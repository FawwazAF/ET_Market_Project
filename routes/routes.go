package routes

import (
	"etmarket/project/constants"
	"etmarket/project/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(e *echo.Echo) {
	//-------------------------Market----------------------------//
	e.GET("/markets", controller.GetAllMarket)
	e.GET("/markets/:market_name", controller.GetSpecificMarket)

	//-------------------------Category----------------------------//
	e.GET("/categories", controller.GetAllCategories) //for register seller

	//-------------------------Shop----------------------------//
	e.GET("/markets/:market_id/shop", controller.GetAllCategoriesMarketIdController)
	e.GET("/markets/:market_id/shop/:category_name", controller.GetCategoryNameMarketIdController)

	//-------------------------Product----------------------------//
	e.GET("/shop/:shop_id/product", controller.GetAllProductInShop)
	e.GET("/shop/:shop_id/product/name/:product_name", controller.GetSpecificProductInShop)
	e.GET("/shop/:shop_id/product/id/:product_id", controller.GetDetailSpecificProduct)

	//--------------------------Customer--------------------------//
	e.POST("/customer/register", controller.RegisterCustomer)
	e.POST("/customer/login", controller.LoginCustomer)
	e.GET("/customer/:customer_id", controller.GetDetailCustomer)
	e.PUT("/customer/:customer_id", controller.UpdateCustomer)

	//--------------------------Driver--------------------------//
	e.POST("/driver/register", controller.RegisterDriver)
	e.POST("/driver/login", controller.LoginDriver)
	e.GET("/driver/:driver_id", controller.GetDetailDriver)
	e.PUT("/driver/:driver_id", controller.UpdateDriver)

	//--------------------------Seller--------------------------//
	e.POST("/seller/register", controller.RegisterSeller)
	e.POST("/seller/login", controller.LoginSeller)
	e.GET("/seller/:seller_id", controller.GetDetailSeller)
	e.PUT("/seller/:seller_id", controller.UpdateSeller)

	//--------------------------Authorized Only--------------------------//
	r := e.Group("")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	//--------------------------Customer--------------------------//
	// r.PUT("/customer/logout/:customer_id", controller.LogoutCustomer)

}
