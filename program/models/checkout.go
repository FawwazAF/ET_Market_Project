package models

import "gorm.io/gorm"

type Checkout struct {
	gorm.Model
	TotalQty   int    `json:"total_qty" form:"total_qty"`
	TotalPrice int    `json:"total_price" form:"total_price"`
	CustomerID uint   `json:"customer_id" form:"customer_id"`
	PaymentID  uint   `json:"payment_id" form:"payment_id"`
	Status     string `json:"status" form:"status" gorm:"type:enum('searching', 'delivery', 'completed')"`
}
