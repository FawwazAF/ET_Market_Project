package controller

import (
	"etmarket/project/lib/database"
	"etmarket/project/middlewares"
	"net/http"

	"github.com/labstack/echo"
)

func GetOrderList(c echo.Context) error {
	logged_driver_id := middlewares.ExtractToken(c)
	if logged_driver_id == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "please login first",
		})
	}
	markets, err := database.GetOrdertoTake()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  markets,
	})
}
