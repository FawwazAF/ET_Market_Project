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
	mock_driver = models.Driver{
		Name:     "Riska",
		Email:    "riska@gmail.com",
		Password: "123",
	}
	mock_driver_login = models.Driver{
		Email:    "riska@gmail.com",
		Password: "123",
	}
	mock_driver_wrong_email = models.Driver{
		Email:    "rizka@gmail.com",
		Password: "123",
	}
)

func TestDriverRegisterControllerSuccess(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}

	// cleaning data before testing
	db.Migrator().DropTable(&models.Driver{})
	db.AutoMigrate(&models.Driver{})

	// setting controller
	body, _ := json.Marshal(mock_driver)
	// r := ioutil.NopCloser(bytes.NewReader(body))
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/driver/register")

	RegisterDriver(context)

	// Unmarshal respose string to struct
	type Response struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	var response Response
	res_body := res.Body.String()

	json.Unmarshal([]byte(res_body), &response)

	t.Run("POST /driver/register", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, "Riska", response.Name)
		assert.Equal(t, "riska@gmail.com", response.Email)
	})
	db.Migrator().DropTable(&models.Driver{})
}

func TestDriverLoginControllerSuccess(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}

	// cleaning data before testing
	db.Migrator().DropTable(&models.Driver{})
	db.AutoMigrate(&models.Driver{})

	// preparate dummy data
	driver := models.Driver{
		Name:     "Riska",
		Email:    "riska@gmail.com",
		Password: "123",
		Address:  "Jl Bogor Raya",
		Gender:   "F",
	}
	convert_pwd := []byte(driver.Password) //convert pass from string to byte
	hashed_pwd := EncryptPwd(convert_pwd)
	driver.Password = hashed_pwd
	if err := db.Save(&driver).Error; err != nil {
		t.Error(err)
	}

	// setting controller
	body, _ := json.Marshal(mock_driver_login)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/driver/login")

	LoginDriver(context)

	// Unmarshal respose string to struct
	type Response struct {
		ID    uint   `json:"id"`
		Email string `json:"email"`
	}

	var response Response
	res_body := res.Body.String()

	json.Unmarshal([]byte(res_body), &response)

	t.Run("POST /driver/login", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, "riska@gmail.com", response.Email)
	})
	db.Migrator().DropTable(&models.Driver{})
}

func TestDriverLoginControllerWrongEmail(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}

	// cleaning data before testing
	db.Migrator().DropTable(&models.Driver{})
	db.AutoMigrate(&models.Driver{})

	// preparate dummy data
	driver := models.Driver{
		Name:     "Riska",
		Email:    "riska@gmail.com",
		Password: "123",
		Address:  "Jl Bogor Raya",
		Gender:   "F",
	}
	convert_pwd := []byte(driver.Password) //convert pass from string to byte
	hashed_pwd := EncryptPwd(convert_pwd)
	driver.Password = hashed_pwd
	if err := db.Save(&driver).Error; err != nil {
		t.Error(err)
	}

	// setting controller
	body, _ := json.Marshal(mock_driver_wrong_email)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/driver/login")

	LoginDriver(context)

	var response models.DriverResponse
	res_body := res.Body.String()

	json.Unmarshal([]byte(res_body), &response)

	t.Run("POST /driver/login", func(t *testing.T) {
		assert.Equal(t, false, response.Status)
		assert.Equal(t, "User Unauthorized. Email or Password not equal", response.Message)
	})
	db.Migrator().DropTable(&models.Driver{})
}

func TestGetDetailDriver(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}

	// cleaning data before testing
	db.Migrator().DropTable(&models.Driver{})
	db.AutoMigrate(&models.Driver{})

	// preparate dummy data
	driver := models.Driver{
		Name:     "Riska",
		Email:    "riska@gmail.com",
		Password: "123",
		Address:  "Jl Bogor Raya",
		Gender:   "F",
	}
	convert_pwd := []byte(driver.Password) //convert pass from string to byte
	hashed_pwd := EncryptPwd(convert_pwd)
	driver.Password = hashed_pwd
	if err := db.Save(&driver).Error; err != nil {
		t.Error(err)
	}

	//Make Token
	var user models.Driver
	if err := config.DB.Where("email = ?", driver.Email).First(&user).Error; err != nil {
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
	context.SetPath("/driver/:driver_id")
	context.SetParamNames("driver_id")
	context.SetParamValues("1")
	middleware.JWT([]byte(constants.SECRET_JWT))(GetDetailDriverTesting())(context)

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

	t.Run("PUT /driver/:driver_id", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, "riska@gmail.com", response.Email)
	})
	db.Migrator().DropTable(&models.Driver{})
}

func TestUpdateDriver(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}

	// cleaning data before testing
	db.Migrator().DropTable(&models.Driver{})
	db.AutoMigrate(&models.Driver{})

	// preparate dummy data
	driver := models.Driver{
		Name:     "Riska",
		Email:    "riska@gmail.com",
		Password: "123",
		Address:  "Jl Bogor Raya",
		Gender:   "F",
	}
	convert_pwd := []byte(driver.Password) //convert pass from string to byte
	hashed_pwd := EncryptPwd(convert_pwd)
	driver.Password = hashed_pwd
	if err := db.Save(&driver).Error; err != nil {
		t.Error(err)
	}

	//Make Token
	var user models.Driver
	if err := config.DB.Where("email = ?", driver.Email).First(&user).Error; err != nil {
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
	context.SetPath("/driver/:driver_id")
	context.SetParamNames("driver_id")
	context.SetParamValues("1")
	middleware.JWT([]byte(constants.SECRET_JWT))(UpdateDriverTesting())(context)

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

	t.Run("PUT /driver/:driver_id", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, "riska@gmail.com", response.Email)
	})
	db.Migrator().DropTable(&models.Driver{})
}

func TestLogoutDriver(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}

	// cleaning data before testing
	db.Migrator().DropTable(&models.Driver{})
	db.AutoMigrate(&models.Driver{})

	// preparate dummy data
	driver := models.Driver{
		Name:     "Riska",
		Email:    "riska@gmail.com",
		Password: "123",
		Address:  "Jl Bogor Raya",
		Gender:   "F",
	}
	convert_pwd := []byte(driver.Password) //convert pass from string to byte
	hashed_pwd := EncryptPwd(convert_pwd)
	driver.Password = hashed_pwd
	if err := db.Save(&driver).Error; err != nil {
		t.Error(err)
	}

	//Make Token
	var user models.Driver
	if err := config.DB.Where("email = ?", driver.Email).First(&user).Error; err != nil {
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
	context.SetPath("/driver/logout/:driver_id")
	context.SetParamNames("driver_id")
	context.SetParamValues("1")
	middleware.JWT([]byte(constants.SECRET_JWT))(LogoutDriverTesting())(context)

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

	t.Run("PUT /driver/:driver_id", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, "", response.Token)
	})
	db.Migrator().DropTable(&models.Driver{})
}
