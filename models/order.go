package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Qty        int  `json:"qty" form:"qty"`
	Price      int  `json:"price" form:"price"`
	CheckoutID uint `json:"checkout_id" form:"checkout_id"`
	ProductID  uint `json:"product_id" form:"product_id"`
}
