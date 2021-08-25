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

func TestGetAllPaymentMethod(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}
	// cleaning data before testing
	db.Migrator().DropTable(&models.Payment{})
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

func TestCheckoutTransaction(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}
	// cleaning data before testing
	db.Migrator().DropTable(&models.Cart{})
	db.Migrator().DropTable(&models.Order{})
	db.Migrator().DropTable(&models.Product{})
	db.Migrator().DropTable(&models.Payment{})
	db.Migrator().DropTable(&models.Seller{})
	db.Migrator().DropTable(&models.Customer{})
	db.Migrator().DropTable(&models.Checkout{})
	db.AutoMigrate(&models.Cart{})
	db.AutoMigrate(&models.Order{})
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.Payment{})
	db.AutoMigrate(&models.Customer{})
	db.AutoMigrate(&models.Seller{})
	db.AutoMigrate(&models.Checkout{})

	// preparate dummy data
	//Customer dummy
	newCustomer := models.Customer{
		Name:     "jojo",
		Email:    "jojo@123",
		Password: "jj123",
		Address:  "bandung",
		Gender:   "M",
	}
	if err := db.Save(&newCustomer).Error; err != nil {
		t.Error(err)
	}
	//Seller Dummy
	newSeller := models.Seller{
		Name:       "ade",
		Email:      "ade@123",
		Password:   "123",
		Address:    "bandung",
		Gender:     "M",
		MarketID:   uint(1),
		CategoryID: uint(1),
		Hp:         628123,
	}
	if err := db.Save(&newSeller).Error; err != nil {
		t.Error(err)
	}
	//Product dummy
	newProduct := models.Product{
		Name:        "bawang",
		Price:       1000,
		Stock:       30,
		Description: "unit per ons",
		SellerID:    uint(1),
	}
	if err := db.Save(&newProduct).Error; err != nil {
		t.Error(err)
	}
	//Cart Dummy
	newCarts := models.Cart{
		Price:      1000,
		Qty:        20,
		ProductID:  uint(1),
		CustomerID: uint(1),
	}
	if err := db.Save(&newCarts).Error; err != nil {
		t.Error(err)
	}
	//Payment dummy
	var newPayment models.Payment
	newPayment.Name = "BCA"
	if err := db.Save(&newPayment).Error; err != nil {
		t.Error(err)
	}
	//Make Token
	var customer models.Customer
	if err := config.DB.Where("email = ?", newCustomer.Email).First(&customer).Error; err != nil {
		t.Error(err)
	}

	reqBody := models.Checkout{
		PaymentID: 1,
	}
	body, _ := json.Marshal(reqBody)
	token, _ := middlewares.CreateToken(int(customer.ID))
	// bearer := "Bearer" + token
	// setting controller
	e := echo.New()
	e.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/checkout")

	//Make Checkout
	middleware.JWT([]byte(constants.SECRET_JWT))(CheckoutTransactionTesting())(context)

	type Response struct {
		TotalQty      int    `json:"total_qty"`
		TotalPrice    int    `json:"total_price"`
		CustomerID    uint   `json:"customer_id"`
		PaymentID     uint   `json:"payment_id"`
		Status        string `json:"status"`
		DeliveryPrice int    `json:"delivery_price"`
	}

	var response Response
	resBody2 := res.Body.String()

	json.Unmarshal([]byte(resBody2), &response)

	t.Run("POST /checkout", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, uint(1), response.CustomerID)
		assert.Equal(t, uint(1), response.PaymentID)
		assert.Equal(t, 20, response.TotalQty)
		// assert.Equal(t, 25000, response.TotalPrice)
		assert.Equal(t, 5000, response.DeliveryPrice)
		assert.Equal(t, "searching", response.Status)
	})
}

func TestFinishedCheckout(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}
	// cleaning data before testing
	db.Migrator().DropTable(&models.Customer{})
	db.Migrator().DropTable(&models.Checkout{})
	db.AutoMigrate(&models.Checkout{})
	db.AutoMigrate(&models.Customer{})

	// preparate dummy data
	//Customer dummy
	newCustomer := models.Customer{
		Name:     "jojo",
		Email:    "jojo@gmail.com",
		Password: "123",
		Address:  "Jl Bogor Raya",
		Gender:   "m",
	}
	if err := db.Save(&newCustomer).Error; err != nil {
		t.Error(err)
	}
	//Checkout dummy
	newCheckout := models.Checkout{
		TotalQty:   1,
		TotalPrice: 6500,
		CustomerID: uint(1),
		PaymentID:  uint(1),
		Status:     "delivery",
	}
	if err := db.Save(&newCheckout).Error; err != nil {
		t.Error(err)
	}
	//Make Token
	var customer models.Customer
	if err := config.DB.Where("email = ?", newCustomer.Email).First(&customer).Error; err != nil {
		t.Error(err)
	}

	token, _ := middlewares.CreateToken(int(customer.ID))
	// setting controller
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/checkout/:checkout_id")
	context.SetParamNames("checkout_id")
	context.SetParamValues("1")

	//Make Checkout
	middleware.JWT([]byte(constants.SECRET_JWT))(FinishTransactionTesting())(context)

	type Response struct {
		Status string `json:"status"`
	}

	var response Response
	resBody2 := res.Body.String()

	json.Unmarshal([]byte(resBody2), &response)

	t.Run("PUT /checkout/:checkout_id", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, "completed", response.Status)
	})

}
