package main

import (
	"database/sql"

	"github.com/devfullcycle/20-CleanArch/internal/entity"
	"github.com/devfullcycle/20-CleanArch/internal/infra/database"
	"github.com/devfullcycle/20-CleanArch/internal/infra/web"
	"github.com/devfullcycle/20-CleanArch/internal/usecase"
	"github.com/devfullcycle/20-CleanArch/pkg/events"
	"github.com/google/wire"
)

// ProviderSet para repositório de pedidos
var OrderRepositorySet = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)

// ProviderSet para eventos
var EventSet = wire.NewSet(
	events.NewEvent,
	wire.Bind(new(events.EventInterface), new(*events.Event)),
)

// ProviderSet para casos de uso
var UseCaseSet = wire.NewSet(
	OrderRepositorySet,
	EventSet,
	usecase.NewCreateOrderUseCase,
	usecase.NewListOrdersUseCase,
)

// ProviderSet para handlers web
var WebHandlerSet = wire.NewSet(
	OrderRepositorySet,
	EventSet,
	web.NewWebOrderHandler,
)

// InitializeCreateOrderUseCase inicializa o caso de uso de criação de pedido
func InitializeCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	wire.Build(UseCaseSet)
	return &usecase.CreateOrderUseCase{}
}

// InitializeListOrdersUseCase inicializa o caso de uso de listagem de pedidos
func InitializeListOrdersUseCase(db *sql.DB) *usecase.ListOrdersUseCase {
	wire.Build(UseCaseSet)
	return &usecase.ListOrdersUseCase{}
}

// InitializeWebOrderHandler inicializa o handler web de pedidos
func InitializeWebOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderHandler {
	wire.Build(WebHandlerSet)
	return &web.WebOrderHandler{}
}
