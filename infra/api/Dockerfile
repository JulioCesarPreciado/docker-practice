# Etapa 1: build
FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN go build -o main ./cmd/api

# Etapa 2: runtime
FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/main .
COPY .env .env

EXPOSE 8080
CMD ["./main"]