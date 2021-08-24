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
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	newDriver = models.Driver{
		Name:     "jojo",
		Email:    "jojo@gmail.com",
		Password: "123",
		Address:  "Jl Bogor Raya",
		Gender:   "m",
	}
	newCheckout = models.Checkout{
		TotalQty:   1,
		TotalPrice: 6500,
		Status:     "searching",
	}
	newDelivery = models.Delivery{
		CheckoutID: uint(1),
		DriverID:   uint(1),
		Status:     "progress",
	}
)

func TestGetOrderList(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}
	// cleaning data before testing
	db.Migrator().DropTable(&models.Driver{})
	db.Migrator().DropTable(&models.Order{})
	db.Migrator().DropTable(&models.Checkout{})
	db.Migrator().DropTable(&models.Delivery{})
	db.AutoMigrate(&models.Driver{})
	db.AutoMigrate(&models.Order{})
	db.AutoMigrate(&models.Checkout{})
	db.AutoMigrate(&models.Delivery{})

	// preparate dummy data
	//Customer dummy
	if err := db.Save(&newCheckout).Error; err != nil {
		t.Error(err)
	}
	//Driver dummy
	if err := db.Save(&newDriver).Error; err != nil {
		t.Error(err)
	}
	//Make Token
	var driver models.Driver
	if err := config.DB.Where("email = ?", newDriver.Email).First(&driver).Error; err != nil {
		t.Error(err)
	}

	token, _ := middlewares.CreateToken(int(driver.ID))
	// setting controller
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/driver/orderlist")

	//Make Checkout
	middleware.JWT([]byte(constants.SECRET_JWT))(GetOrderListTesting())(context)

	type Response struct {
		ID     uint   `json:"id"`
		Status string `json:"status"`
	}

	var response []Response
	resBody2 := res.Body.String()

	json.Unmarshal([]byte(resBody2), &response)

	t.Run("GET /driver/orderlist", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, uint(1), response[0].ID)
		assert.Equal(t, "searching", response[0].Status)
	})

	db.Migrator().DropTable(&models.Driver{})
	db.Migrator().DropTable(&models.Order{})
	db.Migrator().DropTable(&models.Checkout{})
	db.Migrator().DropTable(&models.Delivery{})
}

func TestTakeCheckout(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}
	// cleaning data before testing
	db.Migrator().DropTable(&models.Driver{})
	db.Migrator().DropTable(&models.Delivery{})
	db.Migrator().DropTable(&models.Checkout{})
	db.AutoMigrate(&models.Driver{})
	db.AutoMigrate(&models.Delivery{})
	db.AutoMigrate(&models.Checkout{})

	// preparate dummy data
	//Checkout dummy
	if err := db.Save(&newCheckout).Error; err != nil {
		t.Error(err)
	}
	//Driver dummy
	if err := db.Save(&newDriver).Error; err != nil {
		t.Error(err)
	}
	//Make Token
	var driver models.Driver
	if err := config.DB.Where("email = ?", newDriver.Email).First(&driver).Error; err != nil {
		t.Error(err)
	}

	token, _ := middlewares.CreateToken(int(driver.ID))
	// setting controller
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/driver/orderlist/:checkout_id")
	context.SetParamNames("checkout_id")
	context.SetParamValues("1")

	//Make Checkout
	middleware.JWT([]byte(constants.SECRET_JWT))(TakeCheckoutTesting())(context)

	type Response struct {
		CheckoutID uint   `json:"checkout_id"`
		DriverID   uint   `json:"driver_id"`
		Status     string `json:"status"`
	}

	var response Response
	resBody2 := res.Body.String()

	json.Unmarshal([]byte(resBody2), &response)

	t.Run("POST /driver/orderlist/:checkout_id", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, uint(1), response.DriverID)
		assert.Equal(t, uint(1), response.CheckoutID)
	})

	db.Migrator().DropTable(&models.Driver{})
	db.Migrator().DropTable(&models.Delivery{})
	db.Migrator().DropTable(&models.Checkout{})
}

func TestFinishedDelivery(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}
	// cleaning data before testing
	db.AutoMigrate(&models.Driver{})
	db.AutoMigrate(&models.Delivery{})
	db.AutoMigrate(&models.Checkout{})

	// preparate dummy data
	//Checkout dummy
	if err := db.Save(&newCheckout).Error; err != nil {
		t.Error(err)
	}
	//Driver dummy
	if err := db.Save(&newDriver).Error; err != nil {
		t.Error(err)
	}
	//Delivery dummy
	if err := db.Save(&newDelivery).Error; err != nil {
		t.Error(err)
	}
	//Make Token
	var driver models.Driver
	if err := config.DB.Where("email = ?", newDriver.Email).First(&driver).Error; err != nil {
		t.Error(err)
	}

	token, _ := middlewares.CreateToken(int(driver.ID))
	// setting controller
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/driver/orderlist/:checkout_id")
	context.SetParamNames("checkout_id")
	context.SetParamValues("1")

	//Make Checkout
	middleware.JWT([]byte(constants.SECRET_JWT))(FinishedDeliveryTesting())(context)

	type Response struct {
		Status string `json:"status"`
	}

	var response Response
	resBody2 := res.Body.String()

	json.Unmarshal([]byte(resBody2), &response)

	t.Run("PUT /driver/orderlist/:checkout_id", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, "completed", response.Status)
	})

	db.Migrator().DropTable(&models.Driver{})
	db.Migrator().DropTable(&models.Delivery{})
	db.Migrator().DropTable(&models.Checkout{})
}
