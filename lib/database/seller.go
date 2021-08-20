package database

import (
	"etmarket/project/config"
	"etmarket/project/middlewares"
	"etmarket/project/models"
)

func CreateSeller(seller models.Seller) (interface{}, error) {
	if err := config.DB.Save(&seller).Error; err != nil {
		return nil, err
	}
	return seller, nil
}

//login seller with matching data from database
func LoginSeller(email string) (interface{}, error) {
	var seller models.Seller
	var err error
	if err = config.DB.Where("email = ?", email).First(&seller).Error; err != nil {
		return nil, err
	}
	seller.Token, err = middlewares.CreateToken(int(seller.ID))
	if err != nil {
		return nil, err
	}
	if err := config.DB.Save(seller).Error; err != nil {
		return nil, err
	}
	return seller, err
}

func CheckEmailOnSeller(email string) (interface{}, error) {
	var seller models.Seller

	if err := config.DB.Model(&seller).Select("email").Where("email=?", email).First(&seller.Email).Error; err != nil {
		return nil, err
	}

	return seller, nil
}

func GetSellerById(seller_id int) (interface{}, error) {
	var seller models.Seller
	if err := config.DB.Where("id=?", seller_id).First(&seller).Error; err != nil {
		return nil, err
	}
	return seller, nil
}

//get 1 specified seller with Seller struct output
func GetSeller(id int) (models.Seller, error) {
	var seller models.Seller
	if err := config.DB.Find(&seller, "id=?", id).Error; err != nil {
		return seller, err
	}
	return seller, nil
}

//get email seller
func GetEmailSellerById(seller_id int) (string, error) {
	var seller models.Seller

	if err := config.DB.Model(&seller).Select("email").Where("id=?", seller_id).First(&seller.Email).Error; err != nil {
		return "nil", err
	}

	return seller.Email, nil
}

//update seller info from database
func UpdateSeller(seller models.Seller) (interface{}, error) {
	if err := config.DB.Save(&seller).Error; err != nil {
		return nil, err
	}
	return seller, nil
}

func GetAllSellerProduct(seller_id int) (interface{}, error) {
	var product []models.Product
	if err := config.DB.Find(&product, "seller_id = ?", seller_id).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func AddProductToSeller(product models.Product) (models.Product, error) {
	if err := config.DB.Save(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func EditSellerProduct(product_id int, seller_id int, product models.Product) (models.Product, error) {
	if err := config.DB.Where(&product, "id = ? AND seller_id = ?", product_id, seller_id).Save(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func GetSellerbyName(market_id int, category_name string) (interface{}, error) {
	var categories models.Category
	var seller []models.Seller
	search := "%" + category_name + "%"
	if err := config.DB.Find(&categories, "name LIKE ?", search).Error; err != nil {
		return nil, err
	}
	if err := config.DB.Find(&seller, "market_id=? AND category_id=?", market_id, categories.ID).Error; err != nil {
		return nil, err
	}
	return seller, nil
}
