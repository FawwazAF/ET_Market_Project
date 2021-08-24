package database

import (
	"etmarket/project/config"
	"etmarket/project/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mockDBProduct = models.Product{
		Name:        "tomat",
		Price:       2000,
		Stock:       20,
		Description: "ini tomat",
		SellerID:    1,
	}
	mockDBProduct2 = models.Product{
		Name:        "sawi",
		Price:       2500,
		Stock:       10,
		Description: "ini sawi",
		SellerID:    1,
	}
	mockDBProduct3 = models.Product{
		Name:        "teh",
		Price:       5000,
		Stock:       15,
		Description: "ini teh",
		SellerID:    2,
	}
)

func TestGetAllProductByShopId(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator(),DropTable(&models.Product{})
	config.DB.Migrator().AutoMigrate(&models.Product{})
	config.DB.Save(&mockDBProduct)
	config.DB.Save(&mockDBProduct2)
	config.DB.Save(&mockDBProduct3)
	gotProduct, err := GetAllSellerProduct(1)
	if assert.NoError(t, err) {
		assert.Equal(t, "tomat", [0]gotProduct.name)
		assert.Equal(t, 2000, [0]gotProduct.Price)
		assert.Equal(t, 20, [0]gotProduct.Stock)
		assert.Equal(t, "sawi", [1]gotProduct.name)
		assert.Equal(t, 2500, [1]gotProduct.Price)
		assert.Equal(t, 10, [1]gotProduct.Stock)
	}
}

func TestAddProductToSeller(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator(),DropTable(&models.Product{})
	config.DB.Migrator().AutoMigrate(&models.Product{})
	var product []models.Product
	AddProductToSeller(mockDBProduct)
	AddProductToSeller(mockDBProduct2)
	gotProduct, err := config.DB.Find(&product)
	if assert.NoError(t, err) {
		assert.Equal(t, "tomat", [0]gotProduct.name)
		assert.Equal(t, 2000, [0]gotProduct.Price)
		assert.Equal(t, 20, [0]gotProduct.Stock)
		assert.Equal(t, "sawi", [1]gotProduct.name)
		assert.Equal(t, 2500, [1]gotProduct.Price)
		assert.Equal(t, 10, [1]gotProduct.Stock)
	}
}

func TestEditSellerProduct(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator(),DropTable(&models.Product{})
	config.DB.Migrator().AutoMigrate(&models.Product{})
	var product []models.Product
	config.DB.Save(&mockDBProduct2)
	EditSellerProduct(mockDBProduct)
	gotProduct, err := GetEditProduct(1, 1)
	if assert.NoError(t, err) {
		assert.Equal(t, "tomat", [0]gotProduct.name)
		assert.Equal(t, 2000, [0]gotProduct.Price)
		assert.Equal(t, 20, [0]gotProduct.Stock)
	}
}
	mock_db_seller = models.Seller{
		Name:       "Riska",
		Email:      "riska@gmail.com",
		Password:   "123",
		Address:    "bogor",
		Gender:     "F",
		Hp:         62813628282,
		MarketID:   2,
		CategoryID: 1,
	}
	mock_db_seller_empty = models.Seller{
		Name:       "",
		Email:      "",
		Password:   "",
		Address:    "",
		Gender:     "",
		Hp:         62813628282,
		MarketID:   0,
		CategoryID: 0,
	}
	mock_db_seller_edit = models.Seller{
		Name:       "Riska Kurnia",
		Email:      "riskakurnia@gmail.com",
		Password:   "123",
		Address:    "bogor",
		Gender:     "F",
		Hp:         62813628282,
		MarketID:   2,
		CategoryID: 1,
	}
	mock_db_product = models.Product{
		Name:        "apel",
		Price:       20000,
		Stock:       20,
		Description: "per kg",
		SellerID:    1,
	}
	mock_db_product_edit = models.Product{
		Name:        "apel malang",
		Price:       25000,
		Stock:       20,
		Description: "per kg",
		SellerID:    1,
	}
)

func TestCreateSellerSuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Seller{})
	config.DB.Migrator().AutoMigrate(&models.Seller{})
	seller, err := CreateSeller(mock_db_seller)
	if assert.NoError(t, err) {
		assert.Equal(t, "Riska", seller.Name)
		assert.Equal(t, "riska@gmail.com", seller.Email)
		assert.Equal(t, "123", seller.Password)
	}
}

func TestCreateSellerError(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Seller{})
	_, err := CreateSeller(mock_db_seller)
	assert.Error(t, err)
}

func TestLoginSellerSuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Seller{})
	config.DB.Migrator().AutoMigrate(&models.Seller{})
	seller, err := CreateSeller(mock_db_seller)
	seller_login, err := LoginSeller(seller.Email)
	if assert.NoError(t, err) {
		assert.Equal(t, "Riska", seller_login.Name)
		assert.Equal(t, "riska@gmail.com", seller_login.Email)
		assert.Equal(t, "123", seller_login.Password)
	}
}

func TestLoginSellerWrongEmail(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Seller{})
	config.DB.Migrator().AutoMigrate(&models.Seller{})
	_, err := CreateSeller(mock_db_seller)
	_, err = LoginSeller("rizka@gmail.com")
	assert.Error(t, err)
}

func TestGetSellerSuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Seller{})
	config.DB.Migrator().AutoMigrate(&models.Seller{})
	created_seller, _ := CreateSeller(mock_db_seller)
	seller, err := GetSeller(int(created_seller.ID))
	if assert.NoError(t, err) {
		assert.Equal(t, "Riska", seller.Name)
		assert.Equal(t, "riska@gmail.com", seller.Email)
		assert.Equal(t, "123", seller.Password)
	}
}

