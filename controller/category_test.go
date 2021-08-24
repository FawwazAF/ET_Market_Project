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

func TestGetAllCategories(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}
	// cleaning data before testing
	db.Migrator().DropTable(&models.Category{})
	db.AutoMigrate(&models.Category{})

	// preparate dummy data
	//Product dummy
	newCategory := models.Category{
		Name: "sayur",
	}
	if err := db.Save(&newCategory).Error; err != nil {
		t.Error(err)
	}
	// setting controller
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/categories")

	//Func
	GetAllCategories(context)

	type Response struct {
		Name string `json:"name"`
	}

	var response []Response
	resBody2 := res.Body.String()

	json.Unmarshal([]byte(resBody2), &response)

	t.Run("GET /categories", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, "sayur", response[0].Name)
	})
}

func TestGetSellerInMarket(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}
	// cleaning data before testing
	db.Migrator().DropTable(&models.Seller{})
	db.AutoMigrate(&models.Seller{})

	// preparate dummy data
	//Seller Dummy
	newSeller := models.Seller{
		Name:       "ade",
		Email:      "ade@123",
		Password:   "123",
		Address:    "bandung",
		Gender:     "M",
		MarketID:   uint(1),
		CategoryID: uint(2),
		Hp:         628123,
	}
	if err := db.Save(&newSeller).Error; err != nil {
		t.Error(err)
	}
	// setting controller
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/markets/:market_id/seller")
	context.SetParamNames("market_id")
	context.SetParamValues("1")

	//Func
	GetAllCategoriesMarketIdController(context)

	type Response struct {
		Name     string `json:"name"`
		MarketID uint   `json:"market_id"`
	}

	var response []Response
	resBody2 := res.Body.String()

	json.Unmarshal([]byte(resBody2), &response)

	t.Run("GET /markets/:market_id/seller", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, uint(1), response[0].MarketID)
	})
}

func TestGetSellerByCategory(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}
	// cleaning data before testing
	db.Migrator().DropTable(&models.Seller{})
	db.Migrator().DropTable(&models.Market{})
	db.Migrator().DropTable(&models.Category{})
	db.AutoMigrate(&models.Seller{})
	db.AutoMigrate(&models.Market{})
	db.AutoMigrate(&models.Category{})

	// preparate dummy data
	//Category dummy
	newCategory := models.Category{
		Name: "sayur",
	}
	if err := db.Save(&newCategory).Error; err != nil {
		t.Error(err)
	}
	//Category dummy
	newMarket := models.Market{
		Name:    "pasar baru",
		Address: "bandung",
	}
	if err := db.Save(&newMarket).Error; err != nil {
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
	// setting controller
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/markets/:market_id/seller/:category_name")
	context.SetParamNames("market_id", "category_name")
	context.SetParamValues("1", "sayur")

	//Func
	GetCategoryNameMarketIdController(context)

	type Response struct {
		Name         string `json:"name"`
		CategoryName string `json:"category_name"`
	}

	var response []Response
	resBody2 := res.Body.String()

	json.Unmarshal([]byte(resBody2), &response)

	t.Run("GET /markets/:market_id/seller/:category_name", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, "ade", response[0].Name)
		assert.Equal(t, "sayur", response[0].CategoryName)
	})
}
