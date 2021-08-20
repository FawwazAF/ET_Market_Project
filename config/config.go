package config

import (
	"etmarket/project/models"
	"os"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var HTTP_PORT int

func InitDb() {
	var err error
	connectionString := os.Getenv("CONNECTION_STRING")
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrate()
}

func InitPort() {
	var err error
	HTTP_PORT, err = strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		panic(err)
	}
}

func InitMigrate() {
	DB.AutoMigrate(&models.Customer{})
	DB.AutoMigrate(&models.Driver{})
	DB.AutoMigrate(&models.Market{})
	DB.AutoMigrate(&models.Seller{})
	DB.AutoMigrate(&models.Product{})
	DB.AutoMigrate(&models.Market{})
	DB.AutoMigrate(&models.Category{})
	DB.AutoMigrate(&models.Cart{})
	DB.AutoMigrate(&models.Checkout{})
	DB.AutoMigrate(&models.Order{})
	DB.AutoMigrate(&models.Delivery{})
}
