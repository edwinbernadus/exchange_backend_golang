package models

import (
	"gorm.io/gorm"
)

type TableLastPrice struct {
	gorm.Model
	CurrencyPair string
	LastRate     float64
}
