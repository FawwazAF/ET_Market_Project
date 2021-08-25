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

func TestGetAllCompletedDeliveries(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}
	// cleaning data before testing
	db.Migrator().DropTable(&models.Driver{})
	db.Migrator().DropTable(&models.Delivery{})
	db.AutoMigrate(&models.Driver{})
	db.AutoMigrate(&models.Delivery{})

	// preparate dummy data
	//Delivery dummy
	newDelivery := models.Delivery{
		DriverID:   uint(1),
		CheckoutID: uint(1),
		Status:     "completed",
	}
	if err := db.Save(&newDelivery).Error; err != nil {
		t.Error(err)
	}
	//Driver dummy
	newDriver := models.Driver{
		Name:     "jojo",
		Email:    "jojo@gmail.com",
		Password: "123",
		Address:  "Jl Bogor Raya",
		Gender:   "m",
	}
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
	context.SetPath("/driver/history")

	//Make Checkout
	middleware.JWT([]byte(constants.SECRET_JWT))(GetAllCompletedDeliveriesTesting())(context)

	type Response struct {
		DriverID uint   `json:"driver_id"`
		Status   string `json:"status"`
	}

	var response []Response
	resBody2 := res.Body.String()

	json.Unmarshal([]byte(resBody2), &response)

	t.Run("GET /driver/orderlist/order", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, uint(1), response[0].DriverID)
		assert.Equal(t, "completed", response[0].Status)
	})
}
