package database

import (
	"etmarket/project/config"
	"etmarket/project/models"
)

func GetOrdertoTake() ([]models.Checkout, error) {
	var checkouts []models.Checkout
	if err := config.DB.Find(&checkouts, "status=?", "searching").Error; err != nil {
		return checkouts, err
	}
	return checkouts, nil
}

func MakeDelivery(driver_id, checkout_id int) (models.Delivery, error) {
	var checkout models.Checkout
	var driver models.Driver
	var delivery models.Delivery
	if err := config.DB.Find(&checkout, "id=?", checkout_id).Error; err != nil {
		return delivery, err
	}
	if err := config.DB.Find(&driver, "id=?", driver_id).Error; err != nil {
		return delivery, err
	}
	delivery = models.Delivery{
		DriverID:   uint(driver_id),
		CheckoutID: uint(checkout_id),
		Status:     "progress",
	}
	if err := config.DB.Save(&delivery).Error; err != nil {
		return delivery, err
	}
	checkout.Status = "progress"
	if err := config.DB.Save(checkout).Error; err != nil {
		return delivery, err
	}
	return delivery, nil
}

func EditDelivery(driver_id, checkout_id int) (models.Delivery, error) {
	var delivery models.Delivery
	if err := config.DB.Find(&delivery, "driver_id=? AND checkout_id=?", driver_id, checkout_id).Error; err != nil {
		return delivery, err
	}
	delivery.Status = "completed"
	if err := config.DB.Save(delivery).Error; err != nil {
		return delivery, err
	}
	return delivery, nil
}

// func GetSelectedOrder(checkout_id int) (interface{}, error) {
// 	var order []models.Order
// 	if err := config.DB.Where(&order, "checkout_id = ?", checkout_id).Find(&order).Error; err != nil {
// 		return nil, err
// 	}
// 	return order, nil
// }

/*
Author: Patmiza
*/
func GetAllCompletedDeliveries(driver_id int) (interface{}, error) {
	var deliveries []models.Delivery
	if err := config.DB.Find(&deliveries, "driver_id=? AND status= ?", driver_id, "completed").Error; err != nil {
		return nil, err
	}
	return deliveries, nil
}
