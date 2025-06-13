# Etapa de compilación
FROM golang:1.20 AS builder

WORKDIR /app

COPY . .

RUN go build -o auth-server

# Etapa de ejecución
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/auth-server .

EXPOSE 8080

CMD ["./auth-server"]