func TestGetSellerError(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Seller{})
	CreateSeller(mock_db_seller)
	_, err := GetSeller(1)
	assert.Error(t, err)
}

func TestEditSellerSuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Seller{})
	config.DB.Migrator().AutoMigrate(&models.Seller{})
	created_seller, _ := CreateSeller(mock_db_seller)
	seller, err := GetSeller(int(created_seller.ID))
	seller.Name = "Riska Kurnia"
	seller.Email = "riskakurnia@gmail.com"
	seller.Password = "123"
	edit_seller, err := UpdateSeller(seller)
	seller_edited, err := GetSeller(int(edit_seller.ID))
	if assert.NoError(t, err) {
		assert.Equal(t, "Riska Kurnia", seller_edited.Name)
		assert.Equal(t, "riskakurnia@gmail.com", seller_edited.Email)
		assert.Equal(t, "123", seller_edited.Password)
	}
}

func TestEditSellerError(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Seller{})
	_, err := UpdateSeller(mock_db_seller_edit)
	assert.Error(t, err)
}

func TestGetPwdSellerSuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Seller{})
	config.DB.Migrator().AutoMigrate(&models.Seller{})
	created_seller, _ := CreateSeller(mock_db_seller)
	seller := GetPwdSeller(created_seller.Email)
	assert.Equal(t, seller, created_seller.Password)
}

func TestCheckEmailOnSellerSuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Seller{})
	config.DB.Migrator().AutoMigrate(&models.Seller{})
	created_seller, _ := CreateSeller(mock_db_seller)
	seller, err := CheckEmailOnSeller(created_seller.Email)
	if assert.NoError(t, err) {
		assert.Equal(t, seller, created_seller.Email)
	}
}

func TestCheckEmailOnSellerError(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Seller{})
	CreateSeller(mock_db_seller)
	_, err := CheckEmailOnSeller("riska@gmail.com")
	assert.Error(t, err)
}

func TestGetEmailSellerByIdSuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Seller{})
	config.DB.Migrator().AutoMigrate(&models.Seller{})
	created_seller, _ := CreateSeller(mock_db_seller)
	seller, err := GetEmailSellerById(int(created_seller.ID))
	if assert.NoError(t, err) {
		assert.Equal(t, seller, created_seller.Email)
	}
}

func TestGetEmailSellerByIdError(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Seller{})
	CreateSeller(mock_db_seller)
	_, err := GetEmailSellerById(1)
	assert.Error(t, err)
}

func TestGetHPSeller(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Seller{})
	config.DB.Migrator().AutoMigrate(&models.Seller{})
	created_seller, _ := CreateSeller(mock_db_seller)
	hp, err := GetHPSeller(1, created_seller.ID)
	if assert.NoError(t, err) {
		assert.Equal(t, hp, created_seller.Hp)
	}
}

func TestGetHPSellerError(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Seller{})
	CreateSeller(mock_db_seller)
	_, err := GetHPSeller(1, 1)
	assert.Error(t, err)
}

func TestGetSellerByCheckoutIdSuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Seller{})
	config.DB.Migrator().AutoMigrate(&models.Seller{})
	created_seller, _ := CreateSeller(mock_db_seller)
	seller := GetSellerByCheckoutId(1)
	assert.Equal(t, seller.Name, created_seller.Name)
	assert.Equal(t, seller.Email, created_seller.Email)
	assert.Equal(t, seller.Password, created_seller.Password)
}

func TestGetSellerIdByOderIdSuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Seller{})
	config.DB.Migrator().AutoMigrate(&models.Seller{})
	created_seller, _ := CreateSeller(mock_db_seller)
	seller, err := GetSellerIdByOderId(1)
	if assert.NoError(t, err) {
		assert.Equal(t, seller, created_seller.ID)
	}
}

func TestGetSellerIdByOderIdError(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Seller{})
	CreateSeller(mock_db_seller)
	_, err := GetSellerIdByOderId(1)
	assert.Error(t, err)
}

func TestGetSellerbyNameSuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Seller{})
	config.DB.Migrator().AutoMigrate(&models.Seller{})
	created_seller, _ := CreateSeller(mock_db_seller)
	seller, err := GetSellerbyName(1, "sayur-sayuran")
	if assert.NoError(t, err) {
		assert.Equal(t, seller, created_seller.ID)
	}
}

func TestGetSellerbyNameError(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Seller{})
	CreateSeller(mock_db_seller)
	_, err := GetSellerbyName(1, "sayur-sayuran")
	assert.Error(t, err)
}

func TestEditSellerProductSuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Product{})
	config.DB.Migrator().AutoMigrate(&models.Product{})
	config.DB.Migrator().DropTable(&models.Seller{})
	config.DB.Migrator().AutoMigrate(&models.Seller{})
	created_seller, _ := CreateSeller(mock_db_seller)
	created_product, _ := AddProductToSeller(mock_db_product)
	product, err := GetEditProduct(int(created_product.ID), int(created_seller.ID))
	product.Name = "apel malang"
	product.Price = 25000
	product.Stock = 20
	edit_product, err := EditSellerProduct(product)
	if assert.NoError(t, err) {
		assert.Equal(t, "apel malang", edit_product.Name)
		assert.Equal(t, 25000, edit_product.Price)
		assert.Equal(t, 20, edit_product.Stock)
	}
}
