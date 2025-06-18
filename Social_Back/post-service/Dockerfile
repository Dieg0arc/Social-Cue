# Etapa de construcción
FROM golang:1.23.3 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o post-server

# Etapa de ejecución
FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder /app/post-server .

EXPOSE 8080

CMD ["./post-server"]
