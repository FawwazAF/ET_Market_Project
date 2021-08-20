package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	CustomerID int `json:"customer_id" form:"customer_id"`
	ProductID  int `json:"product_id" form:"product_id"`
	Price      int `json:"price" form:"price"`
	Qty        int `json:"qty" form:"qty"`
}
