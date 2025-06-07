package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"authservice/config"
	"authservice/models"
	"authservice/utils"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type RequestReset struct {
	Email string `json:"email"`
}

type ConfirmReset struct {
	Email       string `json:"email"`
	Code        string `json:"code"`
	NewPassword string `json:"new_password"`
}

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
			log.Println("→ Usuario no encontrado")
			return c.JSON(http.StatusOK, echo.Map{"message": "Si existe una cuenta, se enviará un correo"})
		}

		resets := config.GetCollection("password_resets", cfg)

		// Eliminar códigos antiguos antes de generar uno nuevo
		_, _ = resets.DeleteMany(ctx, bson.M{"user_id": user.ID})

		code, err := utils.GenerateRandomCode()
		if err != nil {
			log.Println("→ Error generando código:", err)
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "No se pudo generar el código"})
		}

		hash, _ := bcrypt.GenerateFromPassword([]byte(code), bcrypt.DefaultCost)
		now := time.Now()
		expires := now.Add(2 * time.Minute)

		log.Println("CREACIÓN DE CÓDIGO")
		log.Println("→ now:", now)
		log.Println("→ expires:", expires)

		_, err = resets.InsertOne(ctx, models.PasswordReset{
			UserID:    user.ID,
			CodeHash:  string(hash),
			ExpiresAt: expires,
			CreatedAt: now,
		})
		if err != nil {
			log.Println("→ Error guardando código:", err)
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error al guardar código"})
		}

		if err := utils.SendResetEmail(user.Email, code); err != nil {
			log.Println("→ Error enviando correo:", err)
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "No se pudo enviar el correo"})
		}

		log.Println("→ Código enviado correctamente a:", user.Email)
		return c.JSON(http.StatusOK, echo.Map{"message": "Instrucciones enviadas"})
	}
}

func ConfirmPasswordReset(cfg *config.Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req ConfirmReset
		if err := c.Bind(&req); err != nil {
			log.Println("→ Error en datos recibidos:", err)
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Datos inválidos"})
		}

		usersCol := config.GetCollection("users", cfg)
		resetsCol := config.GetCollection("password_resets", cfg)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		var user models.User
		if err := usersCol.FindOne(ctx, bson.M{"email": req.Email}).Decode(&user); err != nil {
			log.Println("→ Usuario no encontrado")
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Código inválido"})
		}

		var pr models.PasswordReset
		opts := options.FindOne().SetSort(bson.M{"created_at": -1}) // Si quedara alguno, usa el más reciente
		if err := resetsCol.FindOne(ctx, bson.M{"user_id": user.ID}, opts).Decode(&pr); err != nil {
			log.Println("→ Código no encontrado para el usuario")
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Código inválido"})
		}

		log.Println("VALIDACIÓN DE CÓDIGO")
		log.Println("→ now:", time.Now())
		log.Println("→ expiresAt:", pr.ExpiresAt)
		log.Println("→ diff:", time.Until(pr.ExpiresAt.UTC()))

		if time.Until(pr.ExpiresAt.UTC()) <= 0 {
			log.Println("→ Código expirado")
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Código expirado"})
		}

		if bcrypt.CompareHashAndPassword([]byte(pr.CodeHash), []byte(req.Code)) != nil {
			log.Println("→ Código incorrecto")
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Código incorrecto"})
		}

		newHash, _ := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
		_, err := usersCol.UpdateByID(ctx, user.ID, bson.M{"$set": bson.M{"password": string(newHash)}})
		if err != nil {
			log.Println("→ Error actualizando contraseña:", err)
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "No se pudo actualizar"})
		}

		_, _ = resetsCol.DeleteMany(ctx, bson.M{"user_id": user.ID})
		log.Println("→ Contraseña restablecida correctamente para:", req.Email)

		return c.JSON(http.StatusOK, echo.Map{"message": "Contraseña actualizada"})
	}
}
