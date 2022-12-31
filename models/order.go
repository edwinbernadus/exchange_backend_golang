package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	// Name     string
	// Students []Student

	CancelDate     time.Time
	IsFillComplete bool
	RequestRate    float64

	StartAmount       float64
	LeftAmount        float64
	TotalTransactions float64

	CurrencyPair string
	Uuid         string

	IsOpenOrder bool

	IsCancelled bool
	IsBuy       bool

	UserName string
}

func (input *Order) SubstractAmount(diffAmount float64) {
	input.LeftAmount = input.LeftAmount - diffAmount
	if input.LeftAmount == 0 {
		input.IsOpenOrder = false
		input.IsFillComplete = true
	}
}
