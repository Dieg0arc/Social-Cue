# Etapa de compilación
FROM golang:1.20 AS builder

WORKDIR /app

# Copia todos los archivos al contenedor de compilación
COPY . .

# Compila el binario
RUN go build -o user-server

# Etapa de ejecución
FROM alpine:latest

WORKDIR /root/

# Copia el binario compilado
COPY --from=builder /app/user-server .

# Expón el puerto que uses (ajusta si es distinto)
EXPOSE 8080

# Ejecuta el binario al iniciar el contenedor
CMD ["./user-server"]
