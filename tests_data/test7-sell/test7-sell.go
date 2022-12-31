package test4_buy

import (
	"fiff_golang_draft/helper"
	"fiff_golang_draft/models"
	"fiff_golang_draft/services"
	"fmt"
	"testing"
)

func Execute(t *testing.T) {
	fmt.Println("test four - sell mode")
	helper.TestingMode = true
	deleteAllOrders()
	createBuyOrder()
	createBuyOrderTwo()
	executeSellOrder()
	checkTest(t)
}

func checkTest(t *testing.T) {
	// buy 1 - 1500
	// buy 1 - 1600
	// sell 3
	checkTotalRowsVisibleForTransaction(t)
	checkBuyOrder(t)
	checkSellOrder(t)
}

func checkBuyOrder(t *testing.T) {

	var db = helper.GetDb()
	{
		var result1 models.Order
		db.Where("is_buy = ?", true).First(&result1)
		if result1.IsOpenOrder {
			var msg = "buy order should close"
			t.Errorf(msg)
		}

		if result1.LeftAmount != 0 {
			var msg = "buy order left amount should 0"
			t.Errorf(msg)
		}

		if result1.IsFillComplete == false {
			var msg = "buy order should fill complete"
			t.Errorf(msg)
		}
	}

	{
		var result1 models.Order
		db.Where("is_buy = ?", true).Last(&result1)
		if result1.IsOpenOrder {
			var msg = "[2] buy order should close"
			t.Errorf(msg)
		}

		if result1.LeftAmount != 0 {
			var msg = "[2] buy order left amount should 0"
			t.Errorf(msg)
		}

		if result1.IsFillComplete == false {
			var msg = "[2] buy order should fill complete"
			t.Errorf(msg)
		}
	}

}

func checkSellOrder(t *testing.T) {
	var result1 models.Order
	var db = helper.GetDb()
	db.Where("is_buy = ?", false).First(&result1)
	if result1.IsOpenOrder == false {
		var msg = "sell order should still open"
		t.Errorf(msg)
	}

	if result1.LeftAmount != 1 {
		var msg = "sell order should left 1"
		t.Errorf(msg)
	}

	if result1.IsFillComplete {
		var msg = "sell order should fill complete false"
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
		if total != 1 {
			var msg = "visible for sell transaction should 1"
			t.Errorf(msg)
		}
	}

}

func executeSellOrder() {
	fmt.Println("create order sell")
	var input = models.InputOrder{
		RequestRate:  500.10,
		CurrencyPair: "btcusd",
		//Uuid:         uuid,
		Amount: 3,
		IsBuy:  false,
	}
	services.ExecuteTransaction(input)
}

func createBuyOrder() {
	var input = models.InputOrder{
		RequestRate:  1500.10,
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
		RequestRate:  1600.10,
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
