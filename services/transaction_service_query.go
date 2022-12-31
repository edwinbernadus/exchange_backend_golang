package services

import (
	"fiff_golang_draft/helper"
	"fiff_golang_draft/models"

	"gorm.io/gorm"
)

func InitParamDb() {
	var db = helper.GetDb()

	var tableLastPrice models.TableLastPrice
	db.First(&tableLastPrice)

	if tableLastPrice.ID == 0 {
		var tableInput = models.TableLastPrice{
			Model:        gorm.Model{},
			CurrencyPair: "btcusd",
			LastRate:     0,
		}
		db.Create(&tableInput)
	}
}

func findLastRateItem(currencyPair string, tx *gorm.DB) models.TableLastPrice {
	var result models.TableLastPrice
	tx.Where("currency_pair = ?", currencyPair).First(&result)
	return result
}

func findMatchOrder(input models.Order, tx *gorm.DB) models.Order {
	var isBuy = !input.IsBuy
	var result models.Order
	if isBuy {
		result = findMatchOrderBuy(input, tx)
	} else {
		result = findMatchOrderSell(input, tx)
	}
	return result
}

func findMatchOrderBuy(input models.Order, tx *gorm.DB) models.Order {
	var result models.Order

	var isBuy = true
	var isOpenOrder = true
	var requestRate = input.RequestRate
	var orderBy = "request_rate"

	var tempList []models.Order
	tx.
		Where("is_buy = ? and is_open_order = ? and request_rate >= ?", isBuy, isOpenOrder, requestRate).
		Find(&tempList)

	tx.Order(orderBy).
		Where("is_buy = ? and is_open_order = ? and request_rate >= ?", isBuy, isOpenOrder, requestRate).
		First(&result)
	return result
}

func findMatchOrderSell(input models.Order, tx *gorm.DB) models.Order {

	var result models.Order

	var isBuy = false
	var isOpenOrder = true
	var requestRate = input.RequestRate
	var orderBy = "request_rate desc"

	var tempList []models.Order
	tx.
		Where("is_buy = ? and is_open_order = ? and request_rate <= ?", isBuy, isOpenOrder, requestRate).
		Find(&tempList)

	tx.Order(orderBy).
		Where("is_buy = ? and is_open_order = ? and request_rate <= ?", isBuy, isOpenOrder, requestRate).
		First(&result)
	return result

	//var result models.Order
	//var isOpenOrder = true
	//var isBuy = !input.IsBuy
	//
	//var tempList []models.Order
	//tx.Where("is_buy = ? and is_open_order = ?", isBuy, isOpenOrder).Find(&tempList)
	//
	//var orderBy = "request_rate desc"
	//if isBuy == false {
	//	orderBy = "request_rate"
	//}
	//tx.Order(orderBy).Where("is_buy = ? and is_open_order = ?", isBuy, isOpenOrder).First(&result)
	//return result
}
