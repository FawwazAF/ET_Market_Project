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

func TestGetAllMarket(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}
	// cleaning data before testing
	db.AutoMigrate(&models.Market{})
	// preparate dummy data
	var newMarket models.Market
	newMarket.Name = "pasar baru"
	newMarket.Address = "bandung"
	if err := db.Save(&newMarket).Error; err != nil {
		t.Error(err)
	}
	// setting controller
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/markets")

	GetAllMarket(context)

	// Unmarshal respose string to struct
	type Response struct {
		Name    string `json:"name"`
		Address string `json:"address"`
	}

	var response Response
	resBody := res.Body.String()

	json.Unmarshal([]byte(resBody), &response)

	t.Run("GET /markets", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, "pasar baru", response.Name)
		assert.Equal(t, "bandung", response.Address)
	})
	db.Migrator().DropTable(&models.Market{})
}

func TestGetSpecificMarket(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}
	// cleaning data before testing
	db.AutoMigrate(&models.Market{})
	// preparate dummy data
	var newMarket models.Market
	newMarket.Name = "pasar baru"
	newMarket.Address = "bandung"
	if err := db.Save(&newMarket).Error; err != nil {
		t.Error(err)
	}
	// setting controller
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/markets")
	context.SetParamNames("market_name")
	context.SetParamValues("pasar")

	GetSpecificMarket(context)

	// Unmarshal respose string to struct
	type Response struct {
		ID      uint   `json:"id"`
		Name    string `json:"name"`
		Address string `json:"address"`
	}

	var response []Response
	resBody := res.Body.String()

	json.Unmarshal([]byte(resBody), &response)

	t.Run("GET /markets/:market_name", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, uint(1), response[0].ID)
		assert.Equal(t, "pasar baru", response[0].Name)
		assert.Equal(t, "bandung", response[0].Address)
	})
	db.Migrator().DropTable(&models.Market{})
}
