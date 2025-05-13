# Clean Architecture Challenge

Este projeto implementa uma API de pedidos usando Clean Architecture, com suporte a REST, gRPC e GraphQL.

## Serviços e Portas

- **REST API**: `http://localhost:8000`
  - POST /order - Criar pedido
  - GET /order - Listar pedidos

- **gRPC**: `localhost:50051`
  - CreateOrder - Criar pedido
  - ListOrders - Listar pedidos

- **GraphQL**: `http://localhost:8080`
  - Query: listOrders - Listar pedidos
  - Mutation: createOrder - Criar pedido

- **MySQL**: `localhost:3306`
  - Database: orders
  - User: root
  - Password: root

- **RabbitMQ**: 
  - Management: `http://localhost:15672`
  - AMQP: `localhost:5672`

## Como Executar

### Usando Docker

1. Clone o repositório
2. Execute:
```bash
docker-compose up --build
```

Isso vai subir:
- MySQL
- RabbitMQ
- Aplicação principal (REST, gRPC e GraphQL)
- Cliente gRPC para testes

### Testando os Serviços

#### REST API
```bash
# Criar pedido
curl -X POST http://localhost:8000/order -H "Content-Type: application/json" -d '{"id":"1","price":100,"tax":10}'

# Listar pedidos
curl http://localhost:8000/order
```

#### gRPC
O cliente gRPC já está configurado para testar automaticamente ao subir com Docker.

#### GraphQL
Acesse `http://localhost:8080` no navegador para usar o playground GraphQL.

## Estrutura do Projeto

```
.
├── cmd/
│   ├── ordersystem/    # Aplicação principal
│   └── client/         # Cliente gRPC para testes
├── internal/
│   ├── domain/         # Entidades e regras de negócio
│   ├── usecase/        # Casos de uso
│   └── infra/          # Implementações (REST, gRPC, GraphQL)
└── docker-compose.yaml # Configuração dos containers
``` 