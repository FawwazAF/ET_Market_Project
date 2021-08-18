package routes

import (
	"etmarket/project/controller"

	"github.com/labstack/echo"
)

func New(e *echo.Echo) {
	e.GET("/markets", controller.GetAllMarket)
	e.GET("/markets/:market_name", controller.GetSpecificMarket)
	e.GET("/markets/:id_market/categories", controller.GetAllCategoriesMarketIdController)
	e.GET("/markets/:id_market/categories/:name_category", controller.GetCategoryNameMarketIdController)
}
