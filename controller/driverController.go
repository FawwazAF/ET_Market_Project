package controller

import (
	"etmarket/project/lib/database"
	"etmarket/project/middlewares"
	"etmarket/project/models"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo"
)

/*
Author: Riska
This function for register driver
*/
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

	type Output struct {
		ID    uint   `json:"id"`
		Email string `json:"email"`
		Name  string `json:"name"`
	}

	//set output data
	output := Output{
		ID:    data_driver.ID,
		Email: data_driver.Email,
		Name:  data_driver.Name,
	}

	return c.JSON(http.StatusOK, output)
}

/*
Author: Riska
This function for login driver
*/
func LoginDriver(c echo.Context) error {
	//get user's input
	driver := models.Driver{}
	c.Bind(&driver)

	//compare password on form with db
	get_pwd := database.GetPwdDriver(driver.Email) //get password
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

	type Output struct {
		ID    uint   `json:"id"`
		Email string `json:"email"`
		Token string `json:"token"`
	}

	//set output data
	output := Output{
		ID:    data_driver.ID,
		Email: data_driver.Email,
		Token: data_driver.Token,
	}

	return c.JSON(http.StatusOK, output)
}

/*
Author: Riska
This function for get profile driver
*/
func GetDetailDriver(c echo.Context) error {
	driver_id, err := strconv.Atoi(c.Param("driver_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid driver id",
		})
	}

	//check otorisasi
	logged_in_user_id := middlewares.ExtractToken(c)
	if logged_in_user_id != driver_id {
		return echo.NewHTTPError(http.StatusUnauthorized, "This user unauthorized to get detail")
	}

	data_driver, err := database.GetDriverById(driver_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Cant find driver",
		})
	}

	type Output struct {
		ID      uint   `json:"id"`
		Email   string `json:"email"`
		Name    string `json:"name"`
		Address string `json:"address"`
		Gender  string `json:"gender"`
	}

	//set output data
	output := Output{
		ID:      data_driver.ID,
		Email:   data_driver.Email,
		Name:    data_driver.Name,
		Address: data_driver.Address,
		Gender:  data_driver.Gender,
	}

	return c.JSON(http.StatusOK, output)
}

func GetDetailDriverTesting() echo.HandlerFunc {
	return GetDetailDriver
}

/*
Author: Riska
This function for edit profile driver
*/
func UpdateDriver(c echo.Context) error {
	driver_id, err := strconv.Atoi(c.Param("driver_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	//check otorisasi
	logged_in_user_id := middlewares.ExtractToken(c)
	if logged_in_user_id != driver_id {
		return echo.NewHTTPError(http.StatusUnauthorized, "This user unauthorized to update data")
	}

	email_driver, err := database.GetEmailDriverById(driver_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "cannot get data",
		})
	}

	driver, err := database.GetDriver(driver_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "cannot get data",
		})
	}
	c.Bind(&driver)

	if driver.Email != email_driver {
		//check is email exists?
		is_email_exists, _ := database.CheckEmailOnDriver(driver.Email)
		if is_email_exists != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Email has already exist",
			})
		}
	}

	//encrypt pass user
	convert_pwd := []byte(driver.Password) //convert pass from string to byte
	hashed_pwd := EncryptPwd(convert_pwd)
	driver.Password = hashed_pwd //set new pass

	updated_driver, err := database.UpdateDriver(driver)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "cannot update data",
		})
	}

	type Output struct {
		ID      uint   `json:"id"`
		Email   string `json:"email"`
		Name    string `json:"name"`
		Address string `json:"address"`
		Gender  string `json:"gender"`
	}

	//set output data
	output := Output{
		ID:      updated_driver.ID,
		Email:   updated_driver.Email,
		Name:    updated_driver.Name,
		Address: updated_driver.Address,
		Gender:  updated_driver.Gender,
	}

	return c.JSON(http.StatusOK, output)
}

func UpdateDriverTesting() echo.HandlerFunc {
	return UpdateDriver
}

/*
Author: Riska
This function for logout driver
*/
func LogoutDriver(c echo.Context) error {
	driver_id, err := strconv.Atoi(c.Param("driver_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	//check otorisasi
	logged_in_user_id := middlewares.ExtractToken(c)
	if logged_in_user_id != driver_id {
		return echo.NewHTTPError(http.StatusUnauthorized, "This user unauthorized to logout")
	}

	logout, err := database.GetDriver(driver_id)
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

	type Output struct {
		ID      uint   `json:"id"`
		Email   string `json:"email"`
		Name    string `json:"name"`
		Address string `json:"address"`
		Gender  string `json:"gender"`
	}

	//set output data
	output := Output{
		ID:      driver_updated.ID,
		Email:   driver_updated.Email,
		Name:    driver_updated.Name,
		Address: driver_updated.Address,
		Gender:  driver_updated.Gender,
	}

	return c.JSON(http.StatusOK, output)
}

func LogoutDriverTesting() echo.HandlerFunc {
	return LogoutDriver
}
