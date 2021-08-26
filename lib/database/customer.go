package database

import (
	"etmarket/project/config"
	"etmarket/project/middlewares"
	"etmarket/project/models"
)

/*
Author: Riska
This function for register customer
*/
func CreateCustomer(customer models.Customer) (models.Customer, error) {
	if err := config.DB.Save(&customer).Error; err != nil {
		return customer, err
	}

	return customer, nil
}

/*
Author: Riska
This function for login customer with matching data from database
*/
func LoginCustomer(email string) (models.Customer, error) {
	var customer models.Customer
	var err error
	if err = config.DB.Where("email = ?", email).First(&customer).Error; err != nil {
		return customer, err
	}
	customer.Token, err = middlewares.CreateToken(int(customer.ID))
	if err != nil {
		return customer, err
	}
	if err := config.DB.Save(customer).Error; err != nil {
		return customer, err
	}

	return customer, nil
}

/*
Author: Riska
This function for search password user by email
*/
func GetPwdCustomer(email string) string {
	var customer models.Customer
	config.DB.Where("email = ?", email).First(&customer)
	return customer.Password
}

/*
Author: Riska
This function for check is email customer exists
*/
func CheckEmailOnCustomer(email string) (interface{}, error) {
	var customer models.Customer

	if err := config.DB.Model(&customer).Select("email").Where("email=?", email).First(&customer.Email).Error; err != nil {
		return nil, err
	}

	return customer, nil
}

/*
Author: Riska
This function for get 1 specified customer with interface output
*/
func GetCustomerById(customer_id int) (models.Customer, error) {
	var customer models.Customer
	if err := config.DB.Where("id=?", customer_id).First(&customer).Error; err != nil {
		return customer, err
	}

	return customer, nil
}

/*
Author: Riska
This function for get 1 specified customer with Customer struct output
*/
func GetCustomer(id int) (models.Customer, error) {
	var customer models.Customer
	if err := config.DB.Find(&customer, "id=?", id).Error; err != nil {
		return customer, err
	}
	return customer, nil
}

/*
Author: Riska
This function for get email customer
*/
func GetEmailCustomerById(customer_id int) (string, error) {
	var customer models.Customer

	if err := config.DB.Model(&customer).Select("email").Where("id=?", customer_id).First(&customer.Email).Error; err != nil {
		return "nil", err
	}

	return customer.Email, nil
}

/*
Author: Riska
This function for update customer info from database
*/
func UpdateCustomer(customer models.Customer) (models.Customer, error) {
	if err := config.DB.Save(&customer).Error; err != nil {
		return customer, err
	}

	return customer, nil
}
