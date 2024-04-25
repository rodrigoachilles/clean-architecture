package entity

import (
	"errors"
	"github.com/rodrigoachilles/clean-architecture/pkg/entity"
)

type Order struct {
	ID          entity.ID
	ProductName string
	Price       float64
	Tax         float64
	FinalPrice  float64
}

func NewOrder(productName string, price float64, tax float64) (*Order, error) {
	order := &Order{
		ID:          entity.NewID(),
		ProductName: productName,
		Price:       price,
		Tax:         tax,
	}
	err := order.IsValid()
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (o *Order) IsValid() error {
	if o.ProductName == "" {
		return errors.New("invalid product name")
	}
	if o.Price <= 0 {
		return errors.New("invalid price")
	}
	if o.Tax <= 0 {
		return errors.New("invalid tax")
	}
	return nil
}

func (o *Order) CalculateFinalPrice() error {
	o.FinalPrice = o.Price + o.Tax
	err := o.IsValid()
	if err != nil {
		return err
	}
	return nil
}
