package controller

import (
	"bytes"
	"encoding/json"
	"etmarket/project/config"
	"etmarket/project/constants"
	"etmarket/project/middlewares"
	"etmarket/project/models"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/stretchr/testify/assert"
)

var (
	mock_customer = models.Customer{
		Name:     "Riska",
		Email:    "riska@gmail.com",
		Password: "123",
		Address:  "Jl Bogor Raya",
		Gender:   "F",
	}
	mock_customer_login = models.Customer{
		Email:    "riska@gmail.com",
		Password: "123",
	}
	mock_customer_wrong_email = models.Customer{
		Email:    "rizka@gmail.com",
		Password: "123",
	}
)

func TestCustomerRegisterControllerSuccess(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}

	// cleaning data before testing
	db.Migrator().DropTable(&models.Customer{})
	db.AutoMigrate(&models.Customer{})

	// setting controller
	body, _ := json.Marshal(mock_customer)
	// r := ioutil.NopCloser(bytes.NewReader(body))
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/customer/register")

	RegisterCustomer(context)

	// Unmarshal respose string to struct
	type Response struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	var response Response
	res_body := res.Body.String()

	json.Unmarshal([]byte(res_body), &response)

	t.Run("POST /customer/register", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, "Riska", response.Name)
		assert.Equal(t, "riska@gmail.com", response.Email)
	})
	db.Migrator().DropTable(&models.Customer{})
}

func TestCustomerLoginControllerSuccess(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}

	// cleaning data before testing
	db.Migrator().DropTable(&models.Customer{})
	db.AutoMigrate(&models.Customer{})

	// preparate dummy data
	customer := models.Customer{
		Name:     "Riska",
		Email:    "riska@gmail.com",
		Password: "123",
		Address:  "Jl Bogor Raya",
		Gender:   "F",
	}
	convert_pwd := []byte(customer.Password) //convert pass from string to byte
	hashed_pwd := EncryptPwd(convert_pwd)
	customer.Password = hashed_pwd
	if err := db.Save(&customer).Error; err != nil {
		t.Error(err)
	}

	// setting controller
	body, _ := json.Marshal(mock_customer_login)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/customer/login")

	LoginCustomer(context)

	// Unmarshal respose string to struct
	type Response struct {
		ID    uint   `json:"id"`
		Email string `json:"email"`
	}

	var response Response
	res_body := res.Body.String()

	json.Unmarshal([]byte(res_body), &response)

	t.Run("POST /customer/login", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, "riska@gmail.com", response.Email)
	})
	db.Migrator().DropTable(&models.Customer{})
}

func TestCustomerLoginControllerWrongEmail(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}

	// cleaning data before testing
	db.Migrator().DropTable(&models.Customer{})
	db.AutoMigrate(&models.Customer{})

	// preparate dummy data
	customer := models.Customer{
		Name:     "Riska",
		Email:    "riska@gmail.com",
		Password: "123",
		Address:  "Jl Bogor Raya",
		Gender:   "F",
	}
	convert_pwd := []byte(customer.Password) //convert pass from string to byte
	hashed_pwd := EncryptPwd(convert_pwd)
	customer.Password = hashed_pwd
	if err := db.Save(&customer).Error; err != nil {
		t.Error(err)
	}

	// setting controller
	body, _ := json.Marshal(mock_customer_wrong_email)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/customer/login")

	LoginCustomer(context)

	var response models.CustomerResponse
	res_body := res.Body.String()

	json.Unmarshal([]byte(res_body), &response)

	t.Run("POST /customer/login", func(t *testing.T) {
		assert.Equal(t, false, response.Status)
		assert.Equal(t, "User Unauthorized. Email or Password not equal", response.Message)
	})
	db.Migrator().DropTable(&models.Customer{})
}

