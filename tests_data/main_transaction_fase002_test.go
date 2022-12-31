package tests_data

import (
	"fiff_golang_draft/helper"
	"fiff_golang_draft/models"
	"fiff_golang_draft/services"
	"fmt"
	"gorm.io/gorm"
	"testing"
)

func TestMainTransaction2(t *testing.T) {
	helper.TestingMode = true
	deleteAll()

	initLastPrice()

	fmt.Println("sell list")
	generateSell()

	fmt.Println("buy list")
	generateBuy()
}

func initLastPrice() {
	var db = helper.GetDb()
	var tableInput = models.TableLastPrice{
		Model:        gorm.Model{},
		CurrencyPair: "btcusd",
		LastRate:     10000,
	}
	db.Create(&tableInput)
}

func generateSell() {
	var start = 10000
	for i := 1; i <= 10; i++ {
		var number = start + 500*i
		fmt.Println(number)

		var input = models.InputOrder{
			RequestRate:  float64(number),
			CurrencyPair: "btcusd",
			//Uuid:         uuid,
			Amount: 2,
			IsBuy:  false,
		}
		services.ExecuteTransaction(input)
	}
}

func generateBuy() {
	var start = 10000
	//for i := 1; i <= 10; i++ {
	for i := 1; i <= 1; i++ {
		var number = start - 500*i
		fmt.Println(number)

		var input = models.InputOrder{
			RequestRate:  float64(number),
			CurrencyPair: "btcusd",
			//Uuid:         uuid,
			Amount: 2,
			IsBuy:  true,
		}
		services.ExecuteTransaction(input)
	}
}

func deleteAll() {
	var db = helper.GetDb()
	db.Exec("DELETE FROM orders")
	db.Exec("DELETE FROM table_last_prices")
}
