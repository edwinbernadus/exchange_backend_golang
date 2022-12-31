package test1_buy

import (
	"fiff_golang_draft/helper"
	"fiff_golang_draft/models"
	"fiff_golang_draft/services"
	"fmt"
	"testing"
)

func Execute(t *testing.T) {
	fmt.Println("test one")
	helper.TestingMode = true
	deleteAllOrders()
	createSellOrder()
	executeBuyOrder()
	checkTest(t)
}

func checkTest(t *testing.T) {
	// sell 2
	// buy 2
	checkTotalRowsVisibleForTransaction(t)
	checkSellOrder(t)
	checkBuyOrder(t)
	checkLastPrice(t)
}

func checkLastPrice(t *testing.T) {
	var db = helper.GetDb()
	var result1 models.TableLastPrice
	db.Where("currency_pair = ?", "btcusd").First(&result1)
	if result1.LastRate != 1500.1 {
		var msg = "result1.LastRate is " + fmt.Sprintf("%f", result1.LastRate)
		t.Errorf(msg)
	}
}

func checkSellOrder(t *testing.T) {
	var result1 models.Order
	var db = helper.GetDb()
	db.Where("is_buy = ?", false).First(&result1)
	if result1.IsOpenOrder {
		var msg = "sell order still open"
		t.Errorf(msg)
	}

	if result1.LeftAmount != 0 {
		var msg = "sell order left amount not zero"
		t.Errorf(msg)
	}

	if result1.IsFillComplete == false {
		var msg = "sell order should fill complete"
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
	db.Model(&models.Order{}).Where("is_open_order = ?", true).Count(&total)
	//db.Model(&models.Order{}).Count(&total)
	if total != 0 {
		var msg = "visible for transaction more that 0"
		t.Errorf(msg)
	}
}

func executeBuyOrder() {
	var input = models.InputOrder{
		RequestRate:  2500.10,
		CurrencyPair: "btcusd",
		//Uuid:         uuid,
		Amount: 2,
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
