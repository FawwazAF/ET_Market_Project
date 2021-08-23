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
	mockDBProduct3 = models.Product{
		Name:        "teh",
		Price:       5000,
		Stock:       15,
		Description: "ini teh",
		SellerID:    2,
	}
)

func TestGetAllProductByShopId(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator(),DropTable(&models.Product{})
	config.DB.Migrator().AutoMigrate(&models.Product{})
	config.DB.Save(&mockDBProduct)
	config.DB.Save(&mockDBProduct2)
	config.DB.Save(&mockDBProduct3)
	gotProduct, err := GetAllSellerProduct(1)
	if assert.NoError(t, err) {
		assert.Equal(t, "tomat", [0]gotProduct.name)
		assert.Equal(t, 2000, [0]gotProduct.Price)
		assert.Equal(t, 20, [0]gotProduct.Stock)
		assert.Equal(t, "sawi", [1]gotProduct.name)
		assert.Equal(t, 2500, [1]gotProduct.Price)
		assert.Equal(t, 10, [1]gotProduct.Stock)
	}
}

func TestAddProductToSeller(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator(),DropTable(&models.Product{})
	config.DB.Migrator().AutoMigrate(&models.Product{})
	var product []models.Product
	AddProductToSeller(mockDBProduct)
	AddProductToSeller(mockDBProduct2)
	gotProduct, err := config.DB.Find(&product)
	if assert.NoError(t, err) {
		assert.Equal(t, "tomat", [0]gotProduct.name)
		assert.Equal(t, 2000, [0]gotProduct.Price)
		assert.Equal(t, 20, [0]gotProduct.Stock)
		assert.Equal(t, "sawi", [1]gotProduct.name)
		assert.Equal(t, 2500, [1]gotProduct.Price)
		assert.Equal(t, 10, [1]gotProduct.Stock)
	}
}

func TestEditSellerProduct(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator(),DropTable(&models.Product{})
	config.DB.Migrator().AutoMigrate(&models.Product{})
	var product []models.Product
	config.DB.Save(&mockDBProduct2)
	EditSellerProduct(mockDBProduct)
	gotProduct, err := GetEditProduct(1, 1)
	if assert.NoError(t, err) {
		assert.Equal(t, "tomat", [0]gotProduct.name)
		assert.Equal(t, 2000, [0]gotProduct.Price)
		assert.Equal(t, 20, [0]gotProduct.Stock)
	}
}