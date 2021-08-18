package config

import (
	"etmarket/project/models"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var HTTP_PORT int

func InitDb() {
	var err error
	// connectionString := os.Getenv("CONNECTION_STRING")
	// DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	// if err != nil {
	// 	panic(err)
	// }
	// InitMigrate()
	DB, err = gorm.Open(mysql.Open("root:Minus12345@tcp(localhost:3306)/etmarket_schema?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrate()
}

func InitPort() {
	var err error
	// HTTP_PORT, err = strconv.Atoi(os.Getenv("HTTP_PORT"))
	// if err != nil {
	// 	panic(err)
	// }
	HTTP_PORT, err = strconv.Atoi("8080")
	if err != nil {
		panic(err)
	}
}

func InitMigrate() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Market{})
}
