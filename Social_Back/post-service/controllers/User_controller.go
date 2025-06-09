package controllers

import (
	"context"
	"log"
	"net/http"
	"post-service/config"
	"post-service/models"
	"post-service/utils"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SeguirUsuario crea una relación seguidor → seguido y notifica al seguido si aplica.
func SeguirUsuario(c echo.Context) error {
	// ID del usuario seguido (desde la URL)
	seguidoIDParam := c.Param("id")

	// Estructura para el JSON del cuerpo
	var body struct {
		SeguidorID string `json:"seguidorId"`
	}
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Error al parsear el cuerpo"})
	}

	// Convertir IDs a ObjectID
	seguidoID, err := primitive.ObjectIDFromHex(seguidoIDParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "ID de seguido inválido"})
	}
	seguidorID, err := primitive.ObjectIDFromHex(body.SeguidorID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "ID de seguidor inválido"})
	}

	// Prevenir que se siga a sí mismo
	if seguidoID == seguidorID {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "No puedes seguirte a ti mismo"})
	}

	// Verificar si ya existe la relación
	collection := config.GetCollection("follows")
	filter := bson.M{"seguidorId": seguidorID, "seguidoId": seguidoID}
	var existente models.Follow
	err = collection.FindOne(context.TODO(), filter).Decode(&existente)
	if err == nil {
		// Ya existe
		return c.JSON(http.StatusConflict, echo.Map{"message": "Ya sigues a este usuario"})
	}

	// Registrar follow
	newFollow := models.Follow{
		ID:         primitive.NewObjectID(),
		SeguidorID: seguidorID,
		SeguidoID:  seguidoID,
		Fecha:      time.Now(),
	}
	_, err = collection.InsertOne(context.TODO(), newFollow)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error al guardar follow"})
	}

	log.Println("Follow registrado:", seguidorID.Hex(), "->", seguidoID.Hex())
	utils.CrearNotificacion("follow", seguidorID, seguidoID, "Empezó a seguirte", nil)

	return c.JSON(http.StatusCreated, echo.Map{"message": "Ahora sigues al usuario"})
}

// DejarDeSeguir elimina la relación seguidor -> seguido
func DejarDeSeguir(c echo.Context) error {
	seguidoID, _ := primitive.ObjectIDFromHex(c.Param("id"))

	var body struct {
		SeguidorID string `json:"seguidorId"`
	}
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Datos inválidos"})
	}
	seguidorID, err := primitive.ObjectIDFromHex(body.SeguidorID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "ID inválido"})
	}

	collection := config.GetCollection("follows")
	_, err = collection.DeleteOne(context.TODO(), bson.M{
		"seguidorId": seguidorID,
		"seguidoId":  seguidoID,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error al dejar de seguir"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Dejaste de seguir al usuario"})
}

// VerSeguidores retorna quienes siguen a un usuario
func VerSeguidores(c echo.Context) error {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))

	collection := config.GetCollection("follows")
	cursor, err := collection.Find(context.TODO(), bson.M{"seguidoId": id}) // 👈 CORRECTO
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error al buscar seguidores"})
	}
	defer cursor.Close(context.TODO())

	var seguidores []models.Follow
	cursor.All(context.TODO(), &seguidores)

	return c.JSON(http.StatusOK, seguidores)
}

// VerSeguidos retorna a quién sigue un usuario
func VerSeguidos(c echo.Context) error {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))

	collection := config.GetCollection("follows")
	cursor, err := collection.Find(context.TODO(), bson.M{"seguidorId": id}) // 👈 CORRECTO
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error al buscar seguidos"})
	}
	defer cursor.Close(context.TODO())

	var seguidos []models.Follow
	cursor.All(context.TODO(), &seguidos)

	return c.JSON(http.StatusOK, seguidos)
}

// Ver notificaciones al usuario
func VerNotificaciones(c echo.Context) error {
	userID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "ID inválido"})
	}

	collection := config.GetCollection("notificaciones")
	cursor, err := collection.Find(context.TODO(), bson.M{"para": userID})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error al buscar notificaciones"})
	}
	defer cursor.Close(context.TODO())

	var notis []models.Notificacion
	if err := cursor.All(context.TODO(), &notis); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error al leer notificaciones"})
	}

	return c.JSON(http.StatusOK, notis)
}

// MarcarNotificacionLeida marca una notificación como leída si pertenece al usuario
func MarcarNotificacionLeida(c echo.Context) error {
	userIDParam := c.Param("id")
	notiIDParam := c.Param("notiId")

	// Convertir a ObjectIDs válidos
	userObjID, err := primitive.ObjectIDFromHex(userIDParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "ID de usuario no válido"})
	}

	notiObjID, err := primitive.ObjectIDFromHex(notiIDParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "ID de notificación no válido"})
	}

	collection := config.GetCollection("notificaciones")
	filter := bson.M{"_id": notiObjID, "para": userObjID}
	update := bson.M{"$set": bson.M{"leida": true}}

	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error actualizando la notificación"})
	}
	if result.MatchedCount == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "❌ No se encontró la notificación para este usuario"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "✅ Notificación marcada como leída"})
}

// EliminarNotificacion elimina una notificación específica si pertenece al usuario autenticado.
func EliminarNotificacion(c echo.Context) error {
	userIDParam := c.Param("id")
	notiIDParam := c.Param("notiId")

	// Convertir parámetros a ObjectID
	userID, err := primitive.ObjectIDFromHex(userIDParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "ID de usuario inválido"})
	}
	notiID, err := primitive.ObjectIDFromHex(notiIDParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "ID de notificación inválido"})
	}

	// Obtener colección
	collection := config.GetCollection("notificaciones")

	// Intentar eliminar la notificación, permitiendo coincidencia por ObjectID o string
	filter := bson.M{
		"_id": notiID,
		"$or": []bson.M{
			{"para": userID},
			{"para": userID.Hex()},
		},
	}

	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error al eliminar la notificación"})
	}
	if result.DeletedCount == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Notificación no encontrada"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Notificación eliminada"})
}
