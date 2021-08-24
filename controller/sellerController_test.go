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
	mock_seller = models.Seller{
		Name:     "Riska",
		Email:    "riska@gmail.com",
		Password: "123",
	}
	mock_seller_login = models.Seller{
		Email:    "riska@gmail.com",
		Password: "123",
	}
	mock_seller_wrong_email = models.Seller{
		Email:    "rizka@gmail.com",
		Password: "123",
	}
)

func TestSellerRegisterControllerSuccess(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}

	// cleaning data before testing
	db.Migrator().DropTable(&models.Seller{})
	db.AutoMigrate(&models.Seller{})

	// setting controller
	body, _ := json.Marshal(mock_seller)
	// r := ioutil.NopCloser(bytes.NewReader(body))
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/seller/register")

	RegisterSeller(context)

	// Unmarshal respose string to struct
	type Response struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	var response Response
	res_body := res.Body.String()

	json.Unmarshal([]byte(res_body), &response)

	t.Run("POST /seller/register", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, "Riska", response.Name)
		assert.Equal(t, "riska@gmail.com", response.Email)
	})
	db.Migrator().DropTable(&models.Seller{})
}

func TestSellerLoginControllerSuccess(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}

	// cleaning data before testing
	db.Migrator().DropTable(&models.Seller{})
	db.AutoMigrate(&models.Seller{})

	// preparate dummy data
	seller := models.Seller{
		Name:     "Riska",
		Email:    "riska@gmail.com",
		Password: "123",
		Address:  "Jl Bogor Raya",
		Gender:   "F",
		Hp:       628123828282,
	}
	convert_pwd := []byte(seller.Password) //convert pass from string to byte
	hashed_pwd := EncryptPwd(convert_pwd)
	seller.Password = hashed_pwd
	if err := db.Save(&seller).Error; err != nil {
		t.Error(err)
	}

	// setting controller
	body, _ := json.Marshal(mock_seller_login)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/seller/login")

	LoginSeller(context)

	// Unmarshal respose string to struct
	type Response struct {
		ID    uint   `json:"id"`
		Email string `json:"email"`
	}

	var response Response
	res_body := res.Body.String()

	json.Unmarshal([]byte(res_body), &response)

	t.Run("POST /seller/login", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, "riska@gmail.com", response.Email)
	})
	db.Migrator().DropTable(&models.Seller{})
}

func TestSellerLoginControllerWrongEmail(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}

	// cleaning data before testing
	db.Migrator().DropTable(&models.Seller{})
	db.AutoMigrate(&models.Seller{})

	// preparate dummy data
	seller := models.Seller{
		Name:     "Riska",
		Email:    "riska@gmail.com",
		Password: "123",
		Address:  "Jl Bogor Raya",
		Gender:   "F",
	}
	convert_pwd := []byte(seller.Password) //convert pass from string to byte
	hashed_pwd := EncryptPwd(convert_pwd)
	seller.Password = hashed_pwd
	if err := db.Save(&seller).Error; err != nil {
		t.Error(err)
	}

	// setting controller
	body, _ := json.Marshal(mock_seller_wrong_email)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/seller/login")

	LoginSeller(context)

	var response models.SellerResponse
	res_body := res.Body.String()

	json.Unmarshal([]byte(res_body), &response)

	t.Run("POST /seller/login", func(t *testing.T) {
		assert.Equal(t, false, response.Status)
		assert.Equal(t, "User Unauthorized. Email or Password not equal", response.Message)
	})
	db.Migrator().DropTable(&models.Seller{})
}

func TestGetDetailSeller(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}

	// cleaning data before testing
	db.Migrator().DropTable(&models.Seller{})
	db.AutoMigrate(&models.Seller{})

	// preparate dummy data
	seller := models.Seller{
		Name:     "Riska",
		Email:    "riska@gmail.com",
		Password: "123",
		Address:  "Jl Bogor Raya",
		Gender:   "F",
	}
	convert_pwd := []byte(seller.Password) //convert pass from string to byte
	hashed_pwd := EncryptPwd(convert_pwd)
	seller.Password = hashed_pwd
	if err := db.Save(&seller).Error; err != nil {
		t.Error(err)
	}

	//Make Token
	var user models.Seller
	if err := config.DB.Where("email = ?", seller.Email).First(&user).Error; err != nil {
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
	context.SetPath("/seller/:seller_id")
	context.SetParamNames("seller_id")
	context.SetParamValues("1")
	middleware.JWT([]byte(constants.SECRET_JWT))(GetDetailSellerTesting())(context)

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

	t.Run("PUT /seller/:seller_id", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, "riska@gmail.com", response.Email)
	})
	db.Migrator().DropTable(&models.Seller{})
}

