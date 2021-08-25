package controller

import (
	"encoding/json"
	"etmarket/project/config"
	"etmarket/project/constants"
	"etmarket/project/middlewares"
	"etmarket/project/models"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/stretchr/testify/assert"
)

func TestGetCheckoutStatusInProgress(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}
	// cleaning data before testing
	db.Migrator().DropTable(&models.Checkout{})
	db.Migrator().DropTable(&models.Customer{})
	db.AutoMigrate(&models.Checkout{})
	db.AutoMigrate(&models.Customer{})

	// preparate dummy data
	// Product dummy
	mocknewCheckout := models.Checkout{
		TotalQty:   10,
		TotalPrice: 20000,
		CustomerID: 1,
		PaymentID:  1,
		Status:     "delivery",
	}
	if err := db.Save(&mocknewCheckout).Error; err != nil {
		t.Error(err)
	}

	var newCustomer models.Customer
	newCustomer.Name = "jojo"
	newCustomer.Email = "jojo@123"
	newCustomer.Password = "jj123"
	newCustomer.Address = "bandung"
	newCustomer.Gender = "M"
	if err := db.Save(&newCustomer).Error; err != nil {
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
	q := make(url.Values)
	q.Set("status", "delivery")
	req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/history")
	middleware.JWT([]byte(constants.SECRET_JWT))(GetCheckoutStatusTesting())(context)

	type Response struct {
		TotalQty   int    `json:"total_qty"`
		TotalPrice int    `json:"total_price"`
		CustomerID uint   `json:"customer_id"`
		PaymentID  uint   `json:"payment_id"`
		Status     string `json:"status"`
	}

	var response []Response
	resBody2 := res.Body.String()

	json.Unmarshal([]byte(resBody2), &response)

	t.Run("GET /history", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, 10, response[0].TotalQty)
		assert.Equal(t, 20000, response[0].TotalPrice)
		assert.Equal(t, uint(1), response[0].CustomerID)
		assert.Equal(t, uint(1), response[0].PaymentID)
		assert.Equal(t, "delivery", response[0].Status)
	})
}
