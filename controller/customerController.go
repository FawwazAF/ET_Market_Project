package controller

import (
	"etmarket/project/lib/database"
	"etmarket/project/models"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo"
)

func EncryptPwd(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

func RegisterCustomer(c echo.Context) error {
	//get user's input
	customer := models.Customer{}
	c.Bind(&customer)

	//check is email exists?
	isEmailExists := database.CheckEmailOnCustomer(customer.Email)
	if isEmailExists {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Email has already exist",
		})
	}

	//encrypt pass user
	convert_pwd := []byte(customer.Password) //convert pass from string to byte
	hashed_pwd := EncryptPwd(convert_pwd)
	customer.Password = hashed_pwd //set new pass

	//create new user
	data_customer, err := database.CreateCustomer(customer)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new customer",
		"user":    data_customer,
	})
}

func LoginCustomer(c echo.Context) error {
	//get user's input
	customer := models.Customer{}
	c.Bind(&customer)

	//compare password on form with db
	get_pwd := database.GetPwd(customer.Email) //get password
	err := bcrypt.CompareHashAndPassword([]byte(get_pwd), []byte(customer.Password))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "User Unauthorized. Email or Password not equal",
		})
	}

	//login
	data_customer, err := database.LoginCustomer(customer.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "succes login as a customer",
		"users":  data_customer,
	})
}
