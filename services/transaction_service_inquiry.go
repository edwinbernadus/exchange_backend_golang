package services

import (
	"fiff_golang_draft/helper"
	"fiff_golang_draft/models"

	"gorm.io/gorm"
)

type OrderOutput struct {
	LastRate     float64
	CurrencyPair string

	Buys  []models.Order
	Sells []models.Order
}

func getOrderList(pair string, limit1 int) OrderOutput {
	var db = helper.GetDb()
	var tableLastPrice models.TableLastPrice
	db.Transaction(func(tx *gorm.DB) error {
		tableLastPrice = findLastRateItem(pair, tx)
		return nil
	})

	var lastRate = tableLastPrice.LastRate

	var buys = getLastOrderBuyFromPair(pair, limit1)
	var sells = getLastOrderSellFromPair(pair, limit1)

	var output = OrderOutput{
		CurrencyPair: pair,
		LastRate:     lastRate,
		Buys:         buys,
		Sells:        sells,
	}
	return output
}

func getLastOrderBuyFromPair(pair string, limit1 int) []models.Order {
	var db = helper.GetDb()
	var result []models.Order
	var isBuy = true
	var isOpenOrder = true
	db.Order("request_rate desc").
		Where("is_buy = ? and is_open_order = ? and currency_pair = ?",
			isBuy, isOpenOrder, pair).
		Limit(limit1).Find(&result)
	return result
}

func getLastOrderSellFromPair(pair string, limit1 int) []models.Order {
	var db = helper.GetDb()
	var result []models.Order
	var isBuy = false
	var isOpenOrder = true
	db.Order("request_rate asc").
		Where("is_buy = ? and is_open_order = ? and currency_pair = ?",
			isBuy, isOpenOrder, pair).
		Limit(limit1).Find(&result)

	return result
}
