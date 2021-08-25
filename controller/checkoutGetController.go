package controller

import (
	"etmarket/project/lib/database"
	"etmarket/project/middlewares"
	"net/http"

	"github.com/labstack/echo"
)

// Ihsan
func GetCheckoutStatus(c echo.Context) error {
	logged_in_user_id := middlewares.ExtractToken(c)
	status := c.QueryParam("status")
	inprogress_checkout, err := database.GetHistory(status, logged_in_user_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "request not found",
		})
	}
	if len(inprogress_checkout) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "products not found",
		})
	}
	return c.JSON(http.StatusOK, inprogress_checkout)
}
func GetCheckoutStatusTesting() echo.HandlerFunc {
	return GetCheckoutStatus
}
