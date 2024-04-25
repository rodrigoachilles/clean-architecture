// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"database/sql"
	"github.com/google/wire"
	"github.com/rodrigoachilles/clean-architecture/internal/entity"
	"github.com/rodrigoachilles/clean-architecture/internal/event"
	"github.com/rodrigoachilles/clean-architecture/internal/infra/database"
	"github.com/rodrigoachilles/clean-architecture/internal/infra/web"
	"github.com/rodrigoachilles/clean-architecture/internal/usecase"
	"github.com/rodrigoachilles/clean-architecture/pkg/events"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from wire.go:

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	orderRepository := database.NewOrderRepository(db)
	orderCreated := event.NewOrderCreated()
	createOrderUseCase := usecase.NewCreateOrderUseCase(orderRepository, orderCreated, eventDispatcher)
	return createOrderUseCase
}

func NewListOrdersUseCase(db *sql.DB) *usecase.ListOrdersUseCase {
	orderRepository := database.NewOrderRepository(db)
	listOrdersUseCase := usecase.NewListOrdersUseCase(orderRepository)
	return listOrdersUseCase
}

func NewOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.OrderHandler {
	orderRepository := database.NewOrderRepository(db)
	orderCreated := event.NewOrderCreated()
	orderHandler := web.NewOrderHandler(eventDispatcher, orderRepository, orderCreated)
	return orderHandler
}

// wire.go:

var setOrderRepositoryDependency = wire.NewSet(database.NewOrderRepository, wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)))

var setEventDispatcherDependency = wire.NewSet(events.NewEventDispatcher, event.NewOrderCreated, wire.Bind(new(events.EventInterface), new(*event.OrderCreated)), wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)))

var setOrderCreatedEvent = wire.NewSet(event.NewOrderCreated, wire.Bind(new(events.EventInterface), new(*event.OrderCreated)))
