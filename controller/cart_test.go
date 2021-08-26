package controller

import (
	"bytes"
	"encoding/json"
	"etmarket/project/config"
	"etmarket/project/constants"
	"etmarket/project/middlewares"
	"etmarket/project/models"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
// newDriver = models.Driver{
// 	Name:     "jojo",
// 	Email:    "jojo@gmail.com",
// 	Password: "123",
// 	Address:  "Jl Bogor Raya",
// 	Gender:   "m",
// }
// newCheckout = models.Checkout{
// 	TotalQty:   1,
// 	TotalPrice: 6500,
// 	Status:     "searching",
// }
// newOrder = models.Order{
// 	Qty:        1,
// 	Price:      1000,
// 	Status:     "progress",
// 	CheckoutID: uint(1),
// }
)

func TestInsertProductIntoCart(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}
	// cleaning data before testing
	db.Migrator().DropTable(&models.Product{})
	db.Migrator().DropTable(&models.Customer{})
	db.Migrator().DropTable(&models.Cart{})
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.Customer{})
	db.AutoMigrate(&models.Cart{})

	// preparate dummy data
	//Product dummy
	new_product := models.Product{
		Name:        "bawang",
		Price:       3000,
		Stock:       20,
		Description: "unit per ons",
		SellerID:    uint(1),
	}
	if err := db.Save(&new_product).Error; err != nil {
		t.Error(err)
	}
	//Customer dummy
	new_customer := models.Customer{
		Name:     "jojo",
		Email:    "jojo@gmail.com",
		Password: "123",
		Address:  "Jl Bogor Raya",
		Gender:   "m",
	}
	if err := db.Save(&new_customer).Error; err != nil {
		t.Error(err)
	}
	//Make Token
	var customer models.Customer
	if err := config.DB.Where("email = ?", new_customer.Email).First(&customer).Error; err != nil {
		t.Error(err)
	}
	req_body := models.Cart{
		Qty: 2,
	}
	body, _ := json.Marshal(req_body)
	token, _ := middlewares.CreateToken(int(customer.ID))
	// setting controller
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/seller/:seller_id/product/:product_id")
	context.SetParamNames("seller_id", "product_id")
	context.SetParamValues("1", "1")
	//Make Checkout
	middleware.JWT([]byte(constants.SECRET_JWT))(InsertProductIntoCartTesting())(context)

	type Response struct {
		CustomerID uint `json:"customer_id"`
		ProductID  uint `json:"product_id"`
		Qty        int  `json:"qty"`
		Price      int  `json:"price"`
	}

	var response Response
	req_body2 := res.Body.String()

	json.Unmarshal([]byte(req_body2), &response)

	t.Run("POST /seller/:seller_id/product/:product_id", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, uint(1), response.CustomerID)
		assert.Equal(t, uint(1), response.ProductID)
		assert.Equal(t, 2, response.Qty)
		assert.Equal(t, 6000, response.Price)
	})
}

func TestGetCart(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}
	// cleaning data before testing
	db.Migrator().DropTable(&models.Cart{})
	db.Migrator().DropTable(&models.Customer{})
	db.AutoMigrate(&models.Cart{})
	db.AutoMigrate(&models.Customer{})

	// preparate dummy data
	//Cart dummy
	new_cart := models.Cart{
		CustomerID: 1,
		ProductID:  1,
		Qty:        2,
		Price:      20000,
	}
	if err := db.Save(&new_cart).Error; err != nil {
		t.Error(err)
	}
	//Customer dummy
	new_customer := models.Customer{
		Name:     "jojo",
		Email:    "jojo@gmail.com",
		Password: "123",
		Address:  "Jl Bogor Raya",
		Gender:   "m",
	}
	if err := db.Save(&new_customer).Error; err != nil {
		t.Error(err)
	}
	//Make Token
	var customer models.Customer
	if err := config.DB.Where("email = ?", new_customer.Email).First(&customer).Error; err != nil {
		t.Error(err)
	}

	token, _ := middlewares.CreateToken(int(customer.ID))
	// setting controller
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/cart")

	//Make Checkout
	middleware.JWT([]byte(constants.SECRET_JWT))(GetAllCartsTesting())(context)

	type Response struct {
		CustomerID uint `json:"customer_id"`
		ProductID  uint `json:"product_id"`
		Qty        int  `json:"qty"`
		Price      int  `json:"price"`
	}

	var response []Response
	req_body2 := res.Body.String()

	json.Unmarshal([]byte(req_body2), &response)

	t.Run("GET /cart", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, uint(1), response[0].CustomerID)
		assert.Equal(t, uint(1), response[0].ProductID)
		assert.Equal(t, 2, response[0].Qty)
		assert.Equal(t, 20000, response[0].Price)
	})
}

func TestDeleteProductInCart(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}
	// cleaning data before testing
	db.Migrator().DropTable(&models.Customer{})
	db.Migrator().DropTable(&models.Cart{})
	db.AutoMigrate(&models.Customer{})
	db.AutoMigrate(&models.Cart{})

	// preparate dummy data
	//Cart dummy
	newCart := models.Cart{
		CustomerID: 1,
		ProductID:  1,
		Qty:        2,
		Price:      20000,
	}
	if err := db.Save(&newCart).Error; err != nil {
		t.Error(err)
	}
	//Customer dummy
	new_customer := models.Customer{
		Name:     "jojo",
		Email:    "jojo@gmail.com",
		Password: "123",
		Address:  "Jl Bogor Raya",
		Gender:   "m",
	}
	if err := db.Save(&new_customer).Error; err != nil {
		t.Error(err)
	}
	//Make Token
	var customer models.Customer
	if err := config.DB.Where("email = ?", new_customer.Email).First(&customer).Error; err != nil {
		t.Error(err)
	}
	token, _ := middlewares.CreateToken(int(customer.ID))
	// setting controller
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/cart/:cart_id")
	context.SetParamNames("cart_id")
	context.SetParamValues("1")
	//Make Checkout
	middleware.JWT([]byte(constants.SECRET_JWT))(DeleteProductInCartTesting())(context)

	type Response struct {
		CustomerID uint `json:"customer_id"`
		ProductID  uint `json:"product_id"`
		Qty        int  `json:"qty"`
		Price      int  `json:"price"`
	}

	var response Response
	resBody2 := res.Body.String()

	json.Unmarshal([]byte(resBody2), &response)

	t.Run("DELETE /cart/:cart_id", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, uint(1), response.CustomerID)
		assert.Equal(t, uint(1), response.ProductID)
		assert.Equal(t, 2, response.Qty)
		assert.Equal(t, 20000, response.Price)
	})
}
