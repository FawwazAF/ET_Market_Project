package controller

import (
	"etmarket/project/lib/database"

	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetAllCategoriesMarketIdController(c echo.Context) error {
	market_id, err := strconv.Atoi(c.Param("market_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid market id",
		})
	}
	allCategories, err := database.GetAllCategoriesMarketId(market_id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "category is not found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "success to get all food categories",
		"categories": allCategories,
	})
}

func GetCategoryNameMarketIdController(c echo.Context) error {
	market_id, err := strconv.Atoi(c.Param("market_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid market id",
		})
	}

	category_name := c.Param("category_name")
	selectedCategory, err := database.GetCategoryNameMarketId(market_id, category_name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "food category is not found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":       "success to get a food category",
		"category_name": selectedCategory,
	})
}
