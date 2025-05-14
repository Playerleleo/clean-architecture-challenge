FROM golang:1.23-alpine AS builder

WORKDIR /app

# Instalar dependências necessárias
RUN apk add --no-cache bash curl

# Copiar apenas os arquivos de dependências primeiro
COPY go.mod go.sum ./
RUN go mod download

# Copiar o código fonte
COPY . .

# Compilar a aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/ordersystem

# Imagem final
FROM alpine:latest

WORKDIR /app

# Copiar apenas o binário compilado
COPY --from=builder /app/main .

EXPOSE 8000 50051 8080

CMD ["./main"] 