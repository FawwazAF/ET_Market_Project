package controller

import (
	"etmarket/project/lib/database"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// func Authorized(c echo.Context) (bool, models.User) {
// 	userId, token := middlewares.ExtractTokenUserId(c)
// 	userList, _ := database.GetOneUser(userId)

// 	if userList.Token != token {
// 		return false, userList
// 	}
// 	return true, userList
// }

// Get all product in a shop
func GetAllProductInShop(c echo.Context) error {
	// auth, userList := Authorized(c)
	// if auth == false {
	// 	return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access this account")
	// }
	shop_id, err := strconv.Atoi(c.Param("shop_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid shop id",
		})
	}
	all_products, err := database.GetAllProductByShopId(shop_id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "no products in the shop",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "get all products from this shop success",
		"data":    all_products,
	})
}

// Get specific product in a shop
func GetSpecificProductInShop(c echo.Context) error {
	// auth, userList := Authorized(c)
	// if auth == false {
	// 	return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access this account")
	// }
	shop_id, err := strconv.Atoi(c.Param("shop_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid shop id",
		})
	}
	product_name := c.Param("product_name")
	specific_product, err := database.GetSpecificProductByShopId(shop_id, product_name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "no product in this shop",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "get product from this shop success",
		"data":    specific_product,
	})
}

func GetDetailSpecificProduct(c echo.Context) error {
	// auth, userList := Authorized(c)
	// if auth == false {
	// 	return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access this account")
	// }
	shop_id, err := strconv.Atoi(c.Param("shop_id"))
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
	specific_product, err := database.GetSpecificProductById(shop_id, product_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "no product in this shop",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "get product from this shop success",
		"data":    specific_product,
	})
}
