# Используем образ golang для сборки
FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY configs ./configs
COPY pkg ./pkg

# Собираем приложение
RUN go build -o /app/dndhelper-discord ./cmd/main.go

# Новый образ, который будет запущен
FROM golang:1.22

WORKDIR /app

# Копируем бинарный файл из предыдущего образа
COPY --from=builder /app/dndhelper-discord ./
COPY --from=builder /app/configs ./configs

# Запускаем приложение при запуске контейнера
CMD ["./dndhelper-discord"]
