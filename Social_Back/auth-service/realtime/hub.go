package realtime

import (
    "net/http"

    "github.com/gorilla/websocket"
)

type Hub struct {
    clients    map[*websocket.Conn]bool
    broadcast  chan []byte
    register   chan *websocket.Conn
    unregister chan *websocket.Conn
}

func NewHub() *Hub {
    return &Hub{
        clients:    make(map[*websocket.Conn]bool),
        broadcast:  make(chan []byte),
        register:   make(chan *websocket.Conn),
        unregister: make(chan *websocket.Conn),
    }
}

func (h *Hub) Run() {
    for {
        select {
        case conn := <-h.register:
            h.clients[conn] = true
        case conn := <-h.unregister:
            if _, ok := h.clients[conn]; ok {
                delete(h.clients, conn)
                conn.Close()
            }
        case message := <-h.broadcast:
            for conn := range h.clients {
                conn.WriteMessage(websocket.TextMessage, message)
            }
        }
    }
}

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

func (h *Hub) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        return
    }
    h.register <- conn

    defer func() { h.unregister <- conn }()

    for {
        if _, _, err := conn.ReadMessage(); err != nil {
            break
        }
    }
}

func (h *Hub) Broadcast(message []byte) {
    h.broadcast <- message
}
