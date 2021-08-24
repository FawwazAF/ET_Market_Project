package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name   string   `json:"name" form:"name"`
	Seller []Seller `gorm:"foreignKey:CategoryID;references:ID;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
}
