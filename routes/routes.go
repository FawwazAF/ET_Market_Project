package routes

import (
	"etmarket/project/controller"

	"github.com/labstack/echo"
)

func New(e *echo.Echo) {

	//GET list data of all markets available
	e.GET("/markets", controller.GetAllMarket)
	e.GET("/markets/:market_name", controller.GetSpecificMarket)

}
