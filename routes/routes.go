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
	e.POST("/customer/register", controller.RegisterCustomer) //Riska
	e.POST("/customer/login", controller.LoginCustomer)       //Riska
	e.GET("/payments", controller.GetAllPaymentMethod)

	//--------------------------Driver--------------------------//
	e.POST("/driver/register", controller.RegisterDriver) //Riska
	e.POST("/driver/login", controller.LoginDriver)       //Riska

	//--------------------------Seller--------------------------//
	e.POST("/seller/register", controller.RegisterSeller) //Riska
	e.POST("/seller/login", controller.LoginSeller)       //Riska

	//--------------------------History----------------------//
	e.GET("/history/checkout_id/orders", controller.GetSelectedOrder) // Ihsan

	//--------------------------Authorized Only--------------------------//
	r := e.Group("")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	//--------------------------Checkout--------------------------//
	r.POST("/checkout", controller.CheckoutTransaction)

	//--------------------------History-------------------------//
	r.GET("/history?status=progress", controller.GetCheckoutStatusInProgress) // Ihsan
	r.GET("/history?status=completed", controller.GetCheckoutStatusComplete)  // ihsan

	//--------------------------Customer--------------------------//
	r.GET("/customer/:customer_id", controller.GetDetailCustomer) //Riska
	r.PUT("/customer/:customer_id", controller.UpdateCustomer)    //Riska

	//--------------------------Seller--------------------------//
	r.GET("/seller/:seller_id", controller.GetDetailSeller)                        //Riska
	r.PUT("/seller/:seller_id", controller.UpdateSeller)                           //Riska
	r.GET("/seller/:seller_id/products", controller.GetSellerProducts)             // Ihsan
	r.POST("/seller/:seller_id/products", controller.AddProductToSeller)           // Ihsan
	r.PUT("/seller/:seller_id/products/:product_id", controller.EditSellerProduct) // Ihsan

	//--------------------------Driver--------------------------//
	r.GET("/driver/:driver_id", controller.GetDetailDriver) //Riska
	r.PUT("/driver/:driver_id", controller.UpdateDriver)    //Riska

	//-------------------------Cart----------------------------//
	r.POST("/seller/:seller_id/product/id/:product_id", controller.InsertProductIntoCartController)
	r.GET("/cart", controller.GetAllCartsController)
	r.DELETE("/cart/produtc/:product_id", controller.DeleteProductInCartsController)

	//--------------------------Checkout--------------------------//
	r.POST("/checkout", controller.CheckoutTransaction)

	//--------------------------Customer--------------------------//
	r.GET("/customer/:customer_id", controller.GetDetailCustomer)     //Riska
	r.PUT("/customer/:customer_id", controller.UpdateCustomer)        //Riska
	r.PUT("/customer/logout/:customer_id", controller.LogoutCustomer) //Riska

	//--------------------------Seller--------------------------//
	r.GET("/seller/:seller_id", controller.GetDetailSeller) //Riska
	r.PUT("/seller/:seller_id", controller.UpdateSeller)    //Riska
	r.GET("/seller/:seller_id/products", controller.GetSellerProducts)
	r.POST("/seller/:seller_id/products", controller.AddProductToSeller)
	r.PUT("/seller/:seller_id/products/:product_id", controller.EditSellerProduct)
	r.GET("/seller/orderlist", controller.GetAllOrders)                  //Riska
	r.PUT("/seller/orderlist/:order_id", controller.EditStatusItemOrder) //Riska
	r.PUT("/seller/logout/:seller_id", controller.LogoutSeller)          //Riska

	//--------------------------Driver--------------------------//
	r.GET("/driver/orderlist", controller.GetOrderList)
	r.POST("/driver/orderlist/:checkout_id", controller.TakeCheckout)
	r.PUT("/driver/orderlist/:checkout_id", controller.FinishedDelivery)
	r.GET("/driver/:driver_id", controller.GetDetailDriver)     //Riska
	r.PUT("/driver/:driver_id", controller.UpdateDriver)        //Riska
	r.PUT("/driver/logout/:driver_id", controller.LogoutDriver) //Riska

}
