package controller

import (
	"etmarket/project/lib/database"
	"etmarket/project/middlewares"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

/*
Author: Patmiza
Getting all progress orders for a spesific logged in driver id
*/
func GetAllProgressOrdersController(c echo.Context) error {
	checkout_id, err := strconv.Atoi(c.QueryParam("checkout_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid input",
		})
	}
	logged_in_driver_id := middlewares.ExtractToken(c)
	orders, err := database.GetAllProgressOrders(logged_in_driver_id, checkout_id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if len(orders) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "orders not found",
		})
	}
	return c.JSON(http.StatusOK, orders)
}
func GetOrderDriverTesting() echo.HandlerFunc {
	return GetAllProgressOrdersController
}
