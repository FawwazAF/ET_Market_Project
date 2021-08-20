package models

import "gorm.io/gorm"

type Checkout struct {
	gorm.Model
	CustomerID    uint   `json:"customer_id" form:"customer_id"`
	PaymentMethod uint   `json:"payment_method" form:"payment_method"`
	TotalQty      int    `json:"total_quantity" form:"total_quantity"`
	TotalPrice    int    `json:"total_price" form:"total_price"`
	Status        string `json:"status" form:"status" gorm:"type:enum('searching', 'delivery', 'completed')"`
}
