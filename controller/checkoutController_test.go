package controller

import (
	"encoding/json"
	"etmarket/project/config"
	"etmarket/project/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/labstack/echo"
)

func TestGetAllPaymentMethod(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}
	// cleaning data before testing
	db.AutoMigrate(&models.Payment{})
	// preparate dummy data
	var newPayments models.Payment
	newPayments.Name = "BCA"
	if err := db.Save(&newPayments).Error; err != nil {
		t.Error(err)
	}
	// setting controller
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/payments")

	GetAllPaymentMethod(context)

	// Unmarshal respose string to struct
	type Response struct {
		Name string `json:"name"`
	}

	var response []Response
	resBody := res.Body.String()

	json.Unmarshal([]byte(resBody), &response)

	t.Run("GET /payments", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, 1, len(response))
		assert.Equal(t, "BCA", response[0].Name)
	})
	db.Migrator().DropTable(&models.Payment{})
}

// func TestCheckoutTransaction(t *testing.T) {
// 	// create database connection
// 	db, err := config.ConfigTest()
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	// cleaning data before testing
// 	db.AutoMigrate(&models.Cart{})
// 	db.AutoMigrate(&models.Order{})
// 	db.AutoMigrate(&models.Product{})
// 	db.AutoMigrate(&models.Checkout{})

// 	// preparate dummy data
// 	//Product dummy
// 	var newProduct models.Product
// 	newProduct.ID = 1
// 	newProduct.Name = "bawang"
// 	newProduct.Price = 1000
// 	if err := db.Save(&newProduct).Error; err != nil {
// 		t.Error(err)
// 	}
// 	//Cart Dummy
// 	var newCarts models.Cart
// 	newCarts.Qty = 20
// 	newCarts.ProductID = 1
// 	if err := db.Save(&newCarts).Error; err != nil {
// 		t.Error(err)
// 	}

// 	// setting controller
// 	e := echo.New()
// 	req := httptest.NewRequest(http.MethodGet, "/", nil)
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath("/payments")

// 	GetAllPaymentMethod(context)

// 	// Unmarshal respose string to struct
// 	type Response struct {
// 		Name string `json:"name"`
// 	}

// 	var response []Response
// 	resBody := res.Body.String()

// 	json.Unmarshal([]byte(resBody), &response)

// 	t.Run("GET /payments", func(t *testing.T) {
// 		assert.Equal(t, 200, res.Code)
// 		assert.Equal(t, 1, len(response))
// 		assert.Equal(t, "BCA", response[0].Name)
// 	})
// 	db.Migrator().DropTable(&models.Payment{})
// }
