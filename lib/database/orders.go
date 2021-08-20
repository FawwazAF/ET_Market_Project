package database

import (
	"etmarket/project/config"
	"etmarket/project/models"
)

func GetOrdertoTake() (interface{}, error) {
	var checkouts []models.Checkout
	if err := config.DB.Find(&checkouts, "status=?", "searching").Error; err != nil {
		return nil, err
	}
	return checkouts, nil
}

func MakeDelivery(driver_id, checkout_id int) (interface{}, error) {
	var checkout models.Checkout
	var driver models.Driver
	if err := config.DB.Find(&checkout, "id=?", checkout_id).Error; err != nil {
		return nil, err
	}
	if err := config.DB.Find(&driver, "id=?", driver_id).Error; err != nil {
		return nil, err
	}
	delivery := models.Delivery{
		DriverID:   uint(driver_id),
		CheckoutID: uint(checkout_id),
		Status:     "progress",
	}
	if err := config.DB.Save(&delivery).Error; err != nil {
		return nil, err
	}
	return delivery, nil
}

func EditDelivery(driver_id, checkout_id int) (interface{}, error) {
	var delivery models.Delivery
	if err := config.DB.Find(&delivery, "driver_id=? AND checkout_id=?", driver_id, checkout_id).Error; err != nil {
		return nil, err
	}
	delivery.Status = "completed"
	if err := config.DB.Save(delivery).Error; err != nil {
		return nil, err
	}
	return delivery, nil
}
