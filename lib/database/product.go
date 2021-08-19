package database

import (
	"etmarket/project/config"
	"etmarket/project/models"
)

func GetAllProductByShopId(seller_id int) (interface{}, error) {
	var product []models.Product
	if err := config.DB.Find(&product, "seller_id = ?", seller_id).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func GetSpecificProductByShopId(seller_id int, product_name string) (interface{}, error) {
	var product []models.Product
	search := "%" + product_name
	if err := config.DB.Find(&product, "seller_id = ? AND name LIKE ?", seller_id, search).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func GetSpecificProductById(seller_id int, product_id int) (interface{}, error) {
	var product []models.Product
	if err := config.DB.Find(&product, "seller_id = ? AND id = ?", seller_id, product_id).Error; err != nil {
		return nil, err
	}
	return product, nil
}
