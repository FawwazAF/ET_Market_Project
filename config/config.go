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
	DB.AutoMigrate(&models.Cart{})
	DB.AutoMigrate(&models.Category{})
	DB.AutoMigrate(&models.Checkout{})
	DB.AutoMigrate(&models.Customer{})
	DB.AutoMigrate(&models.Delivery{})
	DB.AutoMigrate(&models.Driver{})
	DB.AutoMigrate(&models.Market{})
	DB.AutoMigrate(&models.Seller{})
	DB.AutoMigrate(&models.Product{})
	DB.AutoMigrate(&models.Order{})
	DB.AutoMigrate(&models.Payment{})
}

func ConfigTest() (*gorm.DB, error) {
	var err error
<<<<<<< HEAD
	connectionStringTest := "root:welcome12345@tcp(localhost:3306)/etmarket_test?charset=utf8&parseTime=True&loc=Local"
=======
	connectionStringTest := "root:root@tcp(localhost:3306)/etmarket_test?charset=utf8&parseTime=True&loc=Local"
>>>>>>> b25d12e06fc46424c2cf04c70b6e29a6cfec4415
	DB, err = gorm.Open(mysql.Open(connectionStringTest), &gorm.Config{})
	if err != nil {
		return DB, err
	}
	return DB, err
}
