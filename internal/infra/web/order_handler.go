package web

import (
	"encoding/json"
	"github.com/rodrigoachilles/clean-architecture/internal/entity"
	"github.com/rodrigoachilles/clean-architecture/internal/usecase"
	"github.com/rodrigoachilles/clean-architecture/pkg/events"
	"github.com/rodrigoachilles/clean-architecture/pkg/log"
	"net/http"
	"strconv"
)

type OrderHandler struct {
	EventDispatcher   events.EventDispatcherInterface
	OrderRepository   entity.OrderRepositoryInterface
	OrderCreatedEvent events.EventInterface
}

func NewOrderHandler(
	EventDispatcher events.EventDispatcherInterface,
	OrderRepository entity.OrderRepositoryInterface,
	OrderCreatedEvent events.EventInterface,
) *OrderHandler {
	return &OrderHandler{
		EventDispatcher:   EventDispatcher,
		OrderRepository:   OrderRepository,
		OrderCreatedEvent: OrderCreatedEvent,
	}
}

func (h *OrderHandler) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		h.list(w, r)
	case http.MethodPost:
		h.create(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *OrderHandler) create(w http.ResponseWriter, r *http.Request) {
	var dto usecase.OrderInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createOrder := usecase.NewCreateOrderUseCase(h.OrderRepository, h.OrderCreatedEvent, h.EventDispatcher)
	output, err := createOrder.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *OrderHandler) list(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	page, err := strconv.Atoi(values.Get("page"))
	if err != nil {
		msg := "The required field 'page' is not correct."
		log.Logger.Err(err).Msg(msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	limit, err := strconv.Atoi(values.Get("limit"))
	if err != nil {
		msg := "The required field 'limit' is not correct."
		log.Logger.Err(err).Msg(msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	dto := usecase.ListOrderInputDTO{
		Page:  page,
		Limit: limit,
		Sort:  values.Get("sort"),
	}

	listOrders := usecase.NewListOrdersUseCase(h.OrderRepository)
	output, err := listOrders.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
