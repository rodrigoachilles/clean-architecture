package database

import (
	"database/sql"
	"fmt"
	"github.com/rodrigoachilles/clean-architecture/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, product_name, price, tax, final_price) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.ProductName, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) FindAll(page int, limit int, sort string) ([]entity.Order, error) {
	offset := (page - 1) * limit

	query := fmt.Sprintf("SELECT * FROM orders ORDER BY %s LIMIT ?, ?", sort)
	rows, err := r.Db.Query(query, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []entity.Order
	for rows.Next() {
		var order entity.Order
		err = rows.Scan()
		if err := rows.Scan(&order.ID, &order.ProductName, &order.Price, &order.Tax, &order.FinalPrice); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}
