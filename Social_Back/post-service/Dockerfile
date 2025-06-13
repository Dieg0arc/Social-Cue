# Etapa de compilación
FROM golang:1.20 AS builder

WORKDIR /app

# Copia el contenido del microservicio
COPY . .

# Compila el ejecutable
RUN go build -o post-server

# Etapa de ejecución
FROM alpine:latest

WORKDIR /root/

# Copia el ejecutable desde la etapa anterior
COPY --from=builder /app/post-server .

# Expón el puerto en el que corre tu microservicio (ajusta si es diferente)
EXPOSE 8080

# Comando por defecto al iniciar el contenedor
CMD ["./post-server"]
