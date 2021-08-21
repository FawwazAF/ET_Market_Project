package controller

import (
	"bytes"
	"encoding/json"
	"etmarket/project/config"
	"etmarket/project/middlewares"
	"etmarket/project/models"
	"fmt"
	"io/ioutil"
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

func TestCheckoutTransaction(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}
	// cleaning data before testing
	db.AutoMigrate(&models.Cart{})
	db.AutoMigrate(&models.Order{})
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.Payment{})
	db.AutoMigrate(&models.Customer{})
	db.AutoMigrate(&models.Checkout{})

	// preparate dummy data
	//Customer dummy
	var newCustomer models.Customer
	newCustomer.Name = "jojo"
	newCustomer.Email = "jojo@123"
	newCustomer.Password = "jj123"
	newCustomer.Address = "bandung"
	newCustomer.Gender = "M"
	if err := db.Save(&newCustomer).Error; err != nil {
		t.Error(err)
	}
	//Product dummy
	var newProduct models.Product
	newProduct.ID = 1
	newProduct.Name = "bawang"
	newProduct.Price = 1000
	if err := db.Save(&newProduct).Error; err != nil {
		t.Error(err)
	}
	//Cart Dummy
	var newCarts models.Cart
	newCarts.Qty = 20
	newCarts.ProductID = 1
	newCarts.CustomerID = 1
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
	// setting controller
	e := echo.New()
	r := ioutil.NopCloser(bytes.NewReader(body))
	req := httptest.NewRequest(http.MethodPost, "/", r)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/checkout")

	//Make Checkout
	CheckoutTransaction(context)

	type Response struct {
		TotalQty   int    `json:"total_qty"`
		TotalPrice int    `json:"total_price"`
		CustomerID uint   `json:"customer_id"`
		PaymentID  uint   `json:"payment_id"`
		Status     string `json:"status"`
	}

	var response Response
	resBody2 := res.Body.String()

	json.Unmarshal([]byte(resBody2), &response)

	t.Run("POST /checkout", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, uint(1), response.CustomerID)
		assert.Equal(t, uint(1), response.PaymentID)
		assert.Equal(t, 20, response.TotalQty)
		assert.Equal(t, 20000, response.TotalPrice)
		assert.Equal(t, "searching", response.Status)
	})

	db.Migrator().DropTable(&models.Cart{})
	db.Migrator().DropTable(&models.Order{})
	db.Migrator().DropTable(&models.Product{})
	db.Migrator().DropTable(&models.Payment{})
	db.Migrator().DropTable(&models.Customer{})
	db.Migrator().DropTable(&models.Checkout{})
}
