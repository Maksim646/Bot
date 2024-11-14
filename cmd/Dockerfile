# Этап сборки
FROM golang:1.23-alpine AS build

# Устанавливаем Git
RUN apk add --no-cache git

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы зависимостей
COPY go.mod go.sum ./
COPY vendor vendor ./

# Копируем весь исходный код
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-X 'main.Version=$(git rev-parse --short HEAD)'" -o /bin/main ./cmd/main.go

# Этап запуска
FROM alpine

# Копируем собранное приложение
COPY --from=build /bin/main /cmd/main

# Указываем команду для запуска приложения
ENTRYPOINT ["/cmd/main"]