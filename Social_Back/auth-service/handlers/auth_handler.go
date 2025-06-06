package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	"authservice/config"
	"authservice/models"
	"authservice/utils"
)

// LoginHandler procesa el inicio de sesión
func LoginHandler(cfg *config.Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req models.LoginRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Datos de entrada inválidos"})
		}

		collection := config.GetCollection("users", cfg)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		var user models.User
		err := collection.FindOne(ctx, bson.M{"email": req.Email}).Decode(&user)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Credenciales inválidas"})
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Credenciales inválidas"})
		}

		token, err := utils.GenerateJWT(user.ID.Hex(), user.Email, user.Rol)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "No se pudo generar token"})
		}

		return c.JSON(http.StatusOK, models.LoginResponse{
			Token: token,
			User:  user,
		})
	}
}

// RegisterHandler procesa el registro de usuarios
func RegisterHandler(cfg *config.Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req models.RegisterRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Datos de entrada inválidos"})
		}

		collection := config.GetCollection("users", cfg)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error al procesar contraseña"})
		}

		nuevo := models.User{
			ID:                  primitive.NewObjectID(),
			Username:            req.Username,
			Email:               req.Email,
			Password:            string(hash),
			Rol:                 req.Rol,
			CodigoInstitucional: req.CodigoInstitucional,
			CreatedAt:           time.Now(),
			UpdatedAt:           time.Now(),
		}

		_, err = collection.InsertOne(ctx, nuevo)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error al registrar"})
		}

		return c.JSON(http.StatusCreated, models.RegisterResponse{
			Message: "Usuario registrado correctamente",
			UserID:  nuevo.ID.Hex(),
		})
	}
}
