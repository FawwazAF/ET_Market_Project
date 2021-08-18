package routes

import (
	"etmarket/project/controller"

	"github.com/labstack/echo"
)

func New(e *echo.Echo) {
	e.GET("/markets/shop/:shop_id", controller.GetAllProductInShop)
	e.GET("/markets/shop/:shop_id/:product_name", controller.GetSpecificProductInShop)
}
