package database

import (
	"errors"
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
	if len(carts) == 0 {
		return errors.New("cart is empty")
	}

	//iterative save from cart to order
	for _, v := range carts {
		order = models.Order{
			Qty:        v.Qty,
			Price:      v.Price,
			ProductID:  v.ProductID,
			CheckoutID: uint(checkout_id),
			Status:     "progress",
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
		PaymentID:  checkout.PaymentID,
		Status:     "searching",
	}
	if err := config.DB.Save(&new_checkout).Error; err != nil {
		return new_checkout, err
	}
	return new_checkout, nil
}

func GetManyPayment() ([]models.Payment, error) {
	var payments []models.Payment
	if err := config.DB.Find(&payments).Error; err != nil {
		return payments, err
	}
	return payments, nil
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
