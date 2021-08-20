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

	//--------------------------Driver--------------------------//
	e.POST("/driver/register", controller.RegisterDriver)
	e.POST("/driver/login", controller.LoginDriver)

	//--------------------------Seller--------------------------//
	e.POST("/seller/register", controller.RegisterSeller)
	e.POST("/seller/login", controller.LoginSeller)

	//--------------------------Authorized Only--------------------------//
	r := e.Group("")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	//--------------------------Customer--------------------------//
	r.GET("/customer/:customer_id", controller.GetDetailCustomer)
	r.PUT("/customer/:customer_id", controller.UpdateCustomer)

	//--------------------------Seller--------------------------//
	r.GET("/seller/:seller_id", controller.GetDetailSeller)
	r.PUT("/seller/:seller_id", controller.UpdateSeller)
	r.GET("/seller/:seller_id/products", controller.GetSellerProducts)
	r.POST("/seller/:seller_id/products", controller.AddProductToSeller)
	r.PUT("/seller/:seller_id/products/:product_id", controller.EditSellerProduct)

	//--------------------------Driver--------------------------//
	r.GET("/driver/:driver_id", controller.GetDetailDriver)
	r.PUT("/driver/:driver_id", controller.UpdateDriver)

	//--------------------------Customer--------------------------//
	// r.PUT("/customer/logout/:customer_id", controller.LogoutCustomer)

}
