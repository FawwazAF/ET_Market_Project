package controller

import (
	"etmarket/project/lib/database"
	"etmarket/project/middlewares"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// Ihsan
func GetCheckoutStatusInProgress(c echo.Context) error {
	checkout_id, err := strconv.Atoi(c.Param("checkout_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid checkout id",
		})
	}

	logged_in_user_id := middlewares.ExtractToken(c)
	if logged_in_user_id != checkout_id {
		return echo.NewHTTPError(http.StatusUnauthorized, "This user unauthorized to get detail")
	}

	status := c.QueryParam("status")
	in_progress_checkout, err := database.GetHistoryInProgress(status, checkout_id)
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
	checkout_id, err := strconv.Atoi(c.Param("checkout_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid checkout id",
		})
	}

	logged_in_user_id := middlewares.ExtractToken(c)
	if logged_in_user_id != checkout_id {
		return echo.NewHTTPError(http.StatusUnauthorized, "This user unauthorized to get detail")
	}

	status := c.QueryParam("status")
	in_complete_checkout, err := database.GetHistoryInProgress(status, checkout_id)
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
