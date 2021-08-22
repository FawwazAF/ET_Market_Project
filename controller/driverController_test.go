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
	mockDBDriver = models.Driver{
		Name:     "Riska",
		Email:    "riska@gmail.com",
		Password: "123",
	}
	mockDBDriverLoginSuccess = models.Driver{
		Email:    "riska@gmail.com",
		Password: "123",
	}
	mockDBDriverLoginWrongEmail = models.Driver{
		Email:    "rizka@gmail.com",
		Password: "123",
	}
	mockDBDriverLoginWrongPassword = models.Driver{
		Email:    "riska@gmail.com",
		Password: "1234",
	}
)

func AddDriverData() bool {
	driver := models.Driver{Name: "Riska", Email: "riska@gmail.com", Password: "$2a$12$S/cx.8/YvBHtbDrRsrd/DumMTsM4M3St0wWP1uonsBZysDw6Hk7Mm"}
	err := config.DB.Debug().Create(&driver)
	if err != nil {
		return false
	}
	return true
}

func TestDriverRegisterControllerSuccess(t *testing.T) {
	config.ConfigTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.Driver{})
	config.DB.Migrator().AutoMigrate(&models.Driver{})
	body, _ := json.Marshal(mockDBDriver)
	r := ioutil.NopCloser(bytes.NewReader(body))
	req := httptest.NewRequest(http.MethodGet, "/", r)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/driver/register")
	if assert.NoError(t, RegisterDriver(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := rec.Body.String()
		var driver_response models.DriverResponse
		fmt.Println(body)
		json.Unmarshal([]byte(body), &driver_response)

		assert.Equal(t, true, driver_response.Status)
		assert.Equal(t, "Registration success", driver_response.Message)
	}
}

func TestDriverRegisterControllerFailEmptyBody(t *testing.T) {
	config.ConfigTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.Driver{})
	config.DB.Migrator().AutoMigrate(&models.Driver{})
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/driver/register")
	if assert.NoError(t, RegisterDriver(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		body := rec.Body.String()
		var responseDriver models.DriverResponse
		fmt.Println(body)
		json.Unmarshal([]byte(body), &responseDriver)

		assert.Equal(t, false, responseDriver.Status)
		assert.Equal(t, "Email/Password cannot empty", responseDriver.Message)
	}
}

func TestDriverRegisterControllerFailNoTable(t *testing.T) {
	config.ConfigTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.Driver{})
	body, _ := json.Marshal(mockDBDriver)
	r := ioutil.NopCloser(bytes.NewReader(body))
	req := httptest.NewRequest(http.MethodGet, "/", r)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/driver/register")
	if assert.NoError(t, RegisterDriver(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		body := rec.Body.String()
		var responseDriver models.DriverResponse
		fmt.Println(body)
		json.Unmarshal([]byte(body), &responseDriver)

		assert.Equal(t, false, responseDriver.Status)
		assert.Equal(t, "Registration failed", responseDriver.Message)
	}
}

func TestDriverLoginControllerSuccess(t *testing.T) {
	config.ConfigTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.Driver{})
	config.DB.Migrator().AutoMigrate(&models.Driver{})
	AddDriverData()
	body, _ := json.Marshal(&mockDBDriverLoginSuccess)
	r := ioutil.NopCloser(bytes.NewReader(body))
	req := httptest.NewRequest(http.MethodGet, "/", r)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/driver/login")
	if assert.NoError(t, LoginDriver(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := rec.Body.String()
		var responseDriver models.DriverResponse
		fmt.Println(body)
		json.Unmarshal([]byte(body), &responseDriver)

		assert.Equal(t, true, responseDriver.Status)
		assert.Equal(t, "Login success", responseDriver.Message)
	}
}

func TestDriverLoginControllerFailNoTable(t *testing.T) {
	config.ConfigTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.Driver{})
	body, _ := json.Marshal(&mockDBDriverLoginWrongEmail)
	r := ioutil.NopCloser(bytes.NewReader(body))
	req := httptest.NewRequest(http.MethodGet, "/", r)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/driver/login")
	if assert.NoError(t, LoginDriver(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		body := rec.Body.String()
		var responseDriver models.DriverResponse
		fmt.Println(body)
		json.Unmarshal([]byte(body), &responseDriver)

		assert.Equal(t, false, responseDriver.Status)
		assert.Equal(t, "Database error", responseDriver.Message)
	}
}

func TestDriverLoginControllerFailWrongEmail(t *testing.T) {
	config.ConfigTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.Driver{})
	config.DB.Migrator().AutoMigrate(&models.Driver{})
	AddDriverData()
	body, _ := json.Marshal(&mockDBDriverLoginWrongEmail)
	r := ioutil.NopCloser(bytes.NewReader(body))
	req := httptest.NewRequest(http.MethodGet, "/", r)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/driver/login")
	if assert.NoError(t, LoginDriver(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		body := rec.Body.String()
		var responseDriver models.DriverResponse
		fmt.Println(body)
		json.Unmarshal([]byte(body), &responseDriver)

		assert.Equal(t, false, responseDriver.Status)
		assert.Equal(t, "Wrong email", responseDriver.Message)
	}
}

func TestDriverLoginControllerFailWrongPassword(t *testing.T) {
	config.ConfigTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.Driver{})
	config.DB.Migrator().AutoMigrate(&models.Driver{})
	AddDriverData()
	body, _ := json.Marshal(&mockDBDriverLoginWrongPassword)
	r := ioutil.NopCloser(bytes.NewReader(body))
	req := httptest.NewRequest(http.MethodGet, "/", r)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/driver/login")
	if assert.NoError(t, LoginDriver(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		body := rec.Body.String()
		var responseDriver models.DriverResponse
		fmt.Println(body)
		json.Unmarshal([]byte(body), &responseDriver)

		assert.Equal(t, false, responseDriver.Status)
		assert.Equal(t, "Wrong password", responseDriver.Message)
	}
}
