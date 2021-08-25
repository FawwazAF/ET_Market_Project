package database

import (
	"etmarket/project/config"
	"etmarket/project/models"
)

func GetHistory(status string, logged_in_user_id int) ([]models.Checkout, error) {
	var checkout []models.Checkout
	if err := config.DB.Find(&checkout, "status = ? AND customer_id = ?", status, logged_in_user_id).Error; err != nil {
		return nil, err
	}
	return checkout, nil
}
