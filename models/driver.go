package models

import "gorm.io/gorm"

type Driver struct {
	gorm.Model
	ID       uint   `gorm:"primarykey"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Address  string `json:"address" form:"address"`
	Gender   string `json:"gender" form:"gender" gorm:"type:enum('M', 'F')"`
	Token    string `json:"token" form:"token"`
}

type ShowDriver struct {
	ID   int
	Name string
}
