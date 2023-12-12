# Базовый образ для установки зависимостей и копирования исходного кода
FROM golang:1.21 AS base


# Установите переменную окружения GOPATH
ENV GOPATH=/

# Копируйте исходный код
COPY ./ /app

# Соберите Go-приложение
WORKDIR /app
RUN go mod download
RUN go build -o app ./cmd/main.go

# Образ для выполнения приложения
FROM base AS app

# Задайте команду для выполнения приложения
CMD ["./app"]