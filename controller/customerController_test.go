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
	mockDBCustomer = models.Customer{
		Name:     "Riska",
		Email:    "riska@gmail.com",
		Password: "123",
		Address:  "Jl Bogor Raya",
		Gender:   "F",
	}
	mockDBCustomerLoginSuccess = models.Customer{
		Email:    "riska@gmail.com",
		Password: "123",
	}
	mockDBCustomerLoginWrongEmail = models.Customer{
		Email:    "rizka@gmail.com",
		Password: "123",
	}
	mockDBCustomerLoginWrongPassword = models.Customer{
		Email:    "riska@gmail.com",
		Password: "1234",
	}
)

func AddCustomerData() bool {
	customer := models.Customer{Name: "Riska", Email: "riska@gmail.com", Password: "$2a$12$S/cx.8/YvBHtbDrRsrd/DumMTsM4M3St0wWP1uonsBZysDw6Hk7Mm", Address: "Jl Bogor Raya", Gender: "F"}
	if err := config.DB.Debug().Create(&customer); err != nil {
		return false
	}
	return true
}

func TestCustomerRegisterControllerSuccess(t *testing.T) {
	config.ConfigTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.Customer{})
	config.DB.Migrator().AutoMigrate(&models.Customer{})
	body, _ := json.Marshal(mockDBCustomer)
	r := ioutil.NopCloser(bytes.NewReader(body))
	req := httptest.NewRequest(http.MethodGet, "/", r)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/customer/register")
	if assert.NoError(t, RegisterCustomer(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := rec.Body.String()
		var customer_response models.CustomerResponse
		fmt.Println(body)
		json.Unmarshal([]byte(body), &customer_response)

		assert.Equal(t, true, customer_response.Status)
		assert.Equal(t, "Registration success", customer_response.Message)
	}
}

func TestCustomerRegisterControllerFailEmptyBody(t *testing.T) {
	config.ConfigTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.Customer{})
	config.DB.Migrator().AutoMigrate(&models.Customer{})
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/customer/register")
	if assert.NoError(t, RegisterCustomer(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		body := rec.Body.String()
		var responseCustomer models.CustomerResponse
		fmt.Println(body)
		json.Unmarshal([]byte(body), &responseCustomer)

		assert.Equal(t, false, responseCustomer.Status)
		assert.Equal(t, "Email/Password cannot empty", responseCustomer.Message)
	}
}

func TestCustomerRegisterControllerFailNoTable(t *testing.T) {
	config.ConfigTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.Customer{})
	body, _ := json.Marshal(mockDBCustomer)
	r := ioutil.NopCloser(bytes.NewReader(body))
	req := httptest.NewRequest(http.MethodGet, "/", r)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/customer/register")
	if assert.NoError(t, RegisterCustomer(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		body := rec.Body.String()
		var responseCustomer models.CustomerResponse
		fmt.Println(body)
		json.Unmarshal([]byte(body), &responseCustomer)

		assert.Equal(t, false, responseCustomer.Status)
		assert.Equal(t, "Registration failed", responseCustomer.Message)
	}
}

func TestCustomerLoginControllerSuccess(t *testing.T) {
	config.ConfigTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.Customer{})
	config.DB.Migrator().AutoMigrate(&models.Customer{})
	AddCustomerData()
	body, _ := json.Marshal(&mockDBCustomerLoginSuccess)
	r := ioutil.NopCloser(bytes.NewReader(body))
	req := httptest.NewRequest(http.MethodGet, "/", r)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/customer/login")
	if assert.NoError(t, LoginCustomer(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := rec.Body.String()
		var responseCustomer models.CustomerResponse
		fmt.Println(body)
		json.Unmarshal([]byte(body), &responseCustomer)

		assert.Equal(t, true, responseCustomer.Status)
		assert.Equal(t, "Login success", responseCustomer.Message)
	}
}

func TestCustomerLoginControllerFailNoTable(t *testing.T) {
	config.ConfigTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.Customer{})
	body, _ := json.Marshal(&mockDBCustomerLoginWrongEmail)
	r := ioutil.NopCloser(bytes.NewReader(body))
	req := httptest.NewRequest(http.MethodGet, "/", r)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/customer/login")
	if assert.NoError(t, LoginCustomer(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		body := rec.Body.String()
		var responseCustomer models.CustomerResponse
		fmt.Println(body)
		json.Unmarshal([]byte(body), &responseCustomer)

		assert.Equal(t, false, responseCustomer.Status)
		assert.Equal(t, "Database error", responseCustomer.Message)
	}
}

func TestCustomerLoginControllerFailWrongEmail(t *testing.T) {
	config.ConfigTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.Customer{})
	config.DB.Migrator().AutoMigrate(&models.Customer{})
	AddCustomerData()
	body, _ := json.Marshal(&mockDBCustomerLoginWrongEmail)
	r := ioutil.NopCloser(bytes.NewReader(body))
	req := httptest.NewRequest(http.MethodGet, "/", r)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/customer/login")
	if assert.NoError(t, LoginCustomer(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		body := rec.Body.String()
		var responseCustomer models.CustomerResponse
		fmt.Println(body)
		json.Unmarshal([]byte(body), &responseCustomer)

		assert.Equal(t, false, responseCustomer.Status)
		assert.Equal(t, "Wrong email", responseCustomer.Message)
	}
}

func TestCustomerLoginControllerFailWrongPassword(t *testing.T) {
	config.ConfigTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.Customer{})
	config.DB.Migrator().AutoMigrate(&models.Customer{})
	AddCustomerData()
	body, _ := json.Marshal(&mockDBCustomerLoginWrongPassword)
	r := ioutil.NopCloser(bytes.NewReader(body))
	req := httptest.NewRequest(http.MethodGet, "/", r)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/customer/login")
	if assert.NoError(t, LoginCustomer(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		body := rec.Body.String()
		var responseCustomer models.CustomerResponse
		fmt.Println(body)
		json.Unmarshal([]byte(body), &responseCustomer)

		assert.Equal(t, false, responseCustomer.Status)
		assert.Equal(t, "Wrong password", responseCustomer.Message)
	}
}

// func TestGetDetailCustomer(t *testing.T) {
// 	// create database connection and create controller
// 	config.ConfigTest()
// 	e := echo.New()
// 	config.DB.Migrator().DropTable(&models.Customer{})
// 	config.DB.Migrator().AutoMigrate(&models.Customer{})
// 	AddCustomerData()
// 	token, _ := middlewares.CreateToken(1)

// 	// setting controller

// 	req := httptest.NewRequest(http.MethodGet, "/", nil)
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath("/customers/:customer_id")
// 	context.SetParamNames("customer_id")
// 	context.SetParamValues("1")

// 	GetDetailCustomer(context)

// 	// Unmarshal respose string to struct
// 	type Response struct {
// 		Name  string `json:"name"`
// 		Email string `json:"email"`
// 	}

// 	var response Response
// 	resBody := res.Body.String()

// 	json.Unmarshal([]byte(resBody), &response)

// 	t.Run("GET /customers/:customer_id", func(t *testing.T) {
// 		assert.Equal(t, 200, res.Code) // response.Data.
// 		assert.Equal(t, "Riska", response.Name)
// 		assert.Equal(t, "riska@gmail.com", response.Email)
// 	})
// }
