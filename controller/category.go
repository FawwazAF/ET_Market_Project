package controller

import (
	"etmarket/project/lib/database"

	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetAllCategories(c echo.Context) error {
	categories, err := database.GetAllCategories()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  categories,
	})
}

/*
Author: Patmiza
Getting all categories of seller in a market
*/
func GetAllCategoriesMarketIdController(c echo.Context) error {
	market_id, err := strconv.Atoi(c.Param("market_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid market id",
		})
	}
	all_shop, err := database.GetAllCategoriesMarketId(market_id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "shop is not found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "success to get all shops",
		"categories": all_shop,
	})
}

func GetSellerController(c echo.Context) error {
	market_id, err := strconv.Atoi(c.Param("market_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid market id",
		})
	}

	category_name := c.Param("category_name")
	list_seller, err := database.GetSellerbyName(market_id, category_name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "seller is not found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":     "success to get a seller",
		"list Seller": list_seller,
	})
}

/*
Author: Patmiza
Getting all sellers or shops in a market by category
*/
func GetCategoryNameMarketIdController(c echo.Context) error {
	market_id, err := strconv.Atoi(c.Param("market_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid market id",
		})
	}

	category_name := c.Param("category_name")
	list_seller, err := database.GetCategoryNameMarketId(market_id, category_name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "seller category is not found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":       "success to get a seller category",
		"category_name": list_seller,
	})
}
