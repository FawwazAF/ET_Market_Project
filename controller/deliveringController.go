package controller

import (
	"etmarket/project/lib/database"
	"etmarket/project/middlewares"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

//Fawwaz
//List of Order to take
func GetOrderList(c echo.Context) error {
	logged_driver_id := middlewares.ExtractToken(c)
	if logged_driver_id == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "please login first",
		})
	}
	checkout, err := database.GetOrdertoTake()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if len(checkout) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "checkout not found",
		})
	}
	type Response struct {
		ID         uint   `json :"id"`
		CustomerID uint   `json:"customer_id"`
		TotalQty   int    `json:"total_qty"`
		TotalPrice int    `json:"total_price"`
		Status     string `json:"status" gorm:"type:enum('searching', 'delivery', 'completed')"`
	}
	var responses []Response
	for i := 0; i < len(checkout); i++ {
		new_array := Response{
			ID:         checkout[i].ID,
			CustomerID: checkout[i].CustomerID,
			TotalQty:   checkout[i].TotalQty,
			TotalPrice: checkout[i].TotalPrice,
			Status:     checkout[i].Status,
		}
		responses = append(responses, new_array)
	}
	return c.JSON(http.StatusOK, responses)
}
func GetOrderListTesting() echo.HandlerFunc {
	return GetOrderList
}

//Fawwaz
//Take Order
func TakeCheckout(c echo.Context) error {
	//Login
	logged_driver_id := middlewares.ExtractToken(c)
	if logged_driver_id == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "please login first",
		})
	}
	//Check Checkout ID
	checkout_id, err := strconv.Atoi(c.Param("checkout_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "input invalid",
		})
	}

	delivery, err := database.MakeDelivery(logged_driver_id, checkout_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "driver or checkout is not exist",
		})
	}

	return c.JSON(http.StatusOK, delivery)
}
func TakeCheckoutTesting() echo.HandlerFunc {
	return TakeCheckout
}

//Fawwaz
//Update status after finishing the delivery
func FinishedDelivery(c echo.Context) error {
	logged_driver_id := middlewares.ExtractToken(c)
	if logged_driver_id == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "please login first",
		})
	}
	//Check Checkout ID
	checkout_id, err := strconv.Atoi(c.Param("checkout_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "input invalid",
		})
	}
	delivery, err := database.EditDelivery(logged_driver_id, checkout_id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, delivery)
}
func FinishedDeliveryTesting() echo.HandlerFunc {
	return FinishedDelivery
}
