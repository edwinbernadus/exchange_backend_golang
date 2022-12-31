package test2_buy

import (
	"fiff_golang_draft/helper"
	"fiff_golang_draft/models"
	"fiff_golang_draft/services"
	"fmt"
	"testing"
)

func Execute(t *testing.T) {
	fmt.Println("test two")
	helper.TestingMode = true
	deleteAllOrders()
	createSellOrder()
	executeBuyOrder()
	checkTest(t)
}

func checkTest(t *testing.T) {
	// sell 2
	// buy 1
	checkTotalRowsVisibleForTransaction(t)
	checkSellOrder(t)
	checkBuyOrder(t)
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
		var msg = "sell order left amount should 1"
		t.Errorf(msg)
	}

	if result1.IsFillComplete {
		var msg = "sell order should fill complete false"
		t.Errorf(msg)
	}
}

func checkBuyOrder(t *testing.T) {
	var result1 models.Order
	var db = helper.GetDb()
	db.Where("is_buy = ?", true).First(&result1)
	if result1.IsOpenOrder {
		var msg = "buy order still open"
		t.Errorf(msg)
	}

	if result1.LeftAmount != 0 {
		var msg = "buy order left amount not zero"
		t.Errorf(msg)
	}

	if result1.IsFillComplete == false {
		var msg = "buy order should fill complete"
		t.Errorf(msg)
	}
}

func checkTotalRowsVisibleForTransaction(t *testing.T) {
	var db = helper.GetDb()
	//var total = int64(0)
	var total int64
	db.Model(&models.Order{}).Where("is_open_order = ? and is_buy = ?", true, false).Count(&total)
	//db.Model(&models.Order{}).Count(&total)
	if total != 1 {
		var msg = "visible for transaction should 1"
		t.Errorf(msg)
	}
}

func executeBuyOrder() {
	var input = models.InputOrder{
		RequestRate:  2500.10,
		CurrencyPair: "btcusd",
		//Uuid:         uuid,
		Amount: 1,
		IsBuy:  true,
	}
	services.ExecuteTransaction(input)
}

func createSellOrder() {
	var input = models.InputOrder{
		RequestRate:  1500.10,
		CurrencyPair: "btcusd",
		//Uuid:         uuid,
		Amount: 2,
		IsBuy:  false,
	}
	services.ExecuteTransaction(input)
}

func deleteAllOrders() {
	var db = helper.GetDb()
	db.Exec("DELETE FROM orders")
}
