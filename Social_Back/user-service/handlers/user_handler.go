package handlers

import (
	"context"
	"net/http"
	"time"

	"user-service/config"
	"user-service/models"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetUserByID obtiene un usuario por su ID desde MongoDB
func GetUserByID(c echo.Context) error {
	idParam := c.Param("id")
	userID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "ID inválido"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := config.GetCollection("users") // debe coincidir con tu colección real
	var user models.User
	err = collection.FindOne(ctx, bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Usuario no encontrado"})
	}

	return c.JSON(http.StatusOK, user)
}
