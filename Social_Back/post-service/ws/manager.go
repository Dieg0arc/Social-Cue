package ws

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

// clients mantiene conexiones WebSocket activas mapeadas por userID (string).
var clients = make(map[string]*websocket.Conn)

// mu es un mutex para evitar condiciones de carrera sobre el mapa clients.
var mu sync.Mutex

// upgrader permite actualizar conexiones HTTP a WebSocket.
// ⚠️ En producción deberías validar el origen en CheckOrigin.
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Permitir conexiones desde cualquier origen
	},
}

// WebSocketHandler establece una conexión WebSocket para un usuario dado por su ID (ruta /ws/:id).
// Esta función registra la conexión activa en el mapa global `clients`.
func WebSocketHandler(c echo.Context) error {
	userID := c.Param("id")

	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println("Error al hacer upgrade WebSocket:", err)
		return err
	}

	// Registrar la conexión
	mu.Lock()
	clients[userID] = conn
	mu.Unlock()

	log.Println("Cliente WebSocket conectado con ID:", userID)

	// Esperar a que se desconecte
	defer func() {
		mu.Lock()
		delete(clients, userID)
		mu.Unlock()
		conn.Close()
		log.Println("🔌 Cliente WebSocket desconectado:", userID)
	}()

	// Mantener la conexión abierta indefinidamente
	select {}
}

// EnviarMensaje envía un mensaje de texto en tiempo real a un usuario específico (si está conectado).
func EnviarMensaje(userID string, mensaje string) {
	mu.Lock()
	conn, ok := clients[userID]
	mu.Unlock()

	if ok {
		log.Println("Enviando mensaje a usuario:", userID)
		if err := conn.WriteMessage(websocket.TextMessage, []byte(mensaje)); err != nil {
			log.Println("Error al enviar WebSocket a", userID, ":", err)
		}
	} else {
		log.Println("ℹUsuario", userID, "no está conectado. No se pudo enviar WebSocket.")
	}
}
