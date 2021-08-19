package routes

import (
	"etmarket/project/controller"

	"github.com/labstack/echo"
)

func New(e *echo.Echo) {
	e.GET("/markets/:market_id/shop", controller.GetAllCategoriesMarketIdController)
	e.GET("/markets/:market_id/shop/:category_name", controller.GetCategoryNameMarketIdController)
}
