package services

import (
	"fiff_golang_draft/helper"
	"fiff_golang_draft/models"
	"fmt"
	"gorm.io/gorm"
	"time"
)

func InitSample() {

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

func generateBuy() {
	var start = 10000
	for i := 1; i <= 10; i++ {
		//for i := 1; i <= 1; i++ {
		var number = start - 500*i
		fmt.Println(number)

		var input = models.InputOrder{
			RequestRate:  float64(number),
			CurrencyPair: "btcusd",
			//Uuid:         uuid,
			Amount: 2,
			IsBuy:  true,
		}
		ExecuteTransaction(input)
	}
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
		ExecuteTransaction(input)
	}
}

func deleteAll() {
	var db = helper.GetDb()
	db.Exec("DELETE FROM orders")
	db.Exec("DELETE FROM table_last_prices")
}

func obsoleteGetSample() []models.Order {
	orders := []models.Order{}

	{
		var item = models.Order{
			Model:             gorm.Model{},
			CancelDate:        time.Time{},
			IsFillComplete:    false,
			RequestRate:       10000,
			StartAmount:       1,
			LeftAmount:        1,
			TotalTransactions: 0,
			CurrencyPair:      "",
			Uuid:              "",
			IsOpenOrder:       true,
			IsCancelled:       false,
			IsBuy:             false,
		}
		orders = append(orders, item)
	}

	{
		var item = models.Order{
			Model:             gorm.Model{},
			CancelDate:        time.Time{},
			IsFillComplete:    false,
			RequestRate:       9000,
			StartAmount:       1,
			LeftAmount:        1,
			TotalTransactions: 0,
			CurrencyPair:      "",
			Uuid:              "",
			IsOpenOrder:       true,
			IsCancelled:       false,
			IsBuy:             false,
		}
		orders = append(orders, item)
	}
	{
		var item = models.Order{
			Model:             gorm.Model{},
			CancelDate:        time.Time{},
			IsFillComplete:    false,
			RequestRate:       8000,
			StartAmount:       1,
			LeftAmount:        1,
			TotalTransactions: 0,
			CurrencyPair:      "",
			Uuid:              "",
			IsOpenOrder:       true,
			IsCancelled:       false,
			IsBuy:             false,
		}
		orders = append(orders, item)
	}
	{
		var item = models.Order{
			Model:             gorm.Model{},
			CancelDate:        time.Time{},
			IsFillComplete:    false,
			RequestRate:       7000,
			StartAmount:       1,
			LeftAmount:        1,
			TotalTransactions: 0,
			CurrencyPair:      "",
			Uuid:              "",
			IsOpenOrder:       true,
			IsCancelled:       false,
			IsBuy:             false,
		}
		orders = append(orders, item)
	}
	{
		var item = models.Order{
			Model:             gorm.Model{},
			CancelDate:        time.Time{},
			IsFillComplete:    false,
			RequestRate:       6000,
			StartAmount:       1,
			LeftAmount:        1,
			TotalTransactions: 0,
			CurrencyPair:      "",
			Uuid:              "",
			IsOpenOrder:       true,
			IsCancelled:       false,
			IsBuy:             false,
		}
		orders = append(orders, item)
	}

	{
		var item = models.Order{
			Model:             gorm.Model{},
			CancelDate:        time.Time{},
			IsFillComplete:    false,
			RequestRate:       5000,
			StartAmount:       1,
			LeftAmount:        1,
			TotalTransactions: 0,
			CurrencyPair:      "",
			Uuid:              "",
			IsOpenOrder:       true,
			IsCancelled:       false,
			IsBuy:             true,
		}
		orders = append(orders, item)
	}
	{
		var item = models.Order{
			Model:             gorm.Model{},
			CancelDate:        time.Time{},
			IsFillComplete:    false,
			RequestRate:       4000,
			StartAmount:       1,
			LeftAmount:        1,
			TotalTransactions: 0,
			CurrencyPair:      "",
			Uuid:              "",
			IsOpenOrder:       true,
			IsCancelled:       false,
			IsBuy:             true,
		}
		orders = append(orders, item)
	}
	{
		var item = models.Order{
			Model:             gorm.Model{},
			CancelDate:        time.Time{},
			IsFillComplete:    false,
			RequestRate:       3000,
			StartAmount:       1,
			LeftAmount:        1,
			TotalTransactions: 0,
			CurrencyPair:      "",
			Uuid:              "",
			IsOpenOrder:       true,
			IsCancelled:       false,
			IsBuy:             true,
		}
		orders = append(orders, item)
	}
	{
		var item = models.Order{
			Model:             gorm.Model{},
			CancelDate:        time.Time{},
			IsFillComplete:    false,
			RequestRate:       2000,
			StartAmount:       1,
			LeftAmount:        1,
			TotalTransactions: 0,
			CurrencyPair:      "",
			Uuid:              "",
			IsOpenOrder:       true,
			IsCancelled:       false,
			IsBuy:             true,
		}
		orders = append(orders, item)
	}
	{
		var item = models.Order{
			Model:             gorm.Model{},
			CancelDate:        time.Time{},
			IsFillComplete:    false,
			RequestRate:       1000,
			StartAmount:       1,
			LeftAmount:        1,
			TotalTransactions: 0,
			CurrencyPair:      "",
			Uuid:              "",
			IsOpenOrder:       true,
			IsCancelled:       false,
			IsBuy:             true,
		}
		orders = append(orders, item)
	}
	return orders
}
