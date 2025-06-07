package utils

import "log"

// SendResetEmail es un stub que simula el envío de un correo con el código
// de recuperación. En un entorno real se debería implementar el envío
// mediante un servidor SMTP o un servicio externo.
func SendResetEmail(to, code string) error {
	log.Printf("Enviando correo de recuperación a %s con código %s", to, code)
	return nil
}
