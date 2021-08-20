package database

import (
	"etmarket/project/config"
	"etmarket/project/models"
)

/*
Author: Patmiza
Getting all categories of seller in a market
*/
func GetAllProgressOrders(driver_id int) (interface{}, error) {
	var orders []models.Order
	if err := config.DB.Find(&orders, "driver_id = ? AND status = ?", driver_id, "progress").Error; err != nil {
		return nil, err
	}
	return orders, nil
}

/*
Author: Riska
This function for get all products on order table by user login
*/
func GetAllProductsOrderBySellerId(seller_id int) (interface{}, error) {
	var results []map[string]interface{}
	var err error
	if err = config.DB.Raw("SELECT drivers.name as driver_name, products.name as product_name, orders.qty, orders.price FROM sellers, products, orders, checkouts, deliveries, drivers WHERE sellers.id = products.seller_id AND products.id = orders.product_id AND orders.checkout_id = checkouts.id AND checkouts.id = deliveries.checkout_id AND orders.status = 'progress' AND deliveries.driver_id = drivers.id AND sellers.id = ?", seller_id).Scan(&results).Error; err != nil {
		return nil, err
	}

	return results, err
}

/*
Author: Riska
This function for get item order
*/
func GetItemOrderByOrderId(order_id int) (models.Order, error) {
	var order models.Order
	if err := config.DB.Find(&order, "id=?", order_id).Error; err != nil {
		return order, err
	}
	return order, nil
}

/*
Author: Riska
This function for update item status on table order
*/
func EditItemStatus(order models.Order) (interface{}, error) {
	if err := config.DB.Save(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}
