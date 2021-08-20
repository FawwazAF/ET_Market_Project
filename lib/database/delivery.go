package database

import (
	"etmarket/project/config"
	"etmarket/project/models"
)

/*
Author: Patmiza
*/
func GetAllCompletedDeliveries(driver_id int) (interface{}, error) {
	var deliveries []models.Delivery
	if err := config.DB.Find(&deliveries, "driver_id= AND status= ?", driver_id, "completed").Error; err != nil {
		return nil, err
	}
	return deliveries, nil
}
