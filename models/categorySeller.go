package models

import "gorm.io/gorm"

type CategorySeller struct {
	gorm.Model
	Name   string   `json:"name" form:"name"`
	Seller []Seller `gorm:"foreignKey:CategorySellerID;references:ID;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
}
