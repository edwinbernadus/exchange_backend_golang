package services

import (
	"fiff_golang_draft/helper"
	"fiff_golang_draft/models"
	"time"
)

func StudentCreate() {
	var db = helper.GetDb()

	var company = models.Company{}
	db.Create(&company)
	var input = models.Student{
		Name:    "student1",
		Age:     1,
		Created: time.Time{},
		Company: company,
	}
	db.Create(&input)
}

func StudentGetList() []*models.Student {
	var db = helper.GetDb()
	var students []*models.Student
	db.Find(&students)
	return students
}

func StudentGetTotal() int64 {
	var db = helper.GetDb()
	count := int64(0)
	db.Model(&models.Student{}).Count(&count)
	return count
}
