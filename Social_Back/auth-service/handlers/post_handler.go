package handlers

import (
    "context"
    "encoding/json"
    "net/http"
    "time"

    "authservice/config"
    "authservice/realtime"

    "github.com/labstack/echo/v4"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func LikePostHandler(cfg *config.Config, hub *realtime.Hub) echo.HandlerFunc {
    return func(c echo.Context) error {
        idParam := c.Param("id")
        postID, err := primitive.ObjectIDFromHex(idParam)
        if err != nil {
            return c.JSON(http.StatusBadRequest, echo.Map{"error": "ID inválido"})
        }

        collection := config.GetCollection("posts", cfg)
        ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancel()

        update := bson.M{"$inc": bson.M{"likes": 1}}
        res := collection.FindOneAndUpdate(ctx, bson.M{"_id": postID}, update)
        if res.Err() != nil {
            return c.JSON(http.StatusNotFound, echo.Map{"error": "Post no encontrado"})
        }

        notif := map[string]any{
            "type":   "like",
            "postId": idParam,
        }
        data, _ := json.Marshal(notif)
        hub.Broadcast(data)

        return c.JSON(http.StatusOK, echo.Map{"message": "Like registrado"})
    }
}
