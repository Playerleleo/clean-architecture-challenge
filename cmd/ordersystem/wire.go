//go:build wireinject
// +build wireinject

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

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	wire.Build(
		database.NewOrderRepository,
		wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
		events.NewEvent,
		wire.Bind(new(events.EventInterface), new(*events.Event)),
		usecase.NewCreateOrderUseCase,
	)
	return &usecase.CreateOrderUseCase{}
}

func NewListOrdersUseCase(db *sql.DB) *usecase.ListOrdersUseCase {
	wire.Build(
		database.NewOrderRepository,
		wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
		usecase.NewListOrdersUseCase,
	)
	return &usecase.ListOrdersUseCase{}
}

func NewWebOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderHandler {
	wire.Build(
		database.NewOrderRepository,
		wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
		events.NewEvent,
		wire.Bind(new(events.EventInterface), new(*events.Event)),
		web.NewWebOrderHandler,
	)
	return &web.WebOrderHandler{}
}
