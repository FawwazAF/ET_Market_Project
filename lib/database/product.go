package database

import (
	"etmarket/project/config"
	"etmarket/project/models"
)

func GetAllProductByShopId(seller_id int) (interface{}, error) {
	var products []models.Product
	if err := config.DB.Find(&product, "seller_id=?", seller_id).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func GetSpecificProductByShopId(seller_id int, product_name string) (interface{}, error) {
	var product []models.Product
	if err := config.DB.Find(&product, "seller_id=?", seller_id, "product_name=?", product_name).Error; err != nil {
		return nil, err
	}
	return product, nil
}
