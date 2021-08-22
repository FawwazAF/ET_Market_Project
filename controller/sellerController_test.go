package controller

import (
	"bytes"
	"encoding/json"
	"etmarket/project/config"
	"etmarket/project/models"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	mockDBSeller = models.Seller{
		Name:     "Riska",
		Email:    "riska@gmail.com",
		Password: "123",
	}
	mockDBSellerLoginSuccess = models.Seller{
		Email:    "riska@gmail.com",
		Password: "123",
	}
	mockDBSellerLoginWrongEmail = models.Seller{
		Email:    "rizka@gmail.com",
		Password: "123",
	}
	mockDBSellerLoginWrongPassword = models.Seller{
		Email:    "riska@gmail.com",
		Password: "1234",
	}
)

func AddSellerData() bool {
	seller := models.Seller{Name: "Riska", Email: "riska@gmail.com", Password: "$2a$12$S/cx.8/YvBHtbDrRsrd/DumMTsM4M3St0wWP1uonsBZysDw6Hk7Mm"}
	if err := config.DB.Debug().Create(&seller); err != nil {
		return false
	}
	return true
}

func TestSellerRegisterControllerSuccess(t *testing.T) {
	config.ConfigTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.Seller{})
	config.DB.Migrator().AutoMigrate(&models.Seller{})
	body, _ := json.Marshal(mockDBSeller)
	r := ioutil.NopCloser(bytes.NewReader(body))
	req := httptest.NewRequest(http.MethodGet, "/", r)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/seller/register")
	if assert.NoError(t, RegisterSeller(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := rec.Body.String()
		var seller_response models.SellerResponse
		fmt.Println(body)
		json.Unmarshal([]byte(body), &seller_response)

		assert.Equal(t, true, seller_response.Status)
		assert.Equal(t, "Registration success", seller_response.Message)
	}
}

func TestSellerRegisterControllerFailEmptyBody(t *testing.T) {
	config.ConfigTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.Seller{})
	config.DB.Migrator().AutoMigrate(&models.Seller{})
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/seller/register")
	if assert.NoError(t, RegisterSeller(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		body := rec.Body.String()
		var responseSeller models.SellerResponse
		fmt.Println(body)
		json.Unmarshal([]byte(body), &responseSeller)

		assert.Equal(t, false, responseSeller.Status)
		assert.Equal(t, "Email/Password cannot empty", responseSeller.Message)
	}
}

func TestSellerRegisterControllerFailNoTable(t *testing.T) {
	config.ConfigTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.Seller{})
	body, _ := json.Marshal(mockDBSeller)
	r := ioutil.NopCloser(bytes.NewReader(body))
	req := httptest.NewRequest(http.MethodGet, "/", r)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/seller/register")
	if assert.NoError(t, RegisterSeller(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		body := rec.Body.String()
		var responseSeller models.SellerResponse
		fmt.Println(body)
		json.Unmarshal([]byte(body), &responseSeller)

		assert.Equal(t, false, responseSeller.Status)
		assert.Equal(t, "Registration failed", responseSeller.Message)
	}
}

func TestSellerLoginControllerSuccess(t *testing.T) {
	config.ConfigTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.Seller{})
	config.DB.Migrator().AutoMigrate(&models.Seller{})
	AddSellerData()
	body, _ := json.Marshal(&mockDBSellerLoginSuccess)
	r := ioutil.NopCloser(bytes.NewReader(body))
	req := httptest.NewRequest(http.MethodGet, "/", r)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/seller/login")
	if assert.NoError(t, LoginSeller(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := rec.Body.String()
		var responseSeller models.SellerResponse
		fmt.Println(body)
		json.Unmarshal([]byte(body), &responseSeller)

		assert.Equal(t, true, responseSeller.Status)
		assert.Equal(t, "Login success", responseSeller.Message)
	}
}

func TestSellerLoginControllerFailNoTable(t *testing.T) {
	config.ConfigTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.Seller{})
	body, _ := json.Marshal(&mockDBSellerLoginWrongEmail)
	r := ioutil.NopCloser(bytes.NewReader(body))
	req := httptest.NewRequest(http.MethodGet, "/", r)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/seller/login")
	if assert.NoError(t, LoginSeller(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		body := rec.Body.String()
		var responseSeller models.SellerResponse
		fmt.Println(body)
		json.Unmarshal([]byte(body), &responseSeller)

		assert.Equal(t, false, responseSeller.Status)
		assert.Equal(t, "Database error", responseSeller.Message)
	}
}

func TestSellerLoginControllerFailWrongEmail(t *testing.T) {
	config.ConfigTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.Seller{})
	config.DB.Migrator().AutoMigrate(&models.Seller{})
	AddSellerData()
	body, _ := json.Marshal(&mockDBSellerLoginWrongEmail)
	r := ioutil.NopCloser(bytes.NewReader(body))
	req := httptest.NewRequest(http.MethodGet, "/", r)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/seller/login")
	if assert.NoError(t, LoginSeller(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		body := rec.Body.String()
		var responseSeller models.SellerResponse
		fmt.Println(body)
		json.Unmarshal([]byte(body), &responseSeller)

		assert.Equal(t, false, responseSeller.Status)
		assert.Equal(t, "Wrong email", responseSeller.Message)
	}
}

func TestSellerLoginControllerFailWrongPassword(t *testing.T) {
	config.ConfigTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.Seller{})
	config.DB.Migrator().AutoMigrate(&models.Seller{})
	AddSellerData()
	body, _ := json.Marshal(&mockDBSellerLoginWrongPassword)
	r := ioutil.NopCloser(bytes.NewReader(body))
	req := httptest.NewRequest(http.MethodGet, "/", r)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/seller/login")
	if assert.NoError(t, LoginSeller(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		body := rec.Body.String()
		var responseSeller models.SellerResponse
		fmt.Println(body)
		json.Unmarshal([]byte(body), &responseSeller)

		assert.Equal(t, false, responseSeller.Status)
		assert.Equal(t, "Wrong password", responseSeller.Message)
	}
}
