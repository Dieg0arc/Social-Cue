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
func CrearNotificacion(tipo string, de, para primitive.ObjectID, mensaje string, postID *primitive.ObjectID) {
	log.Println("Intentando crear notificación")
	log.Println("Tipo:", tipo, "| De:", de.Hex(), "| Para:", para.Hex())

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

	if de.Hex() != para.Hex() {
		log.Println("Enviando notificación WebSocket a:", para.Hex())
		ws.EnviarMensaje(para.Hex(), mensaje)
	} else {
		log.Println("No se envía notificación: receptor y emisor son iguales")
	}

	_, err := config.GetCollection("notificaciones").InsertOne(context.TODO(), noti)
	if err != nil {
		log.Println("Error guardando notificación:", err)
	} else {
		log.Println("Notificación guardada correctamente en MongoDB")
	}
}
