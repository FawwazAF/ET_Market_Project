package database

import (
	"etmarket/project/config"
	"etmarket/project/models"
)

func FindCarts(customer_id int) ([]models.Cart, error) {
	var carts []models.Cart
	if err := config.DB.Find(&carts, "customer_id=?", customer_id).Error; err != nil {
		return carts, err
	}
	return carts, nil
}

func CartMigrate(id, checkout_id int) error {
	var carts []models.Cart
	var order models.Order
	if err := config.DB.Find(&carts, "customer_id=?", id).Error; err != nil {
		return err
	}

	//iterative save from cart to order
	for _, v := range carts {
		order = models.Order{
			Qty:        v.Qty,
			Price:      v.Price,
			ProductID:  v.ProductID,
			CheckoutID: uint(checkout_id),
		}
		if err := config.DB.Save(&order).Error; err != nil {
			return err
		}
	}
	return nil
}

func Checkout(customer_id int, checkout models.Checkout, carts []models.Cart) (models.Checkout, error) {
	total_qty := 0
	total_price := 0
	for _, v := range carts {
		total_qty += v.Qty
		total_price += v.Price
	}
	new_checkout := models.Checkout{
		TotalQty:   total_qty,
		TotalPrice: total_price,
		CustomerID: uint(customer_id),
		DriverID:   checkout.DriverID,
		PaymentID:  checkout.PaymentID,
	}
	if err := config.DB.Save(&new_checkout).Error; err != nil {
		return new_checkout, err
	}
	return new_checkout, nil
}

func CheckoutUpdate(id int) error {
	//Find particular orders and checkout by checkout_id
	var orders []models.Order
	var checkout models.Checkout
	if err := config.DB.Find(&checkout, "id=?", id).Error; err != nil {
		return err
	}
	if err := config.DB.Find(&orders, "checkout_id=?", id).Error; err != nil {
		return err
	}

	//Iterative calculate total_qty and total_price
	total_qty := 0
	total_price := 0
	for _, v := range orders {
		total_qty += v.Qty
		total_price = total_price + (v.Qty * v.Price)
	}
	checkout = models.Checkout{
		TotalQty:   total_qty,
		TotalPrice: total_price,
		DriverID:   checkout.DriverID,
		PaymentID:  checkout.PaymentID,
	}

	if err := config.DB.Save(&checkout).Error; err != nil {
		return err
	}
	return nil
}

func DeleteCart(id int) error {
	var carts []models.Cart
	if err := config.DB.Find(&carts, "customer_id=?", id).Error; err != nil {
		return err
	}
	if err := config.DB.Unscoped().Delete(&carts, "customer_id=?", id).Error; err != nil {
		return err
	}
	return nil
}
