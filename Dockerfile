# Используем официальный образ Go как базовый
FROM golang:1.23-alpine as builder

# Устанавливаем рабочий каталог
WORKDIR /app

# Копируем файлы модулей Go
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем остальные файлы проекта
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o /main

# Используем образ alpine для финального образа
FROM alpine:latest
WORKDIR /
COPY --from=builder /main /main

# Добавляем CA сертификаты для HTTPS
RUN apk --no-cache add ca-certificates

# Открываем порт, который использует приложение
EXPOSE 8080

# Запускаем приложение
CMD ["/main"]