func TestGetDetailCustomer(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}

	// cleaning data before testing
	db.Migrator().DropTable(&models.Customer{})
	db.AutoMigrate(&models.Customer{})

	// preparate dummy data
	customer := models.Customer{
		Name:     "Riska",
		Email:    "riska@gmail.com",
		Password: "123",
		Address:  "Jl Bogor Raya",
		Gender:   "F",
	}
	convert_pwd := []byte(customer.Password) //convert pass from string to byte
	hashed_pwd := EncryptPwd(convert_pwd)
	customer.Password = hashed_pwd
	if err := db.Save(&customer).Error; err != nil {
		t.Error(err)
	}

	//Make Token
	var user models.Customer
	if err := config.DB.Where("email = ?", customer.Email).First(&user).Error; err != nil {
		t.Error(err)
	}

	token, err := middlewares.CreateToken(int(user.ID))
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
	context.SetPath("/customer/:customer_id")
	context.SetParamNames("customer_id")
	context.SetParamValues("1")
	middleware.JWT([]byte(constants.SECRET_JWT))(GetDetailCustomerTesting())(context)

	// Unmarshal respose string to struct
	type Response struct {
		ID      uint   `json:"id"`
		Email   string `json:"email"`
		Name    string `json:"name"`
		Address string `json:"address"`
		Gender  string `json:"gender"`
	}

	var response Response
	res_body := res.Body.String()

	json.Unmarshal([]byte(res_body), &response)

	t.Run("PUT /customer/:customer_id", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, "riska@gmail.com", response.Email)
	})
	db.Migrator().DropTable(&models.Customer{})
}

func TestUpdateCustomer(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}

	// cleaning data before testing
	db.Migrator().DropTable(&models.Customer{})
	db.AutoMigrate(&models.Customer{})

	// preparate dummy data
	customer := models.Customer{
		Name:     "Riska",
		Email:    "riska@gmail.com",
		Password: "123",
		Address:  "Jl Bogor Raya",
		Gender:   "F",
	}
	convert_pwd := []byte(customer.Password) //convert pass from string to byte
	hashed_pwd := EncryptPwd(convert_pwd)
	customer.Password = hashed_pwd
	if err := db.Save(&customer).Error; err != nil {
		t.Error(err)
	}

	//Make Token
	var user models.Customer
	if err := config.DB.Where("email = ?", customer.Email).First(&user).Error; err != nil {
		t.Error(err)
	}

	token, err := middlewares.CreateToken(int(user.ID))
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
	context.SetPath("/customer/:customer_id")
	context.SetParamNames("customer_id")
	context.SetParamValues("1")
	middleware.JWT([]byte(constants.SECRET_JWT))(UpdateCustomerTesting())(context)

	// Unmarshal respose string to struct
	type Response struct {
		ID      uint   `json:"id"`
		Email   string `json:"email"`
		Name    string `json:"name"`
		Address string `json:"address"`
		Gender  string `json:"gender"`
	}

	var response Response
	res_body := res.Body.String()

	json.Unmarshal([]byte(res_body), &response)

	t.Run("PUT /customer/:customer_id", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, "riska@gmail.com", response.Email)
	})
	db.Migrator().DropTable(&models.Customer{})
}

func TestLogoutCustomer(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}

	// cleaning data before testing
	db.Migrator().DropTable(&models.Customer{})
	db.AutoMigrate(&models.Customer{})

	// preparate dummy data
	customer := models.Customer{
		Name:     "Riska",
		Email:    "riska@gmail.com",
		Password: "123",
		Address:  "Jl Bogor Raya",
		Gender:   "F",
	}
	convert_pwd := []byte(customer.Password) //convert pass from string to byte
	hashed_pwd := EncryptPwd(convert_pwd)
	customer.Password = hashed_pwd
	if err := db.Save(&customer).Error; err != nil {
		t.Error(err)
	}

	//Make Token
	var user models.Customer
	if err := config.DB.Where("email = ?", customer.Email).First(&user).Error; err != nil {
		t.Error(err)
	}

	token, err := middlewares.CreateToken(int(user.ID))
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
	context.SetPath("/customer/logout/:customer_id")
	context.SetParamNames("customer_id")
	context.SetParamValues("1")
	middleware.JWT([]byte(constants.SECRET_JWT))(LogoutCustomerTesting())(context)

	// Unmarshal respose string to struct
	type Response struct {
		ID      uint   `json:"id"`
		Email   string `json:"email"`
		Name    string `json:"name"`
		Address string `json:"address"`
		Gender  string `json:"gender"`
		Token   string `json:"token"`
	}

	var response Response
	res_body := res.Body.String()

	json.Unmarshal([]byte(res_body), &response)

	t.Run("PUT /customer/:customer_id", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, "", response.Token)
	})
	db.Migrator().DropTable(&models.Customer{})
}
