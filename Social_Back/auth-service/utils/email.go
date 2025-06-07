package utils

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendResetEmail(to, code string) error {
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	user := os.Getenv("SMTP_USER")
	pass := os.Getenv("SMTP_PASS")
	from := os.Getenv("SMTP_FROM")

	if host == "" || port == "" || user == "" || pass == "" || from == "" {
		return fmt.Errorf("SMTP configuration not set")
	}

	addr := fmt.Sprintf("%s:%s", host, port)
	auth := smtp.PlainAuth("", user, pass, host)

	msg := []byte(fmt.Sprintf("To: %s\r\nSubject: Recuperar contraseña\r\n\r\nTu código de recuperación es: %s\r\n", to, code))

	return smtp.SendMail(addr, auth, from, []string{to}, msg)
}
