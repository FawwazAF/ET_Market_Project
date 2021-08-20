package routes

import (
	"etmarket/project/constants"
	"etmarket/project/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(e *echo.Echo) {

	//--------------------------Customer--------------------------//
	e.POST("/customer/register", controller.RegisterCustomer)
	e.POST("/customer/login", controller.LoginCustomer)

	//--------------------------Driver--------------------------//
	e.POST("/driver/register", controller.RegisterDriver)
	e.POST("/driver/login", controller.LoginDriver)

	//--------------------------Seller--------------------------//
	e.POST("/seller/register", controller.RegisterSeller)
	e.POST("/seller/login", controller.LoginSeller)

	//--------------------------Authorized Only--------------------------//
	r := e.Group("")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	//-------------------------Market----------------------------//
	r.GET("/markets", controller.GetAllMarket)
	r.GET("/markets/:market_name", controller.GetSpecificMarket)

	//-------------------------Category----------------------------//
	r.GET("/categories", controller.GetAllCategories) //for register seller

	//-------------------------Shop----------------------------//
	r.GET("/markets/:market_id/sellers", controller.GetAllCategoriesMarketIdController)
	r.GET("/markets/:market_id/sellers/:category_name", controller.GetSellerController)

	//-------------------------Product----------------------------//
	r.GET("/shop/:shop_id/product", controller.GetAllProductInShop)
	r.GET("/shop/:shop_id/product/name/:product_name", controller.GetSpecificProductInShop)
	r.GET("/shop/:shop_id/product/id/:product_id", controller.GetDetailSpecificProduct)

	//-------------------------Cart----------------------------//
	r.POST("/shop/:shop_id/product/id/:product_id", controller.InsertProductIntoCartController)
	r.GET("/cart", controller.GetProductInCartContorller)
	r.DELETE("/cart/produtc/:product_id", controller.DeleteProductInCartsController)
}
