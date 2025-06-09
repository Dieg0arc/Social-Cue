package utils

import (
	"context"
	"log"
	"post-service/config"
	"post-service/models"
	"post-service/ws"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CrearNotificacion guarda una notificación y la envía por WebSocket si el receptor está conectado.
// Solo se envía si el emisor y receptor son distintos.
func CrearNotificacion(
	tipo string, // Tipo de notificación: "comentario", "like", "follow", etc.
	de primitive.ObjectID, // Usuario que genera la notificación
	para primitive.ObjectID, // Usuario que debe recibir la notificación
	mensaje string, // Mensaje visible de la notificación
	postID *primitive.ObjectID, // ID del post relacionado (opcional)
) {
	log.Println("Intentando crear notificación")
	log.Println("Tipo:", tipo, "| De:", de.Hex(), "| Para:", para.Hex())

	// Crear la notificación
	noti := models.Notificacion{
		ID:      primitive.NewObjectID(),
		Tipo:    tipo,
		De:      de,
		Para:    para,
		Mensaje: mensaje,
		Leida:   false,
		Fecha:   time.Now(),
		PostID:  postID,
	}

	// Evitar auto-notificación
	if de.Hex() != para.Hex() {
		log.Println("Enviando notificación WebSocket a:", para.Hex())
		ws.EnviarMensaje(para.Hex(), mensaje)
	} else {
		log.Println("No se envía notificación: receptor y emisor son iguales")
	}

	// Guardar en MongoDB
	_, err := config.GetCollection("notificaciones").InsertOne(context.TODO(), noti)
	if err != nil {
		log.Println("Error guardando notificación:", err)
	} else {
		log.Println("Notificación guardada correctamente en MongoDB")
	}
}
