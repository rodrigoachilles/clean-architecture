package usecase

import (
	"github.com/google/uuid"
	"github.com/rodrigoachilles/clean-architecture/internal/entity"
	"github.com/rodrigoachilles/clean-architecture/pkg/events"
)

type OrderInputDTO struct {
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	Tax         float64 `json:"tax"`
}

type OrderOutputDTO struct {
	ID          uuid.UUID `json:"id"`
	ProductName string    `json:"product_name"`
	Price       float64   `json:"price"`
	Tax         float64   `json:"tax"`
	FinalPrice  float64   `json:"final_price"`
}

type CreateOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrderCreated    events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewCreateOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	OrderCreated events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		OrderRepository: OrderRepository,
		OrderCreated:    OrderCreated,
		EventDispatcher: EventDispatcher,
	}
}

func (c *CreateOrderUseCase) Execute(input OrderInputDTO) (OrderOutputDTO, error) {
	order, err := entity.NewOrder(input.ProductName, input.Price, input.Tax)
	if err != nil {
		return OrderOutputDTO{}, err
	}
	_ = order.CalculateFinalPrice()
	if err := c.OrderRepository.Save(order); err != nil {
		return OrderOutputDTO{}, err
	}

	dto := OrderOutputDTO{
		ID:          order.ID,
		ProductName: order.ProductName,
		Price:       order.Price,
		Tax:         order.Tax,
		FinalPrice:  order.Price + order.Tax,
	}

	c.OrderCreated.SetPayload(dto)
	_ = c.EventDispatcher.Dispatch(c.OrderCreated)

	return dto, nil
}
