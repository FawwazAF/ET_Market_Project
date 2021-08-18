package models

import "gorm.io/gorm"

type Seller struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Address  string `json:"address" form:"address"`
	Gender   string `json:"gender" form:"gender" gorm:"type:enum('M', 'F')"`
	Token    string `json :"token" form:"token"`

	// 1 to many with category seller and market
	MarketID         uint `json:"market_id" form:"market_id"`
	CategorySellerID uint `json:"category_seller_id" form:"category_seller_id"`
}
