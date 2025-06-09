package ws

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var clients = make(map[string]*websocket.Conn) // userID -> connection
var mu sync.Mutex

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // En producción, usar validación de origen
	},
}

// WebSocketHandler establece una conexión WebSocket para un usuario por su ID.
func WebSocketHandler(c echo.Context) error {
	userID := c.Param("id")
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println("Error al hacer upgrade WebSocket:", err)
		return err
	}

	mu.Lock()
	clients[userID] = conn
	mu.Unlock()

	log.Println("Cliente conectado con ID:", userID)

	defer func() {
		mu.Lock()
		delete(clients, userID)
		mu.Unlock()
		conn.Close()
		log.Println("Cliente desconectado:", userID)
	}()

	// Mantener la conexión abierta
	select {}
}

// EnviarMensaje envía una notificación en tiempo real a un usuario si está conectado.
func EnviarMensaje(userID string, mensaje string) {
	mu.Lock()
	conn, ok := clients[userID]
	mu.Unlock()

	if ok {
		if err := conn.WriteMessage(websocket.TextMessage, []byte(mensaje)); err != nil {
			log.Println("Error al enviar WebSocket:", err)
		}
	}
}
