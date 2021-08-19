package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `json:"name" form:"name"`
	Price       int    `json:"price" form:"price"`
	Quantity    int    `json:"quantity" form:"quantity"`
	Description string `json:"description" form:"address"`

	SellerID uint `json:"seller_id" form:"seller_id"`
}
