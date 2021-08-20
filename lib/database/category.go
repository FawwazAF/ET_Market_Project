package database

import (
	"etmarket/project/config"
	"etmarket/project/models"
)

func GetAllCategories() (interface{}, error) {
	var categories []models.Category
	if err := config.DB.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func GetAllCategoriesMarketId(market_id int) (interface{}, error) {
	var seller []models.Seller
	if err := config.DB.Find(&seller, "market_id=?", market_id).Error; err != nil {
		return nil, err
	}
	return seller, nil
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

func GetCategoryNameMarketId(market_id int, category_name string) (interface{}, error) {
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
