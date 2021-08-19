package database

import (
	"etmarket/project/config"
	"etmarket/project/middlewares"
	"etmarket/project/models"
)

func CreateCustomer(customer models.Customer) (interface{}, error) {
	if err := config.DB.Save(&customer).Error; err != nil {
		return nil, err
	}
	return customer, nil
}

//login customer with matching data from database
func LoginCustomer(email string) (interface{}, error) {
	var customer models.Customer
	var err error
	if err = config.DB.Where("email = ?", email).First(&customer).Error; err != nil {
		return nil, err
	}
	customer.Token, err = middlewares.CreateToken(int(customer.ID))
	if err != nil {
		return nil, err
	}
	if err := config.DB.Save(customer).Error; err != nil {
		return nil, err
	}
	return customer, err
}

//search password user by email
func GetPwd(email string) string {
	var customer models.Customer
	config.DB.Where("email = ?", email).First(&customer)
	return customer.Password
}

func CheckEmailOnCustomer(email string) (interface{}, error) {
	var customer models.Customer

	if err := config.DB.Model(&customer).Select("email").Where("email=?", email).First(&customer.Email).Error; err != nil {
		return nil, err
	}

	return customer, nil
}

func GetManyPayment() (interface{}, error) {
	var payments []models.Payment
	if err := config.DB.Find(&payments).Error; err != nil {
		return nil, err
	}
	return payments, nil
}

func GetManyDrivers() (interface{}, error) {
	var drivers []models.Driver
	if err := config.DB.Find(&drivers).Error; err != nil {
		return drivers, err
	}

	return drivers, nil
}
