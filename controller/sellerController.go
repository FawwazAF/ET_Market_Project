package controller

import (
	"etmarket/project/lib/database"
	"etmarket/project/middlewares"
	"etmarket/project/models"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo"
)

/*
Author: Riska
This function for register seller
*/
func RegisterSeller(c echo.Context) error {
	//get user's input
	seller := models.Seller{}
	c.Bind(&seller)

	//check is email exists?
	is_email_exists, _ := database.CheckEmailOnSeller(seller.Email)
	if is_email_exists != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Email has already exist",
		})
	}

	//encrypt pass user
	convert_pwd := []byte(seller.Password) //convert pass from string to byte
	hashed_pwd := EncryptPwd(convert_pwd)
	seller.Password = hashed_pwd //set new pass

	//create new user
	data_seller, err := database.CreateSeller(seller)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	type Output struct {
		ID    uint   `json:"id"`
		Email string `json:"email"`
		Name  string `json:"name"`
	}

	//set output data
	output := Output{
		ID:    data_seller.ID,
		Email: data_seller.Email,
		Name:  data_seller.Name,
	}

	return c.JSON(http.StatusOK, output)
}

/*
Author: Riska
This function for login seller
*/
func LoginSeller(c echo.Context) error {
	//get user's input
	seller := models.Seller{}
	c.Bind(&seller)

	//compare password on form with db
	get_pwd := database.GetPwdSeller(seller.Email) //get password
	err := bcrypt.CompareHashAndPassword([]byte(get_pwd), []byte(seller.Password))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "User Unauthorized. Email or Password not equal",
		})
	}

	//login
	data_seller, err := database.LoginSeller(seller.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	type Output struct {
		ID    uint   `json:"id"`
		Email string `json:"email"`
		Token string `json:"token"`
	}

	//set output data
	output := Output{
		ID:    data_seller.ID,
		Email: data_seller.Email,
		Token: data_seller.Token,
	}

	return c.JSON(http.StatusOK, output)
}

/*
Author: Riska
This function for get profile seller
*/
func GetDetailSeller(c echo.Context) error {
	//convert id
	seller_id, err := strconv.Atoi(c.Param("seller_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid seller id",
		})
	}

	//check otorisasi
	logged_in_user_id := middlewares.ExtractToken(c)
	if logged_in_user_id != seller_id {
		return echo.NewHTTPError(http.StatusUnauthorized, "This user unauthorized to get detail")
	}

	//get data seller
	data_seller, err := database.GetSellerById(seller_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Cannot find seller",
		})
	}

	type Output struct {
		ID      uint   `json:"id"`
		Email   string `json:"email"`
		Name    string `json:"name"`
		Address string `json:"address"`
		Gender  string `json:"gender"`
		Hp      int64  `json:"hp"`
	}

	//set output data
	output := Output{
		ID:      data_seller.ID,
		Email:   data_seller.Email,
		Name:    data_seller.Name,
		Address: data_seller.Address,
		Gender:  data_seller.Gender,
		Hp:      data_seller.Hp,
	}

	return c.JSON(http.StatusOK, output)
}

func GetDetailSellerTesting() echo.HandlerFunc {
	return GetDetailSeller
}

/*
Author: Riska
This function for edit profile seller
*/
func UpdateSeller(c echo.Context) error {
	//convert id
	seller_id, err := strconv.Atoi(c.Param("seller_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	//check otorisasi
	logged_in_user_id := middlewares.ExtractToken(c)
	if logged_in_user_id != seller_id {
		return echo.NewHTTPError(http.StatusUnauthorized, "This user unauthorized to update data")
	}

	//get email seller
	email_seller, err := database.GetEmailSellerById(seller_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "cannot get data",
		})
	}

	//get seller
	seller, err := database.GetSeller(seller_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "cannot get data",
		})
	}
	c.Bind(&seller)

	//check email
	if seller.Email != email_seller {
		//check is email exists?
		is_email_exists, _ := database.CheckEmailOnSeller(seller.Email)
		if is_email_exists != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Email has already exist",
			})
		}
	}

	//encrypt pass user
	convert_pwd := []byte(seller.Password) //convert pass from string to byte
	hashed_pwd := EncryptPwd(convert_pwd)
	seller.Password = hashed_pwd //set new pass

	//update data seller
	updated_seller, err := database.UpdateSeller(seller)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "cannot update data",
		})
	}

	type Output struct {
		ID      uint   `json:"id"`
		Email   string `json:"email"`
		Name    string `json:"name"`
		Address string `json:"address"`
		Gender  string `json:"gender"`
		Hp      int64  `json:"hp"`
	}

	//set output data
	output := Output{
		ID:      updated_seller.ID,
		Email:   updated_seller.Email,
		Name:    updated_seller.Name,
		Address: updated_seller.Address,
		Gender:  updated_seller.Gender,
		Hp:      updated_seller.Hp,
	}

	return c.JSON(http.StatusOK, output)
}

