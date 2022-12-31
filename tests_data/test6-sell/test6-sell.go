package test4_buy

import (
	"fiff_golang_draft/helper"
	"fiff_golang_draft/models"
	"fiff_golang_draft/services"
	"fmt"
	"testing"
)

func Execute(t *testing.T) {
	fmt.Println("test five - sell mode")
	helper.TestingMode = true
	deleteAllOrders()
	createBuyOrder()
	createBuyOrderTwo()
	executeSellOrder()
	checkTest(t)
}

func checkTest(t *testing.T) {
	// buy 1 - 1000
	// buy 1 - 2000
	// sell 0.5 - 500
	checkTotalRowsVisibleForTransaction(t)
	checkBuyOrder(t)
	checkSellOrder(t)
}

func checkBuyOrder(t *testing.T) {

	var db = helper.GetDb()

	//var result2 []*models.Order
	//db.Order("request_rate").
	//	Where("is_buy = ?", true).Find(&result2)

	//fmt.Println(result2)
	{
		var result1_2000 models.Order
		db.Order("request_rate desc").
			Where("is_buy = ?", true).First(&result1_2000)

		if result1_2000.RequestRate != 2000 {
			var msg = "buy order rate should 2000"
			t.Errorf(msg)
		}

		if result1_2000.IsOpenOrder == false {
			var msg = "buy order should open"
			t.Errorf(msg)
		}

		if result1_2000.LeftAmount != 1 {
			var msg = "buy order left amount should 0.5"
			t.Errorf(msg)
		}
		fmt.Println("item target - test6-sell", result1_2000)

		if result1_2000.IsFillComplete {
			var msg = "buy order should fill complete false"
			t.Errorf(msg)
		}

	}

	{
		var result1_1000 models.Order
		db.Order("request_rate").
			Where("is_buy = ?", true).First(&result1_1000)

		if result1_1000.RequestRate != 1000 {
			var msg = "[2] buy order rate should 1000"
			t.Errorf(msg)
		}

		if result1_1000.IsOpenOrder == false {
			var msg = "[2] buy order should open"
			t.Errorf(msg)
		}

		if result1_1000.LeftAmount != 0.5 {
			var msg = "[2] buy order left amount should 1"
			t.Errorf(msg)
		}

		if result1_1000.IsFillComplete {
			var msg = "[2] buy order should fill complete false"
			t.Errorf(msg)
		}
	}

}

func checkSellOrder(t *testing.T) {
	var result1 models.Order
	var db = helper.GetDb()
	db.Where("is_buy = ?", false).First(&result1)
	if result1.IsOpenOrder {
		var msg = "sell order should close"
		t.Errorf(msg)
	}

	if result1.LeftAmount != 0 {
		var msg = "sell order should left 0"
		t.Errorf(msg)
	}

	if result1.IsFillComplete == false {
		var msg = "sell order should fill complete"
		t.Errorf(msg)
	}
}

func checkTotalRowsVisibleForTransaction(t *testing.T) {
	var db = helper.GetDb()
	{
		var total int64
		var isBuy = true
		db.Model(&models.Order{}).Where("is_open_order = ? and is_buy = ?", true, isBuy).Count(&total)
		if total != 2 {
			var msg = "visible for buy transaction should 2"
			t.Errorf(msg)
		}
	}
	{
		var total int64
		var isBuy = false
		db.Model(&models.Order{}).Where("is_open_order = ? and is_buy = ?", true, isBuy).Count(&total)
		if total != 0 {
			var msg = "visible for sell transaction should 0"
			t.Errorf(msg)
		}
	}

}

func executeSellOrder() {
	fmt.Println("create order sell")
	var input = models.InputOrder{
		RequestRate:  500,
		CurrencyPair: "btcusd",
		//Uuid:         uuid,
		Amount: 0.5,
		IsBuy:  false,
	}
	services.ExecuteTransaction(input)
}

func createBuyOrder() {
	var input = models.InputOrder{
		RequestRate:  1000,
		CurrencyPair: "btcusd",
		//Uuid:         uuid,
		Amount: 1,
		IsBuy:  true,
	}
	fmt.Println("create order buy")
	services.ExecuteTransaction(input)

}

func createBuyOrderTwo() {
	var input = models.InputOrder{
		RequestRate:  2000,
		CurrencyPair: "btcusd",
		//Uuid:         uuid,
		Amount: 1,
		IsBuy:  true,
	}
	fmt.Println("create order buy - 2")
	services.ExecuteTransaction(input)

}

func deleteAllOrders() {
	var db = helper.GetDb()
	db.Exec("DELETE FROM orders")
}
