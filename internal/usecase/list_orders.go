package usecase

import (
	"github.com/rodrigoachilles/clean-architecture/internal/entity"
)

type ListOrderInputDTO struct {
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
	Sort  string `json:"sort"`
}

type ListOrderOutputDTO struct {
	Orders *[]entity.Order `json:"orders"`
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(OrderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *ListOrdersUseCase) Execute(input ListOrderInputDTO) (ListOrderOutputDTO, error) {
	if input.Sort == "" {
		input.Sort = "product_name asc"
	}
	orders, err := c.OrderRepository.FindAll(input.Page, input.Limit, input.Sort)
	if err != nil {
		return ListOrderOutputDTO{Orders: nil}, err
	}
	return ListOrderOutputDTO{Orders: &orders}, nil
}
