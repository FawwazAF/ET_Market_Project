package controller

import (
	"etmarket/project/lib/database"
	"net/http"

	"github.com/labstack/echo"
)

func GetAllMarket(c echo.Context) error {
	markets, err := database.GetManyMarkets()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  markets,
	})
}
