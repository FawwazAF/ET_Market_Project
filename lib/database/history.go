package database

import (
	"etmarket/project/config"
	"etmarket/project/models"
)

func GetHistoryInProgress(status string) (interface{}, error) {
	var checkout []models.Checkout
	var orders models.Order
	if err := config.DB.Where(&orders, "status = ?", status).Find(&checkout).Error; err != nil {
		return nil, err
	}
	return checkout, nil
}

func GetHistoryCompleted(status string) (interface{}, error) {
	var checkout []models.Checkout
	var orders models.Order
	if err := config.DB.Where(&orders, "status = ?", status).Find(&checkout).Error; err != nil {
		return nil, err
	}
	return checkout, nil
}
