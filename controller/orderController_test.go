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

func TestGetAllProgressOrder(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}
	// cleaning data before testing
	db.Migrator().DropTable(&models.Order{})
	db.Migrator().DropTable(&models.Driver{})
	db.Migrator().DropTable(&models.Delivery{})
	db.AutoMigrate(&models.Order{})
	db.AutoMigrate(&models.Driver{})
	db.AutoMigrate(&models.Delivery{})

	// preparate dummy data
	//Delivery dummy
	newDelivery := models.Delivery{
		DriverID:   uint(1),
		CheckoutID: uint(1),
		Status:     "progress",
	}
	if err := db.Save(&newDelivery).Error; err != nil {
		t.Error(err)
	}
	//Order dummy
	newOrder := models.Order{
		Qty:        1,
		Price:      1000,
		Status:     "progress",
		CheckoutID: uint(1),
	}
	if err := db.Save(&newOrder).Error; err != nil {
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
	q := make(url.Values)
	q.Set("checkout_id", "1")
	req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/driver/orderlist/order")

	//Make Checkout
	middleware.JWT([]byte(constants.SECRET_JWT))(GetOrderDriverTesting())(context)

	type Response struct {
		CheckoutID uint `json:"checkout_id"`
		Price      int  `json:"price"`
	}

	var response []Response
	resBody2 := res.Body.String()

	json.Unmarshal([]byte(resBody2), &response)

	t.Run("GET /driver/orderlist/order", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, uint(1), response[0].CheckoutID)
		assert.Equal(t, 1000, response[0].Price)
	})

	db.Migrator().DropTable(&models.Driver{})
	db.Migrator().DropTable(&models.Order{})
	db.Migrator().DropTable(&models.Checkout{})
	db.Migrator().DropTable(&models.Delivery{})
}
