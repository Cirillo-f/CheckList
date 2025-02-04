FROM golang:1.23-alpine AS builder

WORKDIR /usr/local/src

# Работа с зависимостями
COPY ["go.mod", "go.sum", "./"]
RUN go mod download

# Билдим сервис
COPY . ./
RUN go build -o ./bin/api-service cmd/app/main.go

FROM alpine AS runner

COPY --from=builder /usr/local/src/bin/api-service /

# Копируем .env файл в контейнер (если он есть)
COPY .env .env

EXPOSE 8080

CMD ["/api-service"]

