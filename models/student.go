package models

import (
	"gorm.io/gorm"
	"time"
)

type Student struct {
	gorm.Model
	Name      string
	Age       int
	Created   time.Time
	Company   Company
	CompanyID uint
}

type Tree struct {
	gorm.Model
	Name string
}
