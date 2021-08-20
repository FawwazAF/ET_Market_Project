package database

import (
	"etmarket/project/config"
	"etmarket/project/models"
)

func GetHistoryInProgress(status string, checkout_id int) (interface{}, error) {
	var checkout []models.Checkout
	var orders models.Order
	if err := config.DB.Where(&orders, "status = ? AND checkout_id = ?", status, checkout_id).Find(&checkout).Error; err != nil {
		return nil, err
	}
	return checkout, nil
}

func GetHistoryCompleted(status string, checkout_id int) (interface{}, error) {
	var checkout []models.Checkout
	var orders models.Order
	if err := config.DB.Where(&orders, "status = ? AND checkout_id = ?", status, checkout_id).Find(&checkout).Error; err != nil {
		return nil, err
	}
	return checkout, nil
}
