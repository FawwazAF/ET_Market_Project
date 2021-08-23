package database

import (
	"etmarket/project/config"
	"etmarket/project/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mockDBProduct = models.Product{
		Name:        "tomat",
		Price:       2000,
		Stock:       20,
		Description: "ini tomat",
		SellerID:    1,
	}
	mockDBProduct2 = models.Product{
		Name:        "sawi",
		Price:       2500,
		Stock:       10,
		Description: "ini sawi",
		SellerID:    1,
	}
)

func TestGetAllProductByShopId(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator(),DropTable(&models.Product{})
	config.DB.Migrator().AutoMigrate(&models.Product{})
	config.DB.Save(&mockDBProduct)
	gotProduct, err := GetAllProductByShopId(1)
	if assert.NoError(t, err) {
		assert.Equal(t, "tomat", gotProduct.name)
		assert.Equal(t, 2000, gotProduct.Price)
		assert.Equal(t, 20, gotProduct.Stock)
	}
}

func TestGetSpecificProductByShopId(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator(),DropTable(&models.Product{})
	config.DB.Migrator().AutoMigrate(&models.Product{})
	config.DB.Save(&mockDBProduct)
	config.DB.Save(&mockDBProduct2)
	gotProduct, err := GetAllProductByShopId(1, sawi)
	if assert.NoError(t, err) {
		assert.Equal(t, "sawi", [1]gotProduct.name)
		assert.Equal(t, 2500, [1]gotProduct.Price)
		assert.Equal(t, 10, [1]gotProduct.Stock)
	}
}

func TestGetSpecificProductById(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator(),DropTable(&models.Product{})
	config.DB.Migrator().AutoMigrate(&models.Product{})
	config.DB.Save(&mockDBProduct)
	config.DB.Save(&mockDBProduct2)
	gotProduct, err := GetAllProductByShopId(1, 1)
	if assert.NoError(t, err) {
		assert.Equal(t, "tomat", [0]gotProduct.name)
		assert.Equal(t, 2000, [0]gotProduct.Price)
		assert.Equal(t, 20, [0]gotProduct.Stock)
	}
}