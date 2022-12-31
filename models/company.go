package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name     string
	Students []Student
}
