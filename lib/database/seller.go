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
