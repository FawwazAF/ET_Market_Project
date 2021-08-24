package controller

import (
	"etmarket/project/lib/database"
	"net/http"

	"github.com/labstack/echo"
)

//Fawwaz
//Get all data market available
func GetAllMarket(c echo.Context) error {
	markets, err := database.GetManyMarkets()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, markets)
}

//Fawwaz
//Get specific market by name
func GetSpecificMarket(c echo.Context) error {
	key := c.Param("market_name")
	markets, err := database.SearchMarket(key)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, markets)
}
