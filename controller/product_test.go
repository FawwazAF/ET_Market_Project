package controller

import (
	"encoding/json"
	"etmarket/project/config"
	"etmarket/project/models"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestGetAllProductInShop(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}
	// cleaning data before testing
	db.AutoMigrate(&models.Product{})

	// preparate dummy data
	// Product dummy
	mocknewProduct := models.Product{
		Name:        "tomat",
		Price:       1000,
		Stock:       10,
		Description: "per ons",
		SellerID:    1,
	}
	if err := db.Save(&mocknewProduct).Error; err != nil {
		t.Error(err)
	}

	// setting controller
	e := echo.New()
	q := make(url.Values)
	q.Set("seller_id", "1")
	req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/products")

	GetAllProductInShop(context)

	type Response struct {
		Name        string `json:"name"`
		Price       int    `json:"price"`
		Stock       int    `json:"stock"`
		Description string `json:"description"`
		SellerID    uint   `json:"seller_id"`
	}

	var response Response
	resBody2 := res.Body.String()

	json.Unmarshal([]byte(resBody2), &response)

	t.Run("GET /products", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, "tomat", response.Name)
		assert.Equal(t, 1000, response.Price)
		assert.Equal(t, 10, response.Stock)
		assert.Equal(t, "per ons", response.Description)
		assert.Equal(t, 1, response.SellerID)
	})
	db.Migrator().DropTable(&models.Product{})
}

func TestGetSpecificProductInShop(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}
	// cleaning data before testing
	db.AutoMigrate(&models.Product{})

	// preparate dummy data
	// Product dummy
	mocknewProduct := models.Product{
		Name:        "tomat",
		Price:       1000,
		Stock:       10,
		Description: "per ons",
		SellerID:    1,
	}
	if err := db.Save(&mocknewProduct).Error; err != nil {
		t.Error(err)
	}

	// setting controller
	e := echo.New()
	q := make(url.Values)
	q.Set("product_name", "tomat")
	req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/seller")
	context.SetParamNames("seller_id")
	context.SetParamValues("1")

	GetSpecificProductInShop(context)

	type Response struct {
		Name        string `json:"name"`
		Price       int    `json:"price"`
		Stock       int    `json:"stock"`
		Description string `json:"description"`
		SellerID    uint   `json:"seller_id"`
	}

	var response Response
	resBody2 := res.Body.String()

	json.Unmarshal([]byte(resBody2), &response)

	t.Run("GET /seller/product", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, "tomat", response.Name)
		assert.Equal(t, 1000, response.Price)
		assert.Equal(t, 10, response.Stock)
		assert.Equal(t, "per ons", response.Description)
		assert.Equal(t, 1, response.SellerID)
	})
	db.Migrator().DropTable(&models.Product{})
}

func TestGetDetailSpecificProductInShop(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}
	// cleaning data before testing
	db.AutoMigrate(&models.Product{})

	// preparate dummy data
	// Product dummy
	mocknewProduct := models.Product{
		Name:        "tomat",
		Price:       1000,
		Stock:       10,
		Description: "per ons",
		SellerID:    1,
	}
	if err := db.Save(&mocknewProduct).Error; err != nil {
		t.Error(err)
	}

	// setting controller
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/seller")
	context.SetParamNames("seller_id")
	context.SetParamValues("1")
	context.SetPath("/product")
	context.SetParamNames("product_id")
	context.SetParamValues("1")

	GetDetailSpecificProduct(context)

	type Response struct {
		Name        string `json:"name"`
		Price       int    `json:"price"`
		Stock       int    `json:"stock"`
		Description string `json:"description"`
		SellerID    uint   `json:"seller_id"`
	}

	var response Response
	resBody2 := res.Body.String()

	json.Unmarshal([]byte(resBody2), &response)

	t.Run("GET /seller/product", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, "tomat", response.Name)
		assert.Equal(t, 1000, response.Price)
		assert.Equal(t, 10, response.Stock)
		assert.Equal(t, "per ons", response.Description)
		assert.Equal(t, 1, response.SellerID)
	})
	db.Migrator().DropTable(&models.Product{})
}