func TestUpdateSeller(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}

	// cleaning data before testing
	db.Migrator().DropTable(&models.Seller{})
	db.AutoMigrate(&models.Seller{})

	// preparate dummy data
	seller := models.Seller{
		Name:     "Riska",
		Email:    "riska@gmail.com",
		Password: "123",
		Address:  "Jl Bogor Raya",
		Gender:   "F",
	}
	convert_pwd := []byte(seller.Password) //convert pass from string to byte
	hashed_pwd := EncryptPwd(convert_pwd)
	seller.Password = hashed_pwd
	if err := db.Save(&seller).Error; err != nil {
		t.Error(err)
	}

	//Make Token
	var user models.Seller
	if err := config.DB.Where("email = ?", seller.Email).First(&user).Error; err != nil {
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
	context.SetPath("/seller/:seller_id")
	context.SetParamNames("seller_id")
	context.SetParamValues("1")
	middleware.JWT([]byte(constants.SECRET_JWT))(UpdateSellerTesting())(context)

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

	t.Run("PUT /seller/:seller_id", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, "riska@gmail.com", response.Email)
	})
	db.Migrator().DropTable(&models.Seller{})
}

func TestLogoutSeller(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}

	// cleaning data before testing
	db.Migrator().DropTable(&models.Seller{})
	db.AutoMigrate(&models.Seller{})

	// preparate dummy data
	seller := models.Seller{
		Name:     "Riska",
		Email:    "riska@gmail.com",
		Password: "123",
		Address:  "Jl Bogor Raya",
		Gender:   "F",
	}
	convert_pwd := []byte(seller.Password) //convert pass from string to byte
	hashed_pwd := EncryptPwd(convert_pwd)
	seller.Password = hashed_pwd
	if err := db.Save(&seller).Error; err != nil {
		t.Error(err)
	}

	//Make Token
	var user models.Seller
	if err := config.DB.Where("email = ?", seller.Email).First(&user).Error; err != nil {
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
	context.SetPath("/seller/logout/:seller_id")
	context.SetParamNames("seller_id")
	context.SetParamValues("1")
	middleware.JWT([]byte(constants.SECRET_JWT))(LogoutSellerTesting())(context)

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

	t.Run("PUT /seller/:seller_id", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, "", response.Token)
	})
	db.Migrator().DropTable(&models.Seller{})
}

func TestEditStatusItemOrder(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}

	// cleaning data before testing
	db.Migrator().DropTable(&models.Seller{})
	db.AutoMigrate(&models.Seller{})
	db.Migrator().DropTable(&models.Order{})
	db.AutoMigrate(&models.Order{})
	db.Migrator().DropTable(&models.Product{})
	db.AutoMigrate(&models.Product{})

	// preparate dummy data
	seller := models.Seller{
		Name:     "Riska",
		Email:    "riska@gmail.com",
		Password: "123",
		Address:  "Jl Bogor Raya",
		Gender:   "F",
		Hp:       628123828282,
	}
	convert_pwd := []byte(seller.Password) //convert pass from string to byte
	hashed_pwd := EncryptPwd(convert_pwd)
	seller.Password = hashed_pwd
	if err := db.Save(&seller).Error; err != nil {
		t.Error(err)
	}

	product := models.Product{
		Name:        "apel",
		Price:       3000,
		Stock:       20,
		Description: "per ons",
		SellerID:    1,
	}
	if err := db.Save(&product).Error; err != nil {
		t.Error(err)
	}

	order := models.Order{
		Qty:        3,
		Price:      4000,
		CheckoutID: 1,
		ProductID:  1,
		Status:     "progress",
	}
	if err := db.Save(&order).Error; err != nil {
		t.Error(err)
	}

	//Make Token
	var user models.Seller
	if err := config.DB.Where("email = ?", seller.Email).First(&user).Error; err != nil {
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
	context.SetPath("/seller/orderlist/:order_id")
	context.SetParamNames("order_id")
	context.SetParamValues("1")
	middleware.JWT([]byte(constants.SECRET_JWT))(EditStatusItemOrderTesting())(context)

	// Unmarshal respose string to struct
	type Response struct {
		ID         uint   `json:"id"`
		CheckoutID uint   `json:"checkout_id"`
		ProductID  uint   `json:"product_id"`
		Qty        int    `json:"qty"`
		Price      int    `json:"price"`
		Status     string `json:"status"`
	}

	var response Response
	res_body := res.Body.String()

	json.Unmarshal([]byte(res_body), &response)

	t.Run("PUT /seller/orderlist/:order_id", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, "completed", response.Status)
	})
	db.Migrator().DropTable(&models.Seller{})
}

