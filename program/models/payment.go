package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	Name string `json:"name" form:"name"`
}
