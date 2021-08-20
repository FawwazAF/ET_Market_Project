package database

import (
	"etmarket/project/config"
	"etmarket/project/models"
)

func InsertProductIntoCart(customer_id, seller_id, product_id int, carts models.Cart) (interface{}, error) {
	var product models.Product
	if err := config.DB.Find(&product, "id = ? AND seller_id = ?", product_id, seller_id).Error; err != nil {
		return nil, err
	}

	cart := models.Cart{
		CustomerID: uint(customer_id),
		ProductID:  uint(product_id),
		Qty:        carts.Qty,
		Price:      (product.Price * carts.Qty),
	}
	if err := config.DB.Save(&cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}

func GetAllCarts(customer_id int) (interface{}, error) {
	var carts []models.Cart
	if err := config.DB.Find(&carts, "customer_id= ?", customer_id).Error; err != nil {
		return nil, err
	}
	return carts, nil
}

func DeleteProductFromCart(customer_id, product_id int) (interface{}, error) {
	var carts models.Cart
	if err := config.DB.Find(&carts, "customer_id=? AND product_id = ?", customer_id, product_id).Error; err != nil {
		return nil, err
	}
	if err := config.DB.Delete(&carts, "customer_id=? AND product_id = ?", customer_id, product_id).Error; err != nil {
		return nil, err
	}
	return carts, nil
}