func TestGetAllOrders(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}

	// cleaning data before testing
	db.Migrator().DropTable(&models.Seller{})
	db.AutoMigrate(&models.Seller{})
	db.Migrator().DropTable(&models.Product{})
	db.AutoMigrate(&models.Product{})
	db.Migrator().DropTable(&models.Order{})
	db.AutoMigrate(&models.Order{})
	db.Migrator().DropTable(&models.Checkout{})
	db.AutoMigrate(&models.Checkout{})
	db.Migrator().DropTable(&models.Delivery{})
	db.AutoMigrate(&models.Delivery{})
	db.Migrator().DropTable(&models.Driver{})
	db.AutoMigrate(&models.Driver{})

	// preparate dummy data
	seller := models.Seller{
		Name:     "Bambang",
		Email:    "bambang@123",
		Password: "123",
		Address:  "Jl Bogor Raya",
		Gender:   "F",
	}
	convert_pwd := []byte(seller.Password) //convert pass from string to byte
	hashed_pwd := EncryptPwd(convert_pwd)
	seller.Password = hashed_pwd
	if err := db.Save(&seller).Error; err != nil {
		t.Error(err)
	}

	driver := models.Driver{
		Name:     "Budi",
		Email:    "budi@123",
		Password: "123",
		Address:  "Jl Bogor Raya",
		Gender:   "F",
	}
	convert_pwd_driver := []byte(driver.Password) //convert pass from string to byte
	hashed_pwd_driver := EncryptPwd(convert_pwd_driver)
	driver.Password = hashed_pwd_driver
	if err := db.Save(&driver).Error; err != nil {
		t.Error(err)
	}

	product := models.Product{
		Name:        "apel",
		Price:       3000,
		Stock:       20,
		Description: "per ons",
		SellerID:    1,
	}
	if err := db.Save(&product).Error; err != nil {
		t.Error(err)
	}

	checkout := models.Checkout{
		TotalQty:   2,
		TotalPrice: 6000,
		CustomerID: 1,
		PaymentID:  1,
		Status:     "searching",
	}
	if err := db.Save(&checkout).Error; err != nil {
		t.Error(err)
	}

	order := models.Order{
		Qty:        2,
		Price:      6000,
		CheckoutID: 1,
		ProductID:  1,
		Status:     "progress",
	}
	if err := db.Save(&order).Error; err != nil {
		t.Error(err)
	}

	delivery := models.Delivery{
		DriverID:   1,
		CheckoutID: 1,
		Status:     "progress",
	}
	if err := db.Save(&delivery).Error; err != nil {
		t.Error(err)
	}

	//Make Token
	var user models.Seller
	if err := config.DB.Where("email = ?", seller.Email).First(&user).Error; err != nil {
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
	context.SetPath("/seller/orderlist")
	middleware.JWT([]byte(constants.SECRET_JWT))(GetAllOrdersTesting())(context)

	// Unmarshal respose string to struct
	type Response struct {
		DriverName  string `json:"driver_name"`
		ProductName string `json:"product_name"`
		Qty         int    `json:"qty"`
		Price       int    `json:"price"`
	}

	var response []Response
	res_body := res.Body.String()

	json.Unmarshal([]byte(res_body), &response)

	t.Run("GET /seller/orderlist", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, "Budi", response[0].DriverName)
		assert.Equal(t, "apel", response[0].ProductName)
		assert.Equal(t, 2, response[0].Qty)
		assert.Equal(t, 6000, response[0].Price)
	})
	db.Migrator().DropTable(&models.Seller{})
	db.Migrator().DropTable(&models.Product{})
	db.Migrator().DropTable(&models.Order{})
	db.Migrator().DropTable(&models.Checkout{})
	db.Migrator().DropTable(&models.Delivery{})
	db.Migrator().DropTable(&models.Driver{})
}
