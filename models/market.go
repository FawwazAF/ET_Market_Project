package models

import "gorm.io/gorm"

type Market struct {
	gorm.Model
	Name    string `json:"name" form:"name"`
	Address string `json:"address" form:"address"`
}
