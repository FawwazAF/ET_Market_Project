package database

import (
	"etmarket/project/config"
	"etmarket/project/models"
)

func GetOrdertoTake() (interface{}, error) {
	var checkouts []models.Checkout
	if err := config.DB.Find(&checkouts, "status=searching").Error; err != nil {
		return nil, err
	}
	return checkouts, nil
}
