# Etapa de construcción con Go moderno
FROM golang:1.23.3 AS builder

WORKDIR /app

# Copiar los archivos de módulos y descargarlos
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copiar el resto del código fuente
COPY . .

# Compilar el binario
RUN go build -o auth-server

# Etapa final: imagen ligera para ejecución
FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder /app/auth-server .

# Exponer el puerto (ajústalo si usas otro)
EXPOSE 8080

# Comando por defecto
CMD ["./auth-server"]
