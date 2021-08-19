package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	Qty        int  `json:"qty" form:"qty"`
	Price      int  `json:"price" form:"price"`
	CustomerID uint `json:"customer_id" form:"customer_id"`
	ProductID  uint `json:"product_id" form:"product_id"`
}
