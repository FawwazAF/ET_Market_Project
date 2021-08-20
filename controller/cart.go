package controller

import (
	"etmarket/project/lib/database"
	"etmarket/project/middlewares"
	"etmarket/project/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

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

func GetAllCartsController(c echo.Context) error {
	logged_in_user_id := middlewares.ExtractToken(c)
	carts, err := database.GetAllCarts(logged_in_user_id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"Carts":  carts,
	})
}

func DeleteProductInCartsController(c echo.Context) error {
	product_id, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid product id",
		})
	}
	logged_in_user_id := middlewares.ExtractToken(c)
	carts, err := database.DeleteProductFromCart(logged_in_user_id, product_id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"Carts":  carts,
	})
}
