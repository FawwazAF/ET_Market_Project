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
	e.GET("/markets/:market_id/seller", controller.GetAllCategoriesMarketIdController)
	e.GET("/markets/:market_id/seller/:category_name", controller.GetCategoryNameMarketIdController)

	//-------------------------Product----------------------------//
	e.GET("/seller/:seller_id/product", controller.GetAllProductInShop)
	e.GET("/seller/:seller_id/product/name/:product_name", controller.GetSpecificProductInShop)
	e.GET("/seller/:seller_id/product/id/:product_id", controller.GetDetailSpecificProduct)

	//--------------------------Customer--------------------------//
	e.POST("/customer/register", controller.RegisterCustomer)
	e.POST("/customer/login", controller.LoginCustomer)
	e.GET("/payments", controller.GetAllPaymentMethod)

	//--------------------------Driver--------------------------//
	e.POST("/driver/register", controller.RegisterDriver)
	e.POST("/driver/login", controller.LoginDriver)

	//--------------------------Seller--------------------------//
	e.POST("/seller/register", controller.RegisterSeller)
	e.POST("/seller/login", controller.LoginSeller)

	//--------------------------Authorized Only--------------------------//
	r := e.Group("")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	//--------------------------Checkout--------------------------//
	r.POST("/checkout", controller.CheckoutTransaction)
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

	//-------------------------Cart----------------------------//
	r.POST("/seller/:seller_id/product/id/:product_id", controller.InsertProductIntoCartController)
	r.GET("/cart", controller.GetAllCartsController)
	r.DELETE("/cart/produtc/:product_id", controller.DeleteProductInCartsController)

	//--------------------------Checkout--------------------------//
	r.POST("/checkout", controller.CheckoutTransaction)

	//--------------------------Customer--------------------------//
	r.GET("/customer/:customer_id", controller.GetDetailCustomer)
	r.PUT("/customer/:customer_id", controller.UpdateCustomer)
	r.PUT("/customer/logout/:customer_id", controller.LogoutCustomer)

	//--------------------------Seller--------------------------//
	r.GET("/seller/:seller_id", controller.GetDetailSeller)
	r.PUT("/seller/:seller_id", controller.UpdateSeller)
	r.GET("/seller/:seller_id/products", controller.GetSellerProducts)
	r.POST("/seller/:seller_id/products", controller.AddProductToSeller)
	r.PUT("/seller/:seller_id/products/:product_id", controller.EditSellerProduct)
	r.PUT("/seller/logout/:seller_id", controller.LogoutSeller)

	//--------------------------Driver--------------------------//
	r.GET("/driver/:driver_id", controller.GetDetailDriver)
	r.PUT("/driver/:driver_id", controller.UpdateDriver)
	r.PUT("/driver/logout/:driver_id", controller.LogoutDriver)

	r.GET("/driver/orderlist", controller.GetOrderList)

}
