//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"
	"github.com/rodrigoachilles/clean-architecture/internal/entity"
	"github.com/rodrigoachilles/clean-architecture/internal/event"
	"github.com/rodrigoachilles/clean-architecture/internal/infra/database"
	"github.com/rodrigoachilles/clean-architecture/internal/infra/web"
	"github.com/rodrigoachilles/clean-architecture/internal/usecase"
	"github.com/rodrigoachilles/clean-architecture/pkg/events"

	"github.com/google/wire"
)

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setOrderCreatedEvent = wire.NewSet(
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
)

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewCreateOrderUseCase,
	)
	return &usecase.CreateOrderUseCase{}
}

func NewListOrdersUseCase(db *sql.DB) *usecase.ListOrdersUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		usecase.NewListOrdersUseCase,
	)
	return &usecase.ListOrdersUseCase{}
}

func NewOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.OrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		web.NewOrderHandler,
	)
	return &web.OrderHandler{}
}
