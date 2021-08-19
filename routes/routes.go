package routes

import (
	"etmarket/project/constants"
	"etmarket/project/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(e *echo.Echo) {
	//-------------------------Product----------------------------//
	e.GET("/shop/:shop_id/product", controller.GetAllProductInShop)
	e.GET("/shop/:shop_id/product/name/:product_name", controller.GetSpecificProductInShop)
	e.GET("/shop/:shop_id/product/id/:product_id", controller.GetDetailSpecificProduct)

	//--------------------------Customer--------------------------//
	e.POST("/customer/register", controller.RegisterCustomer)
	e.POST("/customer/login", controller.LoginCustomer)

	//--------------------------Driver--------------------------//
	e.POST("/driver/register", controller.RegisterDriver)
	e.POST("/driver/login", controller.LoginDriver)

	//--------------------------Seller--------------------------//
	e.POST("/seller/register", controller.RegisterSeller)
	e.POST("/seller/login", controller.LoginSeller)
	e.GET("/seller/:seller_id/products", controller.GetSellerProducts)
	e.POST("/seller/:seller_id/products", controller.AddProductToSeller)
	e.PUT("/seller/:seller_id/products/:product_id", controller.EditSellerProduct)

	//GET list data of all markets available
	e.GET("/markets", controller.GetAllMarket)
	e.GET("/markets/:market_name", controller.GetSpecificMarket)

	//--------------------------Authorized Only--------------------------//
	r := e.Group("")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

}
