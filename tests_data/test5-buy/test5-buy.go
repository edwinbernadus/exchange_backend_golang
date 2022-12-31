package test4_buy

import (
	"fiff_golang_draft/helper"
	"fiff_golang_draft/models"
	"fiff_golang_draft/services"
	"fmt"
	"testing"
)

func Execute(t *testing.T) {
	fmt.Println("test five - buy mode")
	helper.TestingMode = true
	deleteAllOrders()
	createSellOrder()
	createSellOrderTwo()
	executeBuyOrder()
	checkTest(t)
}

func checkTest(t *testing.T) {
	// sell 1 - 1000
	// sell 1 - 2000
	// buy 0.5 - 3000
	checkTotalRowsVisibleForTransaction(t)
	checkSellOrder(t)
	checkBuyOrder(t)
}

func checkSellOrder(t *testing.T) {

	var db = helper.GetDb()

	var result2 []*models.Order
	db.Order("request_rate").
		Where("is_buy = ?", false).Find(&result2)

	fmt.Println(result2)
	{
		var result1First models.Order
		db.Order("request_rate desc").
			Where("is_buy = ?", false).First(&result1First)

		if result1First.RequestRate != 2000 {
			var msg = "sell order rate should 2000"
			t.Errorf(msg)
		}

		if result1First.IsOpenOrder == false {
			var msg = "sell order should open"
			t.Errorf(msg)
		}

		if result1First.LeftAmount != 0.5 {
			var msg = "sell order left amount should 0.5"
			t.Errorf(msg)
		}
		fmt.Println("item target - test5-buy", result1First)

		if result1First.IsFillComplete {
			var msg = "sell order should fill complete false"
			t.Errorf(msg)
		}

	}

	{
		var result1Last models.Order
		db.Order("request_rate ").
			Where("is_buy = ?", false).First(&result1Last)

		if result1Last.RequestRate != 1000 {
			var msg = "[2] sell order rate should 1000"
			t.Errorf(msg)
		}

		if result1Last.IsOpenOrder == false {
			var msg = "[2] sell order should open"
			t.Errorf(msg)
		}

		if result1Last.LeftAmount != 1 {
			var msg = "[2] sell order left amount should 1"
			t.Errorf(msg)
		}

		if result1Last.IsFillComplete {
			var msg = "[2] sell order should fill complete false"
			t.Errorf(msg)
		}
	}

}

func checkBuyOrder(t *testing.T) {
	var result1 models.Order
	var db = helper.GetDb()
	db.Where("is_buy = ?", true).First(&result1)
	if result1.IsOpenOrder {
		var msg = "buy order should close"
		t.Errorf(msg)
	}

	if result1.LeftAmount != 0 {
		var msg = "buy order should left 0"
		t.Errorf(msg)
	}

	if result1.IsFillComplete == false {
		var msg = "buy order should fill complete"
		t.Errorf(msg)
	}
}

func checkTotalRowsVisibleForTransaction(t *testing.T) {
	var db = helper.GetDb()

	{
		var total int64
		var isBuy = true
		db.Model(&models.Order{}).Where("is_open_order = ? and is_buy = ?", true, isBuy).Count(&total)
		if total != 0 {
			var msg = "visible for buy transaction should 0"
			t.Errorf(msg)
		}
	}

	{
		var total int64
		var isBuy = false
		db.Model(&models.Order{}).Where("is_open_order = ? and is_buy = ?", true, isBuy).Count(&total)

		var items []*models.Order
		db.Model(&models.Order{}).Where("is_open_order = ? and is_buy = ?", true, isBuy).Find(&items)

		if total != 2 {
			var msg = "visible for sell transaction should 2"
			t.Errorf(msg)
		}
	}

}

func executeBuyOrder() {
	fmt.Println("create order buy")
	var input = models.InputOrder{
		RequestRate:  3000,
		CurrencyPair: "btcusd",
		//Uuid:         uuid,
		Amount: 0.5,
		IsBuy:  true,
	}
	services.ExecuteTransaction(input)
}

func createSellOrder() {
	var input = models.InputOrder{
		RequestRate:  1000,
		CurrencyPair: "btcusd",
		//Uuid:         uuid,
		Amount: 1,
		IsBuy:  false,
	}
	fmt.Println("create order sell")
	services.ExecuteTransaction(input)

}

func createSellOrderTwo() {
	var input = models.InputOrder{
		RequestRate:  2000,
		CurrencyPair: "btcusd",
		//Uuid:         uuid,
		Amount: 1,
		IsBuy:  false,
	}
	fmt.Println("create order sell - 2")
	services.ExecuteTransaction(input)

}

func deleteAllOrders() {
	var db = helper.GetDb()
	db.Exec("DELETE FROM orders")
}
