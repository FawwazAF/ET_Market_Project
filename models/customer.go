package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Address  string `json:"address" form:"address"`
	Gender   string `json:"gender" form:"gender" gorm:"type:enum('M', 'F')"`
	Token    string `json:"token" form:"token"`
}

type CustomerLogin struct {
	Name    string `json:"name"`
	Email   string `gorm:"UNIQUE" json:"email"`
	Address string `json:"address" form:"address"`
	Gender  string `json:"gender" form:"gender" gorm:"type:enum('M', 'F')"`
	Token   string `json:"token" form:"token"`
}

type CustomerResponse struct {
	Status  bool       `json:"status"`
	Message string     `json:"message"`
	Data    []Customer `json:"data"`
}

type CustomerResponseSingle struct {
	Status  bool     `json:"status"`
	Message string   `json:"message"`
	Data    Customer `json:"data"`
}

type CustomerResponseLogin struct {
	Status  bool          `json:"status"`
	Message string        `json:"message"`
	Data    CustomerLogin `json:"data"`
}
