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
<<<<<<< HEAD
	}

=======
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

func GetCustomerById(customer_id int) (interface{}, error) {
	var customer models.Customer
	if err := config.DB.Where("id=?", customer_id).First(&customer).Error; err != nil {
		return nil, err
	}
	return customer, nil
}

//get 1 specified customer with Customer struct output
func GetCustomer(id int) (models.Customer, error) {
	var customer models.Customer
	if err := config.DB.Find(&customer, "id=?", id).Error; err != nil {
		return customer, err
	}
	return customer, nil
}

//get email customer
func GetEmailCustomerById(customer_id int) (string, error) {
	var customer models.Customer

	if err := config.DB.Model(&customer).Select("email").Where("id=?", customer_id).First(&customer.Email).Error; err != nil {
		return "nil", err
	}

	return customer.Email, nil
}

//update customer info from database
func UpdateCustomer(customer models.Customer) (interface{}, error) {
	if err := config.DB.Save(&customer).Error; err != nil {
		return nil, err
	}
>>>>>>> 48353d70c37039491df3ae60b3ea01b095da5dc6
	return customer, nil
}
