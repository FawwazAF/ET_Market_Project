package controller

import (
	"etmarket/project/lib/database"
	"etmarket/project/middlewares"
	"net/http"

	"github.com/labstack/echo"
)

/*
Author: Patmiza
Getting all completed deliveries for a spesific logged in driver id
*/
func GetAllCompletedDeliveriesController(c echo.Context) error {
	logged_in_driver_id := middlewares.ExtractToken(c)
	deliveries, err := database.GetAllCompletedDeliveries(logged_in_driver_id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if len(deliveries) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "delivery not found",
		})
	}
	return c.JSON(http.StatusOK, deliveries)
}
func GetAllCompletedDeliveriesTesting() echo.HandlerFunc {
	return GetAllCompletedDeliveriesController
}
