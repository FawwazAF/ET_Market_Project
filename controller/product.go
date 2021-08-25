package controller

import (
	"etmarket/project/lib/database"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// Get all product in a shop
func GetAllProductInShop(c echo.Context) error {
	seller_id, err := strconv.Atoi(c.QueryParam("seller_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid shop id",
		})
	}

	all_products, err := database.GetAllProductByShopId(seller_id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "no products in the shop",
		})
	}
	return c.JSON(http.StatusOK, all_products)
}

// Get specific product name in a shop
func GetSpecificProductInShop(c echo.Context) error {
	seller_id, err := strconv.Atoi(c.Param("seller_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid shop id",
		})
	}

	product_name := c.QueryParam("product_name")
	specific_product, err := database.GetSpecificProductByShopId(seller_id, product_name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "no product in this shop",
		})
	}
	return c.JSON(http.StatusOK, specific_product)
}

// Get specific product by id and seller id
func GetDetailSpecificProduct(c echo.Context) error {
	seller_id, err := strconv.Atoi(c.Param("seller_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id shop",
		})
	}
	product_id, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id product",
		})
	}
	specific_product, err := database.GetSpecificProductById(seller_id, product_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "no product in this shop",
		})
	}
	type Output struct {
		ID          uint   `json:"id"`
		Name        string `json:"name"`
		Stock       int    `json:"stock"`
		Price       int    `json:"price"`
		Description string `json:"description"`
		SellerID    uint   `json:"seller_id"`
	}

	//set output data
	output := Output{
		ID:          specific_product.ID,
		Name:        specific_product.Name,
		Stock:       specific_product.Stock,
		Price:       specific_product.Price,
		Description: specific_product.Description,
		SellerID:    specific_product.SellerID,
	}
	return c.JSON(http.StatusOK, output)
}
