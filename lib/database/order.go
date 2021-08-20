package database

import (
	"etmarket/project/config"
	"etmarket/project/models"
)

/*
Author: Patmiza
Getting all categories of seller in a market
*/
func GetAllProgressOrders(driver_id int) (interface{}, error) {
	var orders []models.Order
	if err := config.DB.Find(&orders, "driver_id = ? AND status = ?", driver_id, "progress").Error; err != nil {
		return nil, err
	}
	return orders, nil
}
