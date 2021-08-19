package controller

import (
	"etmarket/project/lib/database"
	"etmarket/project/models"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo"
)

func RegisterDriver(c echo.Context) error {
	//get user's input
	driver := models.Driver{}
	c.Bind(&driver)

	//check is email exists?
	is_email_exists, _ := database.CheckEmailOnDriver(driver.Email)
	if is_email_exists != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Email has already exist",
		})
	}

	//encrypt pass user
	convert_pwd := []byte(driver.Password) //convert pass from string to byte
	hashed_pwd := EncryptPwd(convert_pwd)
	driver.Password = hashed_pwd //set new pass

	//create new user
	data_driver, err := database.CreateDriver(driver)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new driver",
		"user":    data_driver,
	})
}

func LoginDriver(c echo.Context) error {
	//get user's input
	driver := models.Driver{}
	c.Bind(&driver)

	//compare password on form with db
	get_pwd := database.GetPwd(driver.Email) //get password
	err := bcrypt.CompareHashAndPassword([]byte(get_pwd), []byte(driver.Password))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "User Unauthorized. Email or Password not equal",
		})
	}

	//login
	data_driver, err := database.LoginDriver(driver.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "succes login as a driver",
		"users":  data_driver,
	})
}

func LogoutDriver(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("driver_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	logout, err := database.GetDriver(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "cannot get data",
		})
	}
	logout.Token = ""
	c.Bind(&logout)
	driver_updated, err := database.UpdateDriver(logout)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot logout",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Logout success!",
		"data":    driver_updated,
	})
}
