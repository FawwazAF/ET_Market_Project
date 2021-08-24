package models

import "gorm.io/gorm"

type Delivery struct {
	gorm.Model
	Status     string `json:"status" form:"status" gorm:"type:enum('progress', 'completed')"`
	DriverID   uint   `json:"driver_id" form:"driver_id"`
	CheckoutID uint   `json:"checkout_id" form:"checkout_id"`
}
