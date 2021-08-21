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

type DriverLogin struct {
	ID      uint   `json:"id" gorm:"primarykey"`
	Name    string `json:"name"`
	Email   string `gorm:"UNIQUE" json:"email"`
	Address string `json:"address" form:"address"`
	Gender  string `json:"gender" form:"gender" gorm:"type:enum('M', 'F')"`
	Token   string `json:"token" form:"token"`
}

type DriverResponse struct {
	Status  bool     `json:"status"`
	Message string   `json:"message"`
	Data    []Driver `json:"data"`
}

type DriverResponseSingle struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    Driver `json:"data"`
}

type DriverResponseLogin struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    DriverLogin `json:"data"`
}
