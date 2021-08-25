package database

import (
	"etmarket/project/config"
	"etmarket/project/models"
)

func GetHistoryInProgress(status string, logged_in_user_id int) ([]models.Checkout, error) {
	var checkout []models.Checkout
	if err := config.DB.Find(&checkout, "status = ? AND customer_id = ?", status, logged_in_user_id).Error; err != nil {
		return nil, err
	}
	return checkout, nil
}

func GetHistoryCompleted(status string, logged_in_user_id int) ([]models.Checkout, error) {
	var checkout []models.Checkout
	var orders models.Order
	if err := config.DB.Find(&orders, "status = ? AND customer_id = ?", status, logged_in_user_id).Error; err != nil {
		return nil, err
	}
	return checkout, nil
}
