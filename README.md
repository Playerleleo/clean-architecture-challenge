# Clean Architecture Challenge

Este projeto implementa um sistema de pedidos usando Clean Architecture com múltiplas interfaces:
- REST API
- gRPC
- GraphQL

## Requisitos

- Go 1.21+
- Docker e Docker Compose
- grpcurl (para testar o gRPC)

## Portas Utilizadas

- REST API: 8000
- gRPC: 50051
- GraphQL: 8080
- MySQL: 3306

## Como Executar

1. Clone o repositório
2. Execute o Docker Compose para subir o banco de dados:
```bash
docker compose up -d
```

3. Execute a aplicação:
```bash
go run ./cmd/ordersystem/...
```

## Testando a Aplicação

Você pode usar o arquivo `api.http` para testar todas as interfaces:

### REST API
- Criar pedido: POST http://localhost:8000/order
- Listar pedidos: GET http://localhost:8000/order

### gRPC
- Criar pedido: 
```bash
grpcurl -plaintext -d '{"id": "1", "price": 100, "tax": 10}' localhost:50051 pb.OrderService.CreateOrder
```
- Listar pedidos:
```bash
grpcurl -plaintext localhost:50051 pb.OrderService.ListOrders
```

### GraphQL
- Acesse http://localhost:8080 para o GraphQL Playground
- Use as queries e mutations definidas no arquivo `api.http` 