func UpdateSellerTesting() echo.HandlerFunc {
	return UpdateSeller
}

// Ihsan
func GetSellerProducts(c echo.Context) error {
	// Auth
	seller_id := middlewares.ExtractToken(c)
	// Get data
	all_products_selected_seller, err := database.GetAllSellerProduct(seller_id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "no products in the shop",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "get all products from this shop success",
		"data":    all_products_selected_seller,
	})
}

// Ihsan
func AddProductToSeller(c echo.Context) error {
	// Auth
	seller_id := middlewares.ExtractToken(c)
	// http request body
	product := models.Product{}
	product.SellerID = uint(seller_id)
	c.Bind(&product)
	// Add data from request
	product_added, err := database.AddProductToSeller(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot add product",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    product_added,
	})
}

// ihsan
func EditSellerProduct(c echo.Context) error {
	seller_id := middlewares.ExtractToken(c)
	product_id, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid product id",
		})
	}

	product, err := database.GetEditProduct(product_id, seller_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "cannot get data",
		})
	}
	c.Bind(&product)

	product_edited, err := database.EditSellerProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot edit product",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    product_edited,
	})
}

/*
Author: Riska
This function for get all order by seller id
*/
func GetAllOrders(c echo.Context) error {
	logged_in_user_id := middlewares.ExtractToken(c)
	list_product, err := database.GetAllProductsOrderBySellerId(logged_in_user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot get list product",
		})
	}

	type ProductOnSeller struct {
		DriverName  string `json:"driver_name"`
		ProductName string `json:"product_name"`
		Qty         int    `json:"qty"`
		Price       int    `json:"price"`
	}

	var output []ProductOnSeller
	for i := 0; i < len(list_product); i++ {
		new_array := ProductOnSeller{
			DriverName:  list_product[i].DriverName,
			ProductName: list_product[i].ProductName,
			Qty:         list_product[i].Qty,
			Price:       list_product[i].Price,
		}
		output = append(output, new_array)
	}

	return c.JSON(http.StatusOK, output)
}

func GetAllOrdersTesting() echo.HandlerFunc {
	return GetAllOrders
}

/*
Author: Riska
This function for edit status item order
*/
func EditStatusItemOrder(c echo.Context) error {
	logged_in_user_id := middlewares.ExtractToken(c)

	order_id, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid order id",
		})
	}

	seller_id, err := database.GetSellerIdByOrderId(order_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot get seller id",
		})
	}

	//check otorisasi
	if logged_in_user_id != seller_id {
		return echo.NewHTTPError(http.StatusUnauthorized, "This user unauthorized edit status order")
	}

	order, err := database.GetItemOrderByOrderId(order_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot get order",
		})
	}

	order.Status = "completed"
	c.Bind(&order)

	update_status, err := database.EditItemStatus(order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot update status product",
		})
	}

	type Output struct {
		ID         uint   `json:"id"`
		CheckoutID uint   `json:"checkout_id"`
		ProductID  uint   `json:"product_id"`
		Qty        int    `json:"qty"`
		Price      int    `json:"price"`
		Status     string `json:"status"`
	}

	//set output data
	output := Output{
		ID:         update_status.ID,
		CheckoutID: update_status.CheckoutID,
		ProductID:  update_status.ProductID,
		Qty:        update_status.Qty,
		Price:      update_status.Price,
		Status:     update_status.Status,
	}

	return c.JSON(http.StatusOK, output)
}

func EditStatusItemOrderTesting() echo.HandlerFunc {
	return EditStatusItemOrder
}

/*
Author: Riska
This function for logout seller
*/
func LogoutSeller(c echo.Context) error {
	seller_id, err := strconv.Atoi(c.Param("seller_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	//check otorisasi
	logged_in_user_id := middlewares.ExtractToken(c)
	if logged_in_user_id != seller_id {
		return echo.NewHTTPError(http.StatusUnauthorized, "This user unauthorized to logout")
	}

	logout, err := database.GetSeller(seller_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "cannot get data",
		})
	}
	logout.Token = ""
	c.Bind(&logout)
	seller_updated, err := database.UpdateSeller(logout)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot logout",
		})
	}

	type Output struct {
		ID      uint   `json:"id"`
		Email   string `json:"email"`
		Name    string `json:"name"`
		Address string `json:"address"`
		Gender  string `json:"gender"`
		Hp      int64  `json:"hp"`
	}

	//set output data
	output := Output{
		ID:      seller_updated.ID,
		Email:   seller_updated.Email,
		Name:    seller_updated.Name,
		Address: seller_updated.Address,
		Gender:  seller_updated.Gender,
		Hp:      seller_updated.Hp,
	}

	return c.JSON(http.StatusOK, output)
}

func LogoutSellerTesting() echo.HandlerFunc {
	return LogoutSeller
}
