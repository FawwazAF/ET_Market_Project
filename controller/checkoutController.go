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

	"github.com/labstack/echo"
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

	//Migrate data from cart to order
	if err := database.CartMigrate(logged_in_user_id, int(new_checkout.ID)); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Hard Delete cart
	if err := database.DeleteCart(logged_in_user_id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	SendEmail(new_checkout.ID, logged_in_user_id)
	// SendNotifSms(int(new_checkout.ID), logged_in_user_id)
	return c.JSON(http.StatusOK, new_checkout)
}

func SendEmail(checkout_id uint, customer_id int) {
	email_customer, _ := database.GetEmailCustomerById(customer_id)
	customer, _ := database.GetCustomer(customer_id)
	name_customer := customer.Name
	total_price := database.GetTotalPrice(checkout_id)

	list_order, _ := database.GetListOrderOnCheckoutForCustomer(checkout_id)
	fmt.Println(list_order)
	fmt.Println("len", len(list_order))
	// fmt.Println(total_price)

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

	fmt.Println("show_list \n", show_list)

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

// func SendNotifSms(checkout_id, seller_id int) {
// 	fmt.Println("checkout_id ", checkout_id)
// 	fmt.Println("seller_id ", seller_id)
// 	hp_seller, _ := database.GetHPSeller(checkout_id, seller_id)
// 	name_seller, _ := database.GetNameSeller(checkout_id, seller_id)

// 	fmt.Println("hp_seller ", hp_seller)
// 	fmt.Println("name_seller ", name_seller)
// 	// 	// auth := vonage.CreateAuthFromKeySecret("3b4ad40f", "JxF6ztscitiJSqKO")
// 	// 	// smsClient := vonage.NewSMSClient(auth)
// 	// 	// response, errResp, err := smsClient.Send("6281296645463", "6281211732575", "Hello, test test. This is riska, try to send sms via nexmo", vonage.SMSOpts{})

// 	// 	// if err != nil {
// 	// 	// 	panic(err)
// 	// 	// }

// 	// 	// if response.Messages[0].Status == "0" {
// 	// 	// 	fmt.Println("Account Balance: " + response.Messages[0].RemainingBalance)
// 	// 	// } else {
// 	// 	// 	fmt.Println("Error code " + errResp.Messages[0].Status + ": " + errResp.Messages[0].ErrorText)
// 	// 	// }
// }
