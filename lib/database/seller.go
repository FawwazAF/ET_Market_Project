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

func CheckEmailOnSeller(email string) bool {
	var check bool
	var seller models.Seller

	config.DB.Model(&seller).Select("email").Where("email=?", email).First(&seller.Email)

	if seller.Email == "" {
		check = false
	} else {
		check = true
	}

	return check
}

func GetAllSellerProduct(seller_id int) (interface{}, error) {
	var product []models.Product
	if err := config.DB.Find(&product, "SellerID = ?", seller_id).Error; err != nil {
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
	if err := config.DB.Where(&product, "id = ? AND SellerID = ?", product_id, seller_id).Save(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}
