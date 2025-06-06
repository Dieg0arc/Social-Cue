package handlers

import (
    "net/http"

    "authservice/realtime"
)

func WebSocketHandler(hub *realtime.Hub) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        hub.HandleWebSocket(w, r)
    }
}
