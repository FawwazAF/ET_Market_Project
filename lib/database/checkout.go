package database

import (
	"errors"
	"etmarket/project/config"
	"etmarket/project/models"
)

func FindCarts(customer_id int) ([]models.Cart, error) {
	var carts []models.Cart
	var product []models.Product
	if err := config.DB.Find(&carts, "customer_id=?", customer_id).Error; err != nil {
		return carts, err
	}
	for i, v := range carts {
		if err := config.DB.Find(&product, "id=?", v.ProductID).Error; err != nil {
			return carts, err
		}
		if product[i].Stock < v.Qty {
			return carts, errors.New("stock is empty")
		}
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
	for i, v := range carts {
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
		if err := AutoUpdateStock(carts[i]); err != nil {
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

func AutoUpdateStock(cart models.Cart) error {
	var product models.Product
	if err := config.DB.Find(&product, "id=?", cart.ProductID).Error; err != nil {
		return err
	}
	product.Stock = product.Stock - cart.Qty
	if err := config.DB.Save(product).Error; err != nil {
		return err
	}
	return nil
}
