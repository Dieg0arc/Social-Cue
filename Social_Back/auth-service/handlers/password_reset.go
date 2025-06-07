package handlers

import (
	"context"
	"net/http"
	"time"

	"authservice/config"
	"authservice/models"
	"authservice/utils"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// RequestReset representa la solicitud para iniciar el proceso de recuperación.
type RequestReset struct {
	Email string `json:"email"`
}

// ConfirmReset representa la solicitud para confirmar el cambio de contraseña.
type ConfirmReset struct {
	Email       string `json:"email"`
	Code        string `json:"code"`
	NewPassword string `json:"new_password"`
}

// RequestPasswordReset genera un código y lo envía al correo si el usuario existe.
func RequestPasswordReset(cfg *config.Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req RequestReset
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Datos inválidos"})
		}

		users := config.GetCollection("users", cfg)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		var user models.User
		if err := users.FindOne(ctx, bson.M{"email": req.Email}).Decode(&user); err != nil {
			// Respuesta genérica para no revelar si el correo existe o no
			return c.JSON(http.StatusOK, echo.Map{"message": "Si existe una cuenta, se enviará un correo"})
		}

		code, err := utils.GenerateRandomCode()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "No se pudo generar el código"})
		}
		hash, _ := bcrypt.GenerateFromPassword([]byte(code), bcrypt.DefaultCost)

		resets := config.GetCollection("password_resets", cfg)
		_, err = resets.InsertOne(ctx, models.PasswordReset{
			UserID:    user.ID,
			CodeHash:  string(hash),
			ExpiresAt: time.Now().Add(15 * time.Minute),
		})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error al guardar código"})
		}

		if err := utils.SendResetEmail(user.Email, code); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "No se pudo enviar el correo"})
		}

		return c.JSON(http.StatusOK, echo.Map{"message": "Instrucciones enviadas"})
	}
}

// ConfirmPasswordReset verifica el código y actualiza la contraseña del usuario.
func ConfirmPasswordReset(cfg *config.Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req ConfirmReset
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Datos inválidos"})
		}

		usersCol := config.GetCollection("users", cfg)
		resetsCol := config.GetCollection("password_resets", cfg)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		var user models.User
		if err := usersCol.FindOne(ctx, bson.M{"email": req.Email}).Decode(&user); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Código inválido"})
		}

		var pr models.PasswordReset
		if err := resetsCol.FindOne(ctx, bson.M{"user_id": user.ID}).Decode(&pr); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Código inválido"})
		}

		if time.Now().After(pr.ExpiresAt) {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Código expirado"})
		}

		if bcrypt.CompareHashAndPassword([]byte(pr.CodeHash), []byte(req.Code)) != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Código incorrecto"})
		}

		newHash, _ := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
		_, err := usersCol.UpdateByID(ctx, user.ID, bson.M{"$set": bson.M{"password": string(newHash)}})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "No se pudo actualizar"})
		}

		resetsCol.DeleteMany(ctx, bson.M{"user_id": user.ID})

		return c.JSON(http.StatusOK, echo.Map{"message": "Contraseña actualizada"})
	}
}
