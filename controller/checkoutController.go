package controller

import (
	"etmarket/project/lib/database"
	"etmarket/project/middlewares"
	"etmarket/project/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

//Fawwaz
//Get list all payment method before checkout
func GetAllPaymentMethod(c echo.Context) error {
	payments, err := database.GetManyPayment()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, payments)
}

//Fawwaz
//Make a checkout
func CheckoutTransaction(c echo.Context) error {

	var checkout models.Checkout
	c.Bind(&checkout)
	logged_in_user_id := middlewares.ExtractToken(c)

	//Find Carts
	carts, err := database.FindCarts(logged_in_user_id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if len(carts) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "cart is empty",
		})
	}

	//Create new checkout
	new_checkout, err := database.Checkout(logged_in_user_id, checkout, carts)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	//Migrate data from cart to order
	if err := database.CartMigrate(logged_in_user_id, int(new_checkout.ID)); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Hard Delete cart
	if err := database.DeleteCart(logged_in_user_id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, new_checkout)
}

// Ihsan
func GetCheckoutStatusInProgress(c echo.Context) error {
	checkout_id_auth := middlewares.ExtractToken(c)
	status := c.QueryParam("status")
	in_progress_checkout, err := database.GetHistoryInProgress(status, checkout_id_auth)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "request not found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "request success",
		"data":    in_progress_checkout,
	})
}

// Ihsan
func GetCheckoutStatusComplete(c echo.Context) error {
	checkout_id_auth := middlewares.ExtractToken(c)
	status := c.QueryParam("status")
	in_complete_checkout, err := database.GetHistoryInProgress(status, checkout_id_auth)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "request not found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "request success",
		"data":    in_complete_checkout,
	})
}

// Ihsan
func GetSelectedOrder(c echo.Context) error {
	checkout_id, err := strconv.Atoi(c.Param("checkout_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid checkout id",
		})
	}
	data, err := database.GetSelectedOrder(checkout_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succesfully retrieve data",
		"data":    data,
	})
}
