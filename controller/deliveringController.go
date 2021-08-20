package controller

import (
	"etmarket/project/lib/database"
	"etmarket/project/middlewares"
	"net/http"
	"strconv"

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

func TakeCheckout(c echo.Context) error {
	//Login
	logged_driver_id := middlewares.ExtractToken(c)
	if logged_driver_id == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "please login first",
		})
	}
	//Check Checkout ID
	checkout_id, err := strconv.Atoi(c.Param("checkout_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "input invalid",
		})
	}

	delivery, err := database.MakeDelivery(logged_driver_id, checkout_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "driver or checkout is not exist",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  delivery,
	})
}

func FinishedDelivery(c echo.Context) error {
	logged_driver_id := middlewares.ExtractToken(c)
	if logged_driver_id == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "please login first",
		})
	}
	//Check Checkout ID
	checkout_id, err := strconv.Atoi(c.Param("checkout_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "input invalid",
		})
	}
	delivery, err := database.EditDelivery(logged_driver_id, checkout_id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  delivery,
	})
}
