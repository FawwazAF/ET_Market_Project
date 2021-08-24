package controller

import (
	"etmarket/project/lib/database"

	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

/*
Author: Riska
Getting all categories for register seller
*/
func GetAllCategories(c echo.Context) error {
	categories, err := database.GetAllCategories()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, categories)
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
			"message": "seller is not found",
		})
	}
	type Result struct {
		Name       string `json:"name"`
		MarketID   uint   `json:"market_id"`
		CategoryID uint   `json:"category_id"`
	}

	var output []Result
	for i := 0; i < len(all_shop); i++ {
		new_result := Result{
			Name:       all_shop[i].Name,
			MarketID:   all_shop[i].MarketID,
			CategoryID: all_shop[i].CategoryID,
		}
		output = append(output, new_result)
	}
	return c.JSON(http.StatusOK, output)
}

// func GetSellerController(c echo.Context) error {
// 	market_id, err := strconv.Atoi(c.Param("market_id"))
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{
// 			"message": "invalid market id",
// 		})
// 	}

// 	category_name := c.Param("category_name")
// 	list_seller, err := database.GetSellerbyName(market_id, category_name)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 			"message": "seller is not found",
// 		})
// 	}
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message":     "success to get a seller",
// 		"list Seller": list_seller,
// 	})
// }

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
	type Result struct {
		ID           uint   `json:"id"`
		Name         string `json:"name"`
		CategoryName string `json:"category_name"`
	}
	var output []Result
	for i := 0; i < len(list_seller); i++ {
		new_result := Result{
			ID:           list_seller[i].ID,
			Name:         list_seller[i].Name,
			CategoryName: list_seller[i].Category,
		}
		output = append(output, new_result)
	}
	return c.JSON(http.StatusOK, output)
}
