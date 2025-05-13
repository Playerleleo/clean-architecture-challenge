FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Copiar o arquivo .env
COPY .env .env

RUN go build -o ordersystem ./cmd/ordersystem

CMD ["./ordersystem"] 