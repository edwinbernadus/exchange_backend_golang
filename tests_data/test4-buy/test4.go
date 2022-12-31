package test4_buy

import (
	"fiff_golang_draft/helper"
	"fiff_golang_draft/models"
	"fiff_golang_draft/services"
	"fmt"
	"testing"
)

func Execute(t *testing.T) {
	fmt.Println("test four")
	helper.TestingMode = true
	deleteAllOrders()
	createSellOrder()
	createSellOrder()
	executeBuyOrder()
	checkTest(t)
}

func checkTest(t *testing.T) {
	// sell 1
	// sell 1
	// buy 3
	checkTotalRowsVisibleForTransaction(t)
	checkSellOrder(t)
	checkBuyOrder(t)
}

func checkSellOrder(t *testing.T) {

	var db = helper.GetDb()
	{
		var result1 models.Order
		db.Where("is_buy = ?", false).First(&result1)
		if result1.IsOpenOrder {
			var msg = "sell order should close"
			t.Errorf(msg)
		}

		if result1.LeftAmount != 0 {
			var msg = "sell order left amount should 0"
			t.Errorf(msg)
		}

		if result1.IsFillComplete == false {
			var msg = "sell order should fill complete"
			t.Errorf(msg)
		}
	}

	{
		var result1 models.Order
		db.Where("is_buy = ?", false).Last(&result1)
		if result1.IsOpenOrder {
			var msg = "[2] sell order should close"
			t.Errorf(msg)
		}

		if result1.LeftAmount != 0 {
			var msg = "[2] sell order left amount should 0"
			t.Errorf(msg)
		}

		if result1.IsFillComplete == false {
			var msg = "[2] sell order should fill complete"
			t.Errorf(msg)
		}
	}

}

func checkBuyOrder(t *testing.T) {
	var result1 models.Order
	var db = helper.GetDb()
	db.Where("is_buy = ?", true).First(&result1)
	if result1.IsOpenOrder == false {
		var msg = "buy order should still open"
		t.Errorf(msg)
	}

	if result1.LeftAmount != 1 {
		var msg = "buy order should left 1"
		t.Errorf(msg)
	}

	if result1.IsFillComplete {
		var msg = "buy order should fill complete false"
		t.Errorf(msg)
	}
}

func checkTotalRowsVisibleForTransaction(t *testing.T) {
	var db = helper.GetDb()
	{
		var total int64
		var isBuy = true
		db.Model(&models.Order{}).Where("is_open_order = ? and is_buy = ?", true, isBuy).Count(&total)
		if total != 1 {
			var msg = "visible for buy transaction should 1"
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

func executeBuyOrder() {
	fmt.Println("create order buy")
	var input = models.InputOrder{
		RequestRate:  2500.10,
		CurrencyPair: "btcusd",
		//Uuid:         uuid,
		Amount: 3,
		IsBuy:  true,
	}
	services.ExecuteTransaction(input)
}

func createSellOrder() {
	var input = models.InputOrder{
		RequestRate:  1500.10,
		CurrencyPair: "btcusd",
		//Uuid:         uuid,
		Amount: 1,
		IsBuy:  false,
	}
	fmt.Println("create order sell")
	services.ExecuteTransaction(input)

}

func deleteAllOrders() {
	var db = helper.GetDb()
	db.Exec("DELETE FROM orders")
}
