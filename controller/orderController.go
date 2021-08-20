package controller

import (
	"etmarket/project/lib/database"
	"etmarket/project/middlewares"
	"net/http"

	"github.com/labstack/echo"
)

/*
Author: Patmiza
Getting all progress orders for a spesific logged in driver id
*/
func GetAllProgressOrdersController(c echo.Context) error {
	logged_in_driver_id := middlewares.ExtractToken(c)
	orders, err := database.GetAllProgressOrders(logged_in_driver_id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"Orders": orders,
	})
}
