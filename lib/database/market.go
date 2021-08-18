package database

import (
	"etmarket/project/config"
	"etmarket/project/models"
)

func GetManyMarkets() (interface{}, error) {
	var markets []models.Market
	if err := config.DB.Find(&markets).Error; err != nil {
		return nil, err
	}
	return markets, nil
}
