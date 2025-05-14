FROM golang:1.23-alpine

WORKDIR /app

# Instalar wget e netcat para healthcheck
RUN apk add --no-cache wget netcat-openbsd

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Copiar o arquivo .env para o diretório raiz
COPY cmd/ordersystem/.env /app/app_config.env

RUN go build -o ordersystem ./cmd/ordersystem

# Script de espera para o RabbitMQ
COPY wait-for-rabbitmq.sh /app/wait-for-rabbitmq.sh
RUN chmod +x /app/wait-for-rabbitmq.sh

# Adicionar healthcheck para o RabbitMQ
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://rabbitmq:15672 || exit 1

CMD ["/app/wait-for-rabbitmq.sh", "rabbitmq", "5672", "./ordersystem"] 