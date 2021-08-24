package database

import (
	"etmarket/project/config"
	"etmarket/project/models"
)

func GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	if err := config.DB.Find(&categories).Error; err != nil {
		return categories, err
	}
	return categories, nil
}

/*
Author: Patmiza
*/
func GetAllCategoriesMarketId(market_id int) (interface{}, error) {
	var seller []models.Seller
	if err := config.DB.Find(&seller, "market_id=?", market_id).Error; err != nil {
		return nil, err
	}
	return seller, nil
}

/*
Author: Patmiza
*/
// func GetCategoryNameMarketId(market_id int, category_name string) (interface{}, error) {
// 	var categories models.Category
// 	var seller []models.Seller
// 	search := "%" + category_name + "%"
// 	if err := config.DB.Find(&categories, "name LIKE ?", search).Error; err != nil {
// 		return nil, err
// 	}
// 	if err := config.DB.Find(&seller, "market_id=? AND category_id=?", market_id, categories.ID).Error; err != nil {
// 		return nil, err
// 	}
// 	return seller, nil
// }

type SellerByCategory struct {
	ID       uint
	Name     string
	Category string
}

func GetCategoryNameMarketId(market_id int, category_name string) ([]SellerByCategory, error) {
	search := "%" + category_name + "%"

	rows, err := config.DB.Model(&Result{}).Raw("SELECT sellers.id, sellers.name, categories.name AS category FROM sellers, categories, markets WHERE categories.id = sellers.category_id AND markets.id = sellers.market_id AND categories.name LIKE ? AND markets.id = ?", search, market_id).Rows()

	defer rows.Close()

	var result []SellerByCategory
	for rows.Next() {
		config.DB.ScanRows(rows, &result)
	}

	return result, err
}
