package database

import (
	"etmarket/project/config"
	"etmarket/project/middlewares"
	"etmarket/project/models"
)

/*
Author: Riska
This function for register customer
*/
func CreateSeller(seller models.Seller) (interface{}, error) {
	if err := config.DB.Save(&seller).Error; err != nil {
		return nil, err
	}

	//set output data
	output := map[string]interface{}{
		"id":    seller.ID,
		"email": seller.Email,
		"name":  seller.Name,
	}

	return output, nil
}

/*
Author: Riska
This function for login customer with matching data from database
*/
func LoginSeller(email string) (interface{}, error) {
	var seller models.Seller
	var err error
	if err = config.DB.Where("email = ?", email).First(&seller).Error; err != nil {
		return nil, err
	}
	seller.Token, err = middlewares.CreateToken(int(seller.ID))
	if err != nil {
		return nil, err
	}
	if err := config.DB.Save(seller).Error; err != nil {
		return nil, err
	}

	//set output data
	output := map[string]interface{}{
		"id":    seller.ID,
		"email": seller.Email,
		"token": seller.Token,
	}

	return output, nil
}

/*
Author: Riska
This function for search password user by email
*/
func GetPwdSeller(email string) string {
	var seller models.Seller
	config.DB.Where("email = ?", email).First(&seller)
	return seller.Password
}

/*
Author: Riska
This function for get token seller
*/
func GetTokenSeller(seller_id int) string {
	var seller models.Seller
	config.DB.Where("id = ?", seller_id).First(&seller)
	return seller.Token
}

/*
Author: Riska
This function for check is email customer exists
*/
func CheckEmailOnSeller(email string) (interface{}, error) {
	var seller models.Seller

	if err := config.DB.Model(&seller).Select("email").Where("email=?", email).First(&seller.Email).Error; err != nil {
		return nil, err
	}

	return seller, nil
}

/*
Author: Riska
This function for get 1 specified customer with interface output
*/
func GetSellerById(seller_id int) (interface{}, error) {
	var seller models.Seller
	if err := config.DB.Where("id=?", seller_id).First(&seller).Error; err != nil {
		return nil, err
	}
	//set output data
	output := map[string]interface{}{
		"id":     seller.ID,
		"email":  seller.Email,
		"name":   seller.Name,
		"alamat": seller.Address,
		"gender": seller.Gender,
	}

	return output, nil
}

/*
Author: Riska
This function for get 1 specified customer with Customer struct output
*/
func GetSeller(id int) (models.Seller, error) {
	var seller models.Seller
	if err := config.DB.Find(&seller, "id=?", id).Error; err != nil {
		return seller, err
	}
	return seller, nil
}

/*
Author: Riska
This function for get email customer
*/
func GetEmailSellerById(seller_id int) (string, error) {
	var seller models.Seller

	if err := config.DB.Model(&seller).Select("email").Where("id=?", seller_id).First(&seller.Email).Error; err != nil {
		return "nil", err
	}

	return seller.Email, nil
}

/*
Author: Riska
This function for update customer info from database
*/
func UpdateSeller(seller models.Seller) (interface{}, error) {
	if err := config.DB.Save(&seller).Error; err != nil {
		return nil, err
	}

	//set output data
	output := map[string]interface{}{
		"id":     seller.ID,
		"name":   seller.Name,
		"email":  seller.Email,
		"alamat": seller.Address,
		"gender": seller.Gender,
	}

	return output, nil
}

func GetAllSellerProduct(seller_id int) (interface{}, error) {
	var product []models.Product
	if err := config.DB.Find(&product, "seller_id = ?", seller_id).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func AddProductToSeller(product models.Product) (models.Product, error) {
	if err := config.DB.Save(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func EditSellerProduct(product_id int, seller_id int, product models.Product) (models.Product, error) {
	if err := config.DB.Where(&product, "id = ? AND seller_id = ?", product_id, seller_id).Save(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func GetSellerbyName(market_id int, category_name string) (interface{}, error) {
	var categories models.Category
	var seller []models.Seller
	search := "%" + category_name + "%"
	if err := config.DB.Find(&categories, "name LIKE ?", search).Error; err != nil {
		return nil, err
	}
	if err := config.DB.Find(&seller, "market_id=? AND category_id=?", market_id, categories.ID).Error; err != nil {
		return nil, err
	}
	return seller, nil
}

/*
Author: Riska
This function for get seller id
*/
func GetSellerIdByOderId(order_id int) (int, error) {
	var seller_id int
	var err error
	config.DB.Raw("SELECT products.seller_id FROM orders, products WHERE orders.product_id = products_id AND orders.id = ?", order_id).Scan(&seller_id)

	return seller_id, err
}

// func GetHPSeller(checkout_id, seller_id int) (interface{}, error) {
// 	fmt.Println(checkout_id)
// 	fmt.Println(seller_id)
// 	var seller models.Seller
// 	var err error
// 	// if err = config.DB.Model(&seller).Select("sellers.hp").Joins("join products on products.seller_id = sellers.id").Joins("join orders on products.id = orders.product_id").Joins("join checkout on orders.checkout_id = checkouts.id").Where("checkouts.status = ? AND checkouts.id = ? AND sellers.id = ?", "searching", checkout_id, seller_id).Group("checkouts.id").Error; err != nil {
// 	// 	return nil, err
// 	// }
// 	if err = config.DB.Raw("SELECT sellers.hp FROM sellers, products, orders, checkouts WHERE sellers.id = products.seller_id AND products.id = orders.product_id AND orders.checkout_id = checkouts.id AND checkouts.status = 'searching' AND checkouts.id = ? AND sellers.id = ? GROUP BY checkout_id", checkout_id, seller_id).Scan(&seller).Error; err != nil {
// 		return nil, err
// 	}
// 	fmt.Println(seller)
// 	return seller.Hp, err
// }

// func GetNameSeller(checkout_id, seller_id int) (interface{}, error) {
// 	var seller models.Seller
// 	var err error
// 	if err = config.DB.Raw("SELECT sellers.name FROM sellers, products, orders, checkouts WHERE sellers.id = products.seller_id AND products.id = orders.product_id AND orders.checkout_id = checkouts.id AND checkouts.status = 'searching' AND checkouts.id = ? AND sellers.id = ? GROUP BY checkout_id", checkout_id, seller_id).Scan(&seller).Error; err != nil {
// 		return nil, err
// 	}

// 	return seller.Name, err
// }
