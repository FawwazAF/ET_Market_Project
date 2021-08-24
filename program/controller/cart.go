package controller

import (
	"etmarket/project/lib/database"
	"etmarket/project/middlewares"
	"etmarket/project/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

/*
Author: Patmiza
*/
func InsertProductIntoCartController(c echo.Context) error {
	logged_in_user_id := middlewares.ExtractToken(c)

	// checking shop id
	seller_id, err := strconv.Atoi(c.Param("seller_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid customer id",
		})
	}

	// checking product id
	product_id, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid product id",
		})
	}

	var carts models.Cart
	c.Bind(&carts)

	// insert product into cart
	new_cart, err := database.InsertProductIntoCart(logged_in_user_id, seller_id, product_id, carts)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cart is not found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success create new cart",
		"Cart":   new_cart,
	})
}

/*
Author: Patmiza
*/
func GetAllCartsController(c echo.Context) error {
	logged_in_user_id := middlewares.ExtractToken(c)
	carts, err := database.GetAllCarts(logged_in_user_id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if len(carts) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "cart is empty")
	}
	return c.JSON(http.StatusOK, carts)
}

/*
Author: Patmiza
*/
func DeleteProductInCartsController(c echo.Context) error {
	//checking product id
	product_id, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid product id",
		})
	}

	//extracting token of customer id
	logged_in_customer_id := middlewares.ExtractToken(c)
	carts, err := database.DeleteProductFromCart(logged_in_customer_id, product_id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"Carts":  carts,
	})
}
