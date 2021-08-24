package controller

import (
	"bytes"
	"encoding/json"
	"etmarket/project/config"
	"etmarket/project/middlewares"
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

func TestGetSellerProducts(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}
	// cleaning data before testing
	db.AutoMigrate(&models.Seller{})
	db.AutoMigrate(&models.Product{})

	// preparate dummy data
	//Checkout dummy
	var newProduct models.Product
	newProduct.Name = "tomat"
	newProduct.Price = 1000
	newProduct.Stock = 10
	if err := db.Save(&newProduct).Error; err != nil {
		t.Error(err)
	}
	//Seller dummy
	var newSeller models.Seller
	newSeller.Name = "jojo"
	newSeller.Email = "jojo@123"
	newSeller.Password = "jj123"
	newSeller.Address = "bandung"
	newSeller.Gender = "M"
	if err := db.Save(&newSeller).Error; err != nil {
		t.Error(err)
	}
	//Make Token
	var seller models.Seller
	if err := config.DB.Where("email = ?", newSeller.Email).First(&seller).Error; err != nil {
		t.Error(err)
	}

	token, _ := middlewares.CreateToken(int(seller.ID))
	// setting controller
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/seller/products")

	GetSellerProducts(context)

	type Response struct {
		Name        string `json:"name"`
		Price       int    `json:"price"`
		Stock       int    `json:"stock"`
		Description string `json:"description"`
		SellerID    uint   `json:"seller_id"`
	}
	// var Result models.Product
	// var show_list []Result
	// for i := 0; i < len(list_product); i++ {
	// 	new_array := Result{
	// 		Name:  list_product[i].Name,
	// 		Price: list_product[i].Price,
	// 		Stock: list_product[i].Stock,
	// 	}
	// 	show_list = append(show_list, new_array)
	// }

	var response Response
	resBody2 := res.Body.String()

	json.Unmarshal([]byte(resBody2), &response)

	t.Run("GET /seller/products", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, "tomat", response.Name)
		assert.Equal(t, 1000, response.Price)
		assert.Equal(t, 10, response.Stock)
	})

	db.Migrator().DropTable(&models.Seller{})
	db.Migrator().DropTable(&models.Product{})
}

func TestAddProductToSeller(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}
	// cleaning data before testing
	db.AutoMigrate(&models.Seller{})
	db.AutoMigrate(&models.Product{})

	//Seller dummy
	var newSeller models.Seller
	newSeller.Name = "jojo"
	newSeller.Email = "jojo@123"
	newSeller.Password = "jj123"
	newSeller.Address = "bandung"
	newSeller.Gender = "M"
	if err := db.Save(&newSeller).Error; err != nil {
		t.Error(err)
	}
	//Make Token
	var seller models.Seller
	if err := config.DB.Where("email = ?", newSeller.Email).First(&seller).Error; err != nil {
		t.Error(err)
	}

	token, _ := middlewares.CreateToken(int(seller.ID))
	// preparate dummy data
	//Checkout dummy
	mocknewProduct := models.Product{
		Name:        "tomat",
		Price:       1000,
		Stock:       10,
		Description: "per ons",
		SellerID:    seller.ID,
	}

	// setting controller
	body, _ := json.Marshal(mocknewProduct)
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/seller/products")

	AddProductToSeller(context)

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

	t.Run("POST /seller/products", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, "tomat", response.Name)
		assert.Equal(t, 1000, response.Price)
		assert.Equal(t, 10, response.Stock)
		assert.Equal(t, "per ons", response.Description)
		assert.Equal(t, 1, response.SellerID)
	})
	db.Migrator().DropTable(&models.Product{})
	db.Migrator().DropTable(&models.Seller{})
}

func TestEditSellerProduct(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}
	// cleaning data before testing
	db.AutoMigrate(&models.Seller{})
	db.AutoMigrate(&models.Product{})

	// preparate dummy data
	//Checkout dummy
	var newProduct models.Product
	newProduct.Name = "tomat"
	newProduct.Price = 1000
	newProduct.Stock = 10
	if err := db.Save(&newProduct).Error; err != nil {
		t.Error(err)
	}
	//Seller dummy
	var newSeller models.Seller
	newSeller.Name = "jojo"
	newSeller.Email = "jojo@123"
	newSeller.Password = "jj123"
	newSeller.Address = "bandung"
	newSeller.Gender = "M"
	if err := db.Save(&newSeller).Error; err != nil {
		t.Error(err)
	}
	//Make Token
	var seller models.Seller
	if err := config.DB.Where("email = ?", newSeller.Email).First(&seller).Error; err != nil {
		t.Error(err)
	}

	mocknewProduct := models.Product{
		Name:        "kangkung",
		Price:       1500,
		Stock:       11,
		Description: "per ons",
		SellerID:    seller.ID,
	}
	token, _ := middlewares.CreateToken(int(seller.ID))
	body, _ := json.Marshal(mocknewProduct)
	// setting controller
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/seller/products")
	context.SetParamNames("product_id")
	context.SetParamValues("1")

	EditSellerProduct(context)

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

	t.Run("POST /seller/products", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, "kangkung", response.Name)
		assert.Equal(t, 1500, response.Price)
		assert.Equal(t, 11, response.Stock)
		assert.Equal(t, "per ons", response.Description)
		assert.Equal(t, 1, response.SellerID)
	})
	db.Migrator().DropTable(&models.Product{})
	db.Migrator().DropTable(&models.Seller{})
}
