package models

import "gorm.io/gorm"

type Seller struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Address  string `json:"address" form:"address"`
	Gender   string `json:"gender" form:"gender" gorm:"type:enum('M', 'F')"`
	Hp       int64  `json:"hp" form:"hp"`
	Token    string `json:"token" form:"token"`

	// 1 to many with category seller and market
	MarketID   uint `json:"market_id" form:"market_id"`
	CategoryID uint `json:"category_id" form:"category_id"`

	Product []Product `gorm:"foreignKey:SellerID;references:ID;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
}

type SellerLogin struct {
	ID         uint   `json:"id" gorm:"primarykey"`
	Name       string `json:"name"`
	Email      string `gorm:"UNIQUE" json:"email"`
	Address    string `json:"address" form:"address"`
	Gender     string `json:"gender" form:"gender" gorm:"type:enum('M', 'F')"`
	Token      string `json:"token" form:"token"`
	MarketID   uint   `json:"market_id" form:"market_id"`
	CategoryID uint   `json:"category_id" form:"category_id"`
}

type SellerResponse struct {
	Status  bool     `json:"status"`
	Message string   `json:"message"`
	Data    []Seller `json:"data"`
}

type SellerResponseSingle struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    Seller `json:"data"`
}

type SellerResponseLogin struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    SellerLogin `json:"data"`
}
