FROM golang:1.23-alpine AS builder

WORKDIR /usr/local/src

# Работа с зависимостями
COPY ["go.mod", "go.sum", "./"]
RUN go mod download

# Билдим сервис
COPY . ./
RUN go build -o ./bin/db-service db/main.go

FROM alpine AS runner

COPY --from=builder /usr/local/src/bin/db-service /

# Копируем .env файл в контейнер (если он есть)
COPY .env .env

EXPOSE 8081

CMD ["/db-service"]
