package database

import (
	"etmarket/project/config"
	"etmarket/project/models"
)

func GetAllCategoriesMarketId(market_id int) (interface{}, error) {
	var categories []models.Category
	if err := config.DB.Find(&categories, "market_id=?", market_id).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func GetCategoryNameMarketId(market_id int, category_name string) (interface{}, error) {
	var categories []models.Category
	search := "%" + category_name + "%"
	if err := config.DB.Find(&categories, "market_id=?", market_id, "category_name LIKE ", search).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
