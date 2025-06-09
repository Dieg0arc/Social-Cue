package utils

import (
	"io"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

// SubirArchivo guarda un archivo localmente y retorna la ruta relativa
func SubirArchivo(c echo.Context, campo string) (string, error) {
	file, err := c.FormFile(campo)
	if err != nil {
		return "", err
	}

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Crear carpeta si no existe
	err = os.MkdirAll("uploads", os.ModePerm)
	if err != nil {
		return "", err
	}

	dstPath := filepath.Join("uploads", file.Filename)
	dst, err := os.Create(dstPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return "", err
	}

	// Retornar ruta (puedes adaptarlo para S3 en el futuro)
	return "/" + dstPath, nil
}
