FROM golang:alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

# Копируем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем ВСЁ
COPY . .

# Собираем ТОЛЬКО из папки cmd
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd

# Финальный образ
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]