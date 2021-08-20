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

// func Authorized(c echo.Context) (bool, models.Seller) {
// 	seller_id := middlewares.ExtractToken(c)
// 	token := database.GetTokenSeller(seller_id)
// 	seller, _ := database.GetSeller(seller_id)

// 	if seller.Token != token {
// 		return false, seller
// 	}
// 	return true, seller
// }

/*
Author: Riska
This function for register seller
*/
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

/*
Author: Riska
This function for login seller
*/
func LoginSeller(c echo.Context) error {
	//get user's input
	seller := models.Seller{}
	c.Bind(&seller)

	//compare password on form with db
	get_pwd := database.GetPwdSeller(seller.Email) //get password
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

/*
Author: Riska
This function for get profile seller
*/
func GetDetailSeller(c echo.Context) error {
	//convert id
	seller_id, err := strconv.Atoi(c.Param("seller_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid seller id",
		})
	}

	//check otorisasi
	logged_in_user_id := middlewares.ExtractToken(c)
	if logged_in_user_id != seller_id {
		return echo.NewHTTPError(http.StatusUnauthorized, "This user unauthorized to get detail")
	}

	//get data seller
	data_seller, err := database.GetSellerById(seller_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Cannot find seller",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"seller": data_seller,
	})
}

/*
Author: Riska
This function for edit profile seller
*/
func UpdateSeller(c echo.Context) error {
	//convert id
	seller_id, err := strconv.Atoi(c.Param("seller_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	//check otorisasi
	logged_in_user_id := middlewares.ExtractToken(c)
	if logged_in_user_id != seller_id {
		return echo.NewHTTPError(http.StatusUnauthorized, "This user unauthorized to update data")
	}

	//get email seller
	email_seller, err := database.GetEmailSellerById(seller_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "cannot get data",
		})
	}

	//get seller
	seller, err := database.GetSeller(seller_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "cannot get data",
		})
	}
	c.Bind(&seller)

	//check email
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

	//update data seller
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

// Ihsan
func GetSellerProducts(c echo.Context) error {
	// convert parameter to variable
	seller_id, err := strconv.Atoi(c.Param("seller_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid seller id",
		})
	}
	// Auth
	logged_in_user_id := middlewares.ExtractToken(c)
	if logged_in_user_id != seller_id {
		return echo.NewHTTPError(http.StatusUnauthorized, "This user unauthorized to get detail")
	}
	// Get data
	all_products_selected_seller, err := database.GetAllSellerProduct(seller_id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "no products in the shop",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "get all products from this shop success",
		"data":    all_products_selected_seller,
	})
}

// Ihsan
func AddProductToSeller(c echo.Context) error {
	// Convert parameter to variable
	seller_id, err := strconv.Atoi(c.Param("seller_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	// Auth
	logged_in_user_id := middlewares.ExtractToken(c)
	if logged_in_user_id != seller_id {
		return echo.NewHTTPError(http.StatusUnauthorized, "This user unauthorized to get detail")
	}
	// http request body
	product := models.Product{}
	product.SellerID = uint(seller_id)
	c.Bind(&product)
	// Add data from request
	product_added, err := database.AddProductToSeller(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot add product",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    product_added,
	})
}

// ihsan
func EditSellerProduct(c echo.Context) error {
	seller_id, err := strconv.Atoi(c.Param("seller_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid seller id",
		})
	}

	logged_in_user_id := middlewares.ExtractToken(c)
	if logged_in_user_id != seller_id {
		return echo.NewHTTPError(http.StatusUnauthorized, "This user unauthorized to get detail")
	}

	product_id, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid product id",
		})
	}

	product := models.Product{}
	c.Bind(&product)

	product_edited, err := database.EditSellerProduct(seller_id, product_id, product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot edit product",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    product_edited,
	})
}

/*
Author: Riska
This function for get all order by seller id
*/
func GetAllOrders(c echo.Context) error {
	logged_in_user_id := middlewares.ExtractToken(c)
	list_product, err := database.GetAllProductsOrderBySellerId(logged_in_user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot get list product",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": list_product,
	})
}

/*
Author: Riska
This function for edit status item order
*/
func EditStatusItemOrder(c echo.Context) error {
	logged_in_user_id := middlewares.ExtractToken(c)

	order_id, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid order id",
		})
	}

	seller_id, err := database.GetSellerIdByOderId(order_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot get seller id",
		})
	}

	//check otorisasi
	if logged_in_user_id != seller_id {
		return echo.NewHTTPError(http.StatusUnauthorized, "This user unauthorized edit status order")
	}

	order, err := database.GetItemOrderByOrderId(order_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot get order",
		})
	}

	order.Status = "completed"
	c.Bind(&order)

	update_status, err := database.EditItemStatus(order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot update status product",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": update_status,
	})
}

/*
Author: Riska
This function for logout seller
*/
func LogoutSeller(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("seller_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	logout, err := database.GetSeller(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "cannot get data",
		})
	}
	logout.Token = ""
	c.Bind(&logout)
	seller_updated, err := database.UpdateSeller(logout)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot logout",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Logout success!",
		"data":    seller_updated,
	})
}
