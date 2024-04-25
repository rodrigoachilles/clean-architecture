package entity

type OrderRepositoryInterface interface {
	Save(order *Order) error
	FindAll(page int, limit int, sort string) ([]Order, error)
}
