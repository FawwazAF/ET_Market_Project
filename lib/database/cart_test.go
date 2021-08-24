package database

import (
	"etmarket/project/config"
	"etmarket/project/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mock_db_cart = models.Cart{
		CustomerID: 1,
		ProductID:  1,
		Qty:        2,
		Price:      20000,
	}
	mock_db_cart_2 = models.Cart{
		CustomerID: 1,
		ProductID:  2,
		Qty:        5,
		Price:      25000,
	}
	mock_db_cart_3 = models.Cart{
		CustomerID: 1,
		ProductID:  1,
		Qty:        2,
		Price:      40000,
	}
)

func TestInsertProductIntoCartSuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Cart{})
	config.DB.Migrator().AutoMigrate(&models.Cart{})
	var price models.Cart

	cart, err := InsertProductIntoCart(1, 1, 1, price)
	if assert.NoError(t, err) {
		assert.Equal(t, 1, cart.CustomerID)
		assert.Equal(t, 1, cart.ProductID)
		assert.Equal(t, 2, cart.Qty)
		assert.Equal(t, 20000, cart.Price)
	}
}

func TestInsertProductIntoCartError(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Cart{})
	_, err := InsertProductIntoCart(1, 1, 1, mock_db_cart)
	assert.Error(t, err)
}
