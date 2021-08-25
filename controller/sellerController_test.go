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
	context.SetPath("/seller/products/:product_id")
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

	t.Run("POST /seller/products/:product_id", func(t *testing.T) {
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
