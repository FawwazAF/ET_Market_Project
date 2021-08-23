package database

import (
	"etmarket/project/models"
	"testing"
)

var (
	mockDBHistory = models.Checkout{
		TotalQty: 10,
		TotalPrice: 10000,
		CustomerID: 1,
		PaymentID: 1,
		Status: "progress",
	}
	mockDBHistory2 = models.Checkout{
		TotalQty: 12,
		TotalPrice: 12000,
		CustomerID: 1,
		PaymentID: 2,
		Status: "progress",
	}
)

func TestGetHistoryInProgress(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator(),DropTable(&models.Checkout{})
	config.DB.Migrator().AutoMigrate(&models.Checkout{})
	config.DB.Save(&mockDBHistory)
	config.DB.Save(&mockDBHistory2)
	gotCheckoutProgress, err := GetHistoryInProgress("progress",1)
	if assert.NoError(t, err) {
		assert.Equal(t, 10, [0]gotCheckoutProgress.TotalQty)
		assert.Equal(t, 10000, [0]gotCheckoutProgress.TotalPrice)
		assert.Equal(t, 1, [0]gotCheckoutProgress.CustomerID)
		assert.Equal(t, 12, [1]gotCheckoutProgress.TotalQty)
		assert.Equal(t, 12000, [1]gotCheckoutProgress.TotalPrice)
		assert.Equal(t, 1, [1]gotCheckoutProgress.CustomerID)
	}
}
