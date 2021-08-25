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

func TestGetOrderList(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}
	// cleaning data before testing
	db.AutoMigrate(&models.Driver{})
	db.AutoMigrate(&models.Order{})
	db.AutoMigrate(&models.Checkout{})

	// preparate dummy data
	//Customer dummy
	var newCheckout models.Checkout
	newCheckout.TotalQty = 1
	newCheckout.TotalPrice = 1000
	newCheckout.Status = "searching"
	if err := db.Save(&newCheckout).Error; err != nil {
		t.Error(err)
	}
	//Driver dummy
	var newDriver models.Driver
	newDriver.Name = "jojo"
	newDriver.Email = "jojo@123"
	newDriver.Password = "jj123"
	newDriver.Address = "bandung"
	newDriver.Gender = "M"
	if err := db.Save(&newDriver).Error; err != nil {
		t.Error(err)
	}
	//Make Token
	var driver models.Driver
	if err := config.DB.Where("email = ?", newDriver.Email).First(&driver).Error; err != nil {
		t.Error(err)
	}

	token, err := middlewares.CreateToken(int(driver.ID))
	if err != nil {
		panic(err)
	}
	// setting controller
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/driver/orderlist")
	middleware.JWT([]byte(constants.SECRET_JWT))(GetOrderListTesting())(context)
	// GetOrderList(context)
	// (GetOrderListTesting())(context)

	type Response struct {
		ID     uint   `json:"id"`
		Status string `json:"status"`
	}

	var response []Response
	resBody2 := res.Body.String()

	json.Unmarshal([]byte(resBody2), &response)
	fmt.Println(response)
	t.Run("GET /driver/orderlist", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, uint(1), response[0].ID)
		assert.Equal(t, "searching", response[0].Status)
	})

	db.Migrator().DropTable(&models.Driver{})
	db.Migrator().DropTable(&models.Order{})
	db.Migrator().DropTable(&models.Checkout{})
}

// func TestTakeCheckout(t *testing.T) {
// 	// create database connection
// 	db, err := config.ConfigTest()
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	// cleaning data before testing
// 	db.AutoMigrate(&models.Driver{})
// 	db.AutoMigrate(&models.Delivery{})
// 	db.AutoMigrate(&models.Checkout{})

// 	// preparate dummy data
// 	//Checkout dummy
// 	var newCheckout models.Checkout
// 	newCheckout.TotalQty = 1
// 	newCheckout.TotalPrice = 1000
// 	newCheckout.Status = "searching"
// 	if err := db.Save(&newCheckout).Error; err != nil {
// 		t.Error(err)
// 	}
// 	//Driver dummy
// 	var newDriver models.Driver
// 	newDriver.Name = "jojo"
// 	newDriver.Email = "jojo@123"
// 	newDriver.Password = "jj123"
// 	newDriver.Address = "bandung"
// 	newDriver.Gender = "M"
// 	if err := db.Save(&newDriver).Error; err != nil {
// 		t.Error(err)
// 	}
// 	//Make Token
// 	var driver models.Driver
// 	if err := config.DB.Where("email = ?", newDriver.Email).First(&driver).Error; err != nil {
// 		t.Error(err)
// 	}

// 	token, _ := middlewares.CreateToken(int(driver.ID))
// 	// setting controller
// 	e := echo.New()
// 	req := httptest.NewRequest(http.MethodPost, "/", nil)
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath("/driver/orderlist")
// 	context.SetParamNames("checkout_id")
// 	context.SetParamValues("1")

// 	//Make Checkout
// 	TakeCheckout(context)

// 	type Response struct {
// 		CheckoutID uint   `json:"checkout_id"`
// 		DriverID   string `json:"driver_id"`
// 	}

// 	var response Response
// 	resBody2 := res.Body.String()

// 	json.Unmarshal([]byte(resBody2), &response)

// 	t.Run("POST /driver/orderlist/:checkout_id", func(t *testing.T) {
// 		assert.Equal(t, 200, res.Code)
// 		assert.Equal(t, 1, response.DriverID)
// 		assert.Equal(t, 1, response.CheckoutID)
// 	})

// 	db.Migrator().DropTable(&models.Driver{})
// 	db.Migrator().DropTable(&models.Delivery{})
// 	db.Migrator().DropTable(&models.Checkout{})
// }

// func TestFinishedDelivery(t *testing.T) {
// 	// create database connection
// 	db, err := config.ConfigTest()
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	// cleaning data before testing
// 	db.AutoMigrate(&models.Driver{})
// 	db.AutoMigrate(&models.Delivery{})
// 	db.AutoMigrate(&models.Checkout{})

// 	// preparate dummy data
// 	//Checkout dummy
// 	var newCheckout models.Checkout
// 	newCheckout.TotalQty = 1
// 	newCheckout.TotalPrice = 1000
// 	newCheckout.Status = "searching"
// 	if err := db.Save(&newCheckout).Error; err != nil {
// 		t.Error(err)
// 	}
// 	//Driver dummy
// 	var newDriver models.Driver
// 	newDriver.Name = "jojo"
// 	newDriver.Email = "jojo@123"
// 	newDriver.Password = "jj123"
// 	newDriver.Address = "bandung"
// 	newDriver.Gender = "M"
// 	if err := db.Save(&newDriver).Error; err != nil {
// 		t.Error(err)
// 	}
// 	//Delivery dummy
// 	var newDelivery models.Delivery
// 	newDelivery.CheckoutID = uint(1)
// 	newDelivery.DriverID = uint(1)
// 	newDelivery.Status = "progress"
// 	if err := db.Save(&newDelivery).Error; err != nil {
// 		t.Error(err)
// 	}
// 	//Make Token
// 	var driver models.Driver
// 	if err := config.DB.Where("email = ?", newDriver.Email).First(&driver).Error; err != nil {
// 		t.Error(err)
// 	}

// 	token, _ := middlewares.CreateToken(int(driver.ID))
// 	// setting controller
// 	e := echo.New()
// 	req := httptest.NewRequest(http.MethodPut, "/", nil)
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath("/driver/orderlist")
// 	context.SetParamNames("checkout_id")
// 	context.SetParamValues("1")

// 	//Make Checkout
// 	FinishedDelivery(context)

// 	type Response struct {
// 		Status string `json:"status"`
// 	}

// 	var response Response
// 	resBody2 := res.Body.String()

// 	json.Unmarshal([]byte(resBody2), &response)

// 	t.Run("PUT /driver/orderlist/:checkout_id", func(t *testing.T) {
// 		assert.Equal(t, 200, res.Code)
// 		assert.Equal(t, "completed", response.Status)
// 	})

// 	db.Migrator().DropTable(&models.Driver{})
// 	db.Migrator().DropTable(&models.Delivery{})
// 	db.Migrator().DropTable(&models.Checkout{})
// }
