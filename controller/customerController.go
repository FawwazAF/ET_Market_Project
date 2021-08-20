package controller

import (
	"etmarket/project/lib/database"
	"etmarket/project/middlewares"
	"etmarket/project/models"
	"log"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo"
)

/*
Author: Riska
This function for encrypt password
*/
func EncryptPwd(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

/*
Author: Riska
This function for register customer
*/
func RegisterCustomer(c echo.Context) error {
	//get user's input
	customer := models.Customer{}
	c.Bind(&customer)

	//check is email exists?
	is_email_exists, _ := database.CheckEmailOnCustomer(customer.Email)
	if is_email_exists != nil {
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

/*
Author: Riska
This function for login customer
*/
func LoginCustomer(c echo.Context) error {
	//get user's input
	customer := models.Customer{}
	c.Bind(&customer)

	//compare password on form with db
	get_pwd := database.GetPwdCustomer(customer.Email) //get password
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

/*
Author: Riska
This function for get profile customer
*/
func GetDetailCustomer(c echo.Context) error {
	customer_id, err := strconv.Atoi(c.Param("customer_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid customer id",
		})
	}

	//check otorisasi
	logged_in_user_id := middlewares.ExtractToken(c)
	if logged_in_user_id != customer_id {
		return echo.NewHTTPError(http.StatusUnauthorized, "This user unauthorized to get detail")
	}

	//get customer by id
	data_customer, err := database.GetCustomerById(customer_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Cant find customer",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"customer": data_customer,
	})
}

/*
Author: Riska
This function for edit profile customer
*/
func UpdateCustomer(c echo.Context) error {
	customer_id, err := strconv.Atoi(c.Param("customer_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	//check otorisasi
	logged_in_user_id := middlewares.ExtractToken(c)
	if logged_in_user_id != customer_id {
		return echo.NewHTTPError(http.StatusUnauthorized, "This user unauthorized to update data")
	}

	//get email customer
	email_customer, err := database.GetEmailCustomerById(customer_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "cannot get data",
		})
	}

	//get customer
	customer, err := database.GetCustomer(customer_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "cannot get data",
		})
	}
	c.Bind(&customer)

	//check email
	if customer.Email != email_customer {
		//check is email exists?
		is_email_exists, _ := database.CheckEmailOnCustomer(customer.Email)
		if is_email_exists != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Email has already exist",
			})
		}
	}

	//encrypt pass user
	convert_pwd := []byte(customer.Password) //convert pass from string to byte
	hashed_pwd := EncryptPwd(convert_pwd)
	customer.Password = hashed_pwd //set new pass

	//update data customer
	updated_customer, err := database.UpdateCustomer(customer)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "cannot update data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":       "success update customer",
		"data customer": updated_customer,
	})
}

/*
Author: Riska
This function for logout customer
*/
func LogoutCustomer(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("customer_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	logout, err := database.GetCustomer(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "cannot get data",
		})
	}
	logout.Token = ""
	c.Bind(&logout)
	customer_updated, err := database.UpdateCustomer(logout)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot logout",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Logout success!",
		"data":    customer_updated,
	})
}
