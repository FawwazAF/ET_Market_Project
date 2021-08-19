package controller

import (
	"etmarket/project/lib/database"
	"etmarket/project/models"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo"
)

func RegisterSeller(c echo.Context) error {
	//get user's input
	seller := models.Seller{}
	c.Bind(&seller)

	//check is email exists?
	isEmailExists := database.CheckEmailOnSeller(seller.Email)
	if isEmailExists {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Email has already exist",
		})
	}

	//encrypt pass user
	convert_pwd := []byte(seller.Password) //convert pass from string to byte
	hashed_pwd := EncryptPwd(convert_pwd)
	seller.Password = hashed_pwd //set new pass

	//create new user
	data_seller, err := database.CreateSeller(seller)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new seller",
		"user":    data_seller,
	})
}

func LoginSeller(c echo.Context) error {
	//get user's input
	seller := models.Seller{}
	c.Bind(&seller)

	//compare password on form with db
	get_pwd := database.GetPwd(seller.Email) //get password
	err := bcrypt.CompareHashAndPassword([]byte(get_pwd), []byte(seller.Password))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "User Unauthorized. Email or Password not equal",
		})
	}

	//login
	data_seller, err := database.LoginSeller(seller.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "succes login as a seller",
		"users":  data_seller,
	})
}

func GetDetailSeller(c echo.Context) error {
	seller_id, err := strconv.Atoi(c.Param("seller_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid seller id",
		})
	}
	data_seller, err := database.GetSellerById(seller_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Cant find seller",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"customer": data_seller,
	})
}
