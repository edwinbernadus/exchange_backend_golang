package services

import (
	"fiff_golang_draft/helper"
	"fiff_golang_draft/models"
	"math"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func ExecuteTransaction(input models.InputOrder) {

	var db = helper.GetDb()
	db.Transaction(func(tx *gorm.DB) error {
		var order = ConvertToOrder(input)
		tx.Create(&order)
		logic(order, tx)
		return nil
	})
}

func ConvertToOrder(input models.InputOrder) models.Order {
	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	var order = models.Order{
		CancelDate:        time.Time{},
		IsFillComplete:    false,
		RequestRate:       input.RequestRate,
		LeftAmount:        input.Amount,
		TotalTransactions: 0,
		CurrencyPair:      input.CurrencyPair,
		Uuid:              uuid,
		IsOpenOrder:       true,
		StartAmount:       input.Amount,
		IsCancelled:       false,
		IsBuy:             input.IsBuy,
		UserName:          input.UserName,
	}
	return order
}

func logic(input models.Order, tx *gorm.DB) {
	var matchOrder = findMatchOrder(input, tx)
	var isMatch = matchOrder.ID != 0
	if isMatch {
		settlement(&input, &matchOrder, tx)
		var currencyPair = matchOrder.CurrencyPair
		var lastPriceItem = findLastRateItem(currencyPair, tx)
		updateLastPrice(&lastPriceItem, matchOrder, tx)
		notificationLastPrice(lastPriceItem)
		notificationOrderList(currencyPair)
		if input.IsFillComplete == false {
			logic(input, tx)
		}
	}

}

func updateLastPrice(tableLastPrice *models.TableLastPrice, order models.Order, tx *gorm.DB) {
	tableLastPrice.LastRate = order.RequestRate
	tx.Save(&tableLastPrice)
}

func settlement(input *models.Order, matchOrder *models.Order, tx *gorm.DB) {
	var diffAmount = getDiffAmount(input.LeftAmount, matchOrder.LeftAmount)
	input.SubstractAmount(diffAmount)
	matchOrder.SubstractAmount(diffAmount)

	tx.Save(&input)
	tx.Save(&matchOrder)
}

func getDiffAmount(input float64, matchOrder float64) float64 {
	return math.Min(input, matchOrder)
}
