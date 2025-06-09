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

// CrearPost maneja la creación de un nuevo post
func CrearPost(c echo.Context) error {
	post := new(models.Post)

	post.Titulo = c.FormValue("titulo")
	post.Contenido = c.FormValue("contenido")
	post.Tipo = c.FormValue("tipo")
	post.Categoria = c.FormValue("categoria")
	post.Tags = c.Request().Form["tags"]

	autorID, err := primitive.ObjectIDFromHex(c.FormValue("autorId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "ID de autor inválido"})
	}
	post.AutorID = autorID
	post.ID = primitive.NewObjectID()
	post.FechaCreado = time.Now()

	if fileURL, err := utils.SubirArchivo(c, "archivo"); err == nil {
		post.URLArchivo = fileURL
	}

	collection := config.GetCollection("posts")
	_, err = collection.InsertOne(context.TODO(), post)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error al guardar el post"})
	}

	return c.JSON(http.StatusCreated, post)
}

// ObtenerPosts devuelve todos los posts (con filtros)
func ObtenerPosts(c echo.Context) error {
	tipo := c.QueryParam("tipo")
	categoria := c.QueryParam("categoria")
	tag := c.QueryParam("tag")

	filtro := bson.M{}
	if tipo != "" {
		filtro["tipo"] = tipo
	}
	if categoria != "" {
		filtro["categoria"] = categoria
	}
	if tag != "" {
		filtro["tags"] = tag
	}

	collection := config.GetCollection("posts")
	cursor, err := collection.Find(context.TODO(), filtro)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error al buscar los posts"})
	}
	defer cursor.Close(context.TODO())

	var posts []models.Post
	if err := cursor.All(context.TODO(), &posts); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error al leer los posts"})
	}

	return c.JSON(http.StatusOK, posts)
}

// ObtenerPostPorID devuelve un post por su ID
func ObtenerPostPorID(c echo.Context) error {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "ID no válido"})
	}

	var post models.Post
	collection := config.GetCollection("posts")
	err = collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&post)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Post no encontrado"})
	}

	return c.JSON(http.StatusOK, post)
}

// EliminarPost elimina un post por su ID
func EliminarPost(c echo.Context) error {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "ID no válido"})
	}

	collection := config.GetCollection("posts")
	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error al eliminar"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Post eliminado"})
}

// ToggleLike registra o elimina un like, y crea una notificación si aplica
func ToggleLike(c echo.Context) error {
	postID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "ID del post inválido"})
	}

	var data struct {
		UsuarioID string `json:"usuarioId"`
	}
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Cuerpo inválido"})
	}

	userID, err := primitive.ObjectIDFromHex(data.UsuarioID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "ID de usuario inválido"})
	}

	collection := config.GetCollection("likes")
	filter := bson.M{"postId": postID, "usuarioId": userID}

	var existing models.Like
	err = collection.FindOne(context.TODO(), filter).Decode(&existing)
	if err == nil {
		_, _ = collection.DeleteOne(context.TODO(), filter)
		log.Println("Like quitado:", userID.Hex(), "en post", postID.Hex())
		return c.JSON(http.StatusOK, echo.Map{"message": "Like quitado"})
	}

	like := models.Like{
		ID:        primitive.NewObjectID(),
		PostID:    postID,
		UsuarioID: userID,
		Fecha:     time.Now(),
	}
	_, err = collection.InsertOne(context.TODO(), like)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error al dar like"})
	}

	var post models.Post
	err = config.GetCollection("posts").FindOne(context.TODO(), bson.M{"_id": postID}).Decode(&post)
	if err == nil && post.AutorID != userID {
		log.Println("Like registrado por:", userID.Hex(), "→ notificar a", post.AutorID.Hex())
		utils.CrearNotificacion("like", userID, post.AutorID, "Le dio like a tu post", &postID)
	} else {
		log.Println("No se envía notificación de like (es su propio post o error)")
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "Like registrado"})
}

// CrearComentario agrega un comentario y envía notificación si aplica
func CrearComentario(c echo.Context) error {
	idParam := c.Param("id")
	postID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "ID no válido"})
	}

	var comentario models.Comentario
	if err := c.Bind(&comentario); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Error al parsear"})
	}

	comentario.ID = primitive.NewObjectID()
	comentario.PostID = postID
	comentario.Fecha = time.Now()

	collection := config.GetCollection("comentarios")
	_, err = collection.InsertOne(context.TODO(), comentario)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error al guardar"})
	}

	// Obtener autor del post para verificar si se debe notificar
	var post models.Post
	err = config.GetCollection("posts").FindOne(context.TODO(), bson.M{"_id": postID}).Decode(&post)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "No se pudo verificar el autor del post"})
	}

	if comentario.AutorID != post.AutorID {
		utils.CrearNotificacion("comentario", comentario.AutorID, post.AutorID, "Comentó tu post", &postID)
	}

	return c.JSON(http.StatusCreated, comentario)
}

// ObtenerComentarios devuelve todos los comentarios de un post
func ObtenerComentarios(c echo.Context) error {
	postIDParam := c.Param("id")
	postID, err := primitive.ObjectIDFromHex(postIDParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "ID del post inválido"})
	}

	collection := config.GetCollection("comentarios")
	cursor, err := collection.Find(context.TODO(), bson.M{"postId": postID})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error al buscar los comentarios"})
	}
	defer cursor.Close(context.TODO())

	var comentarios []models.Comentario
	if err := cursor.All(context.TODO(), &comentarios); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error al leer los comentarios"})
	}

	return c.JSON(http.StatusOK, comentarios)
}
