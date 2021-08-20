package controller

import (
	"etmarket/project/lib/database"
	"etmarket/project/middlewares"
	"etmarket/project/models"
	"net/http"

	"github.com/labstack/echo"
)

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

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success",
		"checkout": new_checkout,
	})
}
