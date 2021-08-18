package models

import "gorm.io/gorm"

type Market struct {
	gorm.Model
	Name    string   `json:"name" form:"name"`
	Address string   `json:"address" form:"address"`
	Seller  []Seller `gorm:"foreignKey:MarketID;references:ID;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
}
