package controller

import (
	"bytes"
	"etmarket/project/lib/database"
	"etmarket/project/middlewares"
	"etmarket/project/models"
	"fmt"
	"html/template"
	"net/http"
	"net/smtp"
	"strconv"

	"github.com/labstack/echo"
	"github.com/vonage/vonage-go-sdk"
)

//Fawwaz
//Get list all payment method before checkout
func GetAllPaymentMethod(c echo.Context) error {
	payments, err := database.GetManyPayment()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, payments)
}

//Fawwaz
//Make a checkout
func CheckoutTransaction(c echo.Context) error {

	var checkout models.Checkout
	c.Bind(&checkout)
	logged_in_user_id := middlewares.ExtractToken(c)

	//Find Carts
	carts, err := database.FindCarts(logged_in_user_id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if len(carts) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "cart is empty",
		})
	}

	//Create new checkout
	new_checkout, err := database.Checkout(logged_in_user_id, checkout, carts)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	type Response struct {
		ID            uint
		TotalQty      int    `json:"total_qty"`
		TotalPrice    int    `json:"total_price"`
		CustomerID    uint   `json:"customer_id"`
		PaymentID     uint   `json:"payment_id"`
		Status        string `json:"status" gorm:"type:enum('searching', 'delivery', 'completed')"`
		DeliveryPrice int    `json:"delivery_price"`
	}
	response := Response{
		ID:            new_checkout.ID,
		TotalQty:      new_checkout.TotalQty,
		TotalPrice:    new_checkout.TotalPrice,
		CustomerID:    new_checkout.CustomerID,
		PaymentID:     new_checkout.PaymentID,
		Status:        new_checkout.Status,
		DeliveryPrice: 5000,
	}
	//Migrate data from cart to order
	if err := database.CartMigrate(logged_in_user_id, int(new_checkout.ID)); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Hard Delete cart
	if err := database.DeleteCart(logged_in_user_id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	seller := database.GetSellerByCheckoutId(new_checkout.ID)

	SendEmail(new_checkout.ID, logged_in_user_id)
	SendNotifSms(new_checkout.ID, seller.ID, seller.Name)
	return c.JSON(http.StatusOK, response)
}

func CheckoutTransactionTesting() echo.HandlerFunc {
	return CheckoutTransaction
}

/*
Riska
This function for send email notification for customer that the purchase is success
*/
func SendEmail(checkout_id uint, customer_id int) {
	email_customer, _ := database.GetEmailCustomerById(customer_id)
	customer, _ := database.GetCustomer(customer_id)
	name_customer := customer.Name
	total_price := database.GetTotalPrice(checkout_id)

	list_order, _ := database.GetListOrderOnCheckoutForCustomer(checkout_id)
	// Sender data.
	from := "etmarket.group3@gmail.com"
	password := "etmarket2021"

	// Receiver email address.
	to := []string{
		email_customer,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	t, _ := template.ParseFiles("template.html")

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: [E-T Market] Thank you for your purchase \n%s\n\n", mimeHeaders)))

	type Result struct {
		Name      string
		Price     int
		ProductID int
		Qty       int
	}

	var show_list []Result
	for i := 0; i < len(list_order); i++ {
		new_array := Result{
			Name:      list_order[i].Name,
			Price:     list_order[i].Price,
			ProductID: list_order[i].ProductID,
			Qty:       list_order[i].Qty,
		}
		show_list = append(show_list, new_array)
	}

	t.Execute(&body, struct {
		Name       string
		TotalPrice int
		Product    []Result
	}{
		Name:       name_customer,
		TotalPrice: total_price,
		Product:    show_list,
	})

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent!")
}

/*
Riska
This function for send notif sms for seller
*/
func SendNotifSms(checkout_id uint, seller_id uint, seller_name string) {
	hp_seller, _ := database.GetHPSeller(checkout_id, seller_id)
	hp := strconv.FormatInt(hp_seller, 10)

	auth := vonage.CreateAuthFromKeySecret("3b4ad40f", "JxF6ztscitiJSqKO")
	smsClient := vonage.NewSMSClient(auth)
	response, errResp, err := smsClient.Send(hp, hp, "Hai. Ada pesanan baru dari pelanggan. Silahkan cek aplikasimu", vonage.SMSOpts{})

	if err != nil {
		fmt.Println(err)
		return
	}

	if response.Messages[0].Status == "0" {
		fmt.Println("Sms sent. Account Balance: " + response.Messages[0].RemainingBalance)
	} else {
		fmt.Println("Error code " + errResp.Messages[0].Status + ": " + errResp.Messages[0].ErrorText)
	}
}

//Fawwaz
//Finish a checkout
func FinishTransaction(c echo.Context) error {
	logged_customer_id := middlewares.ExtractToken(c)
	if logged_customer_id == 0 {
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
	checkout, err := database.EditCheckout(logged_customer_id, checkout_id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, checkout)
}
func FinishTransactionTesting() echo.HandlerFunc {
	return FinishTransaction
}
