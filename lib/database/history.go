package database

import (
	"etmarket/project/config"
	"etmarket/project/models"
)

func GetHistoryInProgress(checkout_id int) (interface{}, error) {
	var checkout []models.Checkout
	var orders models.Orders
	status := "inprogress"
	if err := config.DB.Where(&orders, "checkout_id = ? AND status = ?", checkout_id, status).Find(&checkout).Error; err != nil {
		return nil, err
	}
	return checkout, nil
}

func GetHistoryCompleted(checkout_id int) (interface{}, error) {
	var checkout []models.Checkout
	var orders models.Orders
	status := "completed"
	if err := config.DB.Where(&orders, "checkout_id = ? AND status = ?", checkout_id, status).Find(&checkout).Error; err != nil {
		return nil, err
	}
	return checkout, nil
}
