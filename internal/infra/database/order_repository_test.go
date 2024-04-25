package database

import (
	"database/sql"
	"fmt"
	"github.com/rodrigoachilles/clean-architecture/internal/entity"
	"github.com/stretchr/testify/suite"
	"testing"

	// sqlite3
	_ "github.com/mattn/go-sqlite3"
)

type OrderRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *OrderRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	_, _ = db.Exec("CREATE TABLE orders (id varchar(255) NOT NULL, product_name varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")
	suite.Db = db
}

func (suite *OrderRepositoryTestSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositoryTestSuite))
}

func (suite *OrderRepositoryTestSuite) TestGivenAnOrder_WhenSave_ThenShouldSaveOrder() {
	order, err := entity.NewOrder("123", 10.0, 2.0)
	suite.NoError(err)
	suite.NoError(order.CalculateFinalPrice())
	repo := NewOrderRepository(suite.Db)
	err = repo.Save(order)
	suite.NoError(err)

	var orderResult entity.Order
	err = suite.Db.QueryRow("SELECT id, product_name, price, tax, final_price FROM orders WHERE id = ?", order.ID).
		Scan(&orderResult.ID, &orderResult.ProductName, &orderResult.Price, &orderResult.Tax, &orderResult.FinalPrice)

	suite.NoError(err)
	suite.NotEmpty(order.ID)
	suite.Equal(order.ProductName, orderResult.ProductName)
	suite.Equal(order.Price, orderResult.Price)
	suite.Equal(order.Tax, orderResult.Tax)
	suite.Equal(order.FinalPrice, orderResult.FinalPrice)
}

func (suite *OrderRepositoryTestSuite) TestGivenPageLimitSort_WhenCallFindAll_ThenShouldReturnListOfOrders() {
	repo := NewOrderRepository(suite.Db)
	for i := 1; i <= 20; i++ {
		order, err := entity.NewOrder(fmt.Sprintf("Product %d", i), float64(i)*10.0, float64(i)*2.0)
		suite.NoError(err)
		suite.NoError(order.CalculateFinalPrice())
		err = repo.Save(order)
		suite.NoError(err)
	}

	orders, err := repo.FindAll(1, 10, "price DESC")
	suite.NoError(err)

	suite.Equal(10, len(orders))

	for i, order := range orders {
		suite.NotEmpty(order.ID)
		expectedProductName := fmt.Sprintf("Product %d", 20-i)
		suite.Equal(expectedProductName, order.ProductName)
		expectedPrice := float64(20-i) * 10.0
		suite.Equal(expectedPrice, order.Price)
		expectedTax := float64(20-i) * 2.0
		suite.Equal(expectedTax, order.Tax)
		expectedFinalPrice := float64(20-i) * 12.0
		suite.Equal(expectedFinalPrice, order.FinalPrice)
	}
}
