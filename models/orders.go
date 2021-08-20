package models

import "gorm.io/gorm"

type Orders struct {
	gorm.Model
	DriverID   uint `json:"driver_id" form:"driver_id"`
	CheckoutID uint `json:"checkout_id" form:"checkout_id"`
	Status     int  `json:"status" form:"status" gorm:"type:enum('in_progress', 'completed')"`
}
