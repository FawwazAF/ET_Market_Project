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
	is_email_exists, _ := database.CheckEmailOnSeller(seller.Email)
	if is_email_exists != nil {
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
		"seller": data_seller,
	})
}

func UpdateSeller(c echo.Context) error {
	seller_id, err := strconv.Atoi(c.Param("seller_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	email_seller, err := database.GetEmailSellerById(seller_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "cannot get data",
		})
	}

	seller, err := database.GetSeller(seller_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "cannot get data",
		})
	}
	c.Bind(&seller)

	if seller.Email != email_seller {
		//check is email exists?
		is_email_exists, _ := database.CheckEmailOnSeller(seller.Email)
		if is_email_exists != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Email has already exist",
			})
		}
	}

	//encrypt pass user
	convert_pwd := []byte(seller.Password) //convert pass from string to byte
	hashed_pwd := EncryptPwd(convert_pwd)
	seller.Password = hashed_pwd //set new pass

	updated_seller, err := database.UpdateSeller(seller)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "cannot update data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":     "success update seller",
		"data seller": updated_seller,
	})
}
