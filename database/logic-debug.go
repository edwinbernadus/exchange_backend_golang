package database

import (
	"fiff_golang_draft/helper"
	"fiff_golang_draft/models"
	"gorm.io/gorm/clause"
	"time"
)

func DebugWebApi() int64 {

	var db = helper.GetDb()

	company := &models.Company{Name: "company1"}
	db.Create(company)

	count := int64(0)
	db.Model(&models.Company{}).Count(&count)

	return count
}
func DebugTestSaveItem() {

	var db = helper.GetDb()

	company := &models.Company{Name: "company1"}
	db.Create(company)

	//student := models.Student{Name: "student1", Age: 21, Created: time.Now()}
	//company := &models.Company{Name: "company1"}
	//_ = append(company.Students, student)
	//db.Create(student)

}

func DebugTestSaveItem2() {

	var db = helper.GetDb()

	company := &models.Company{Name: "company2"}
	db.Create(company)

	// var student = &models.Student{Name: "student1", Age: 21, Created: time.Now(), CompanyID: company.ID}
	var student = &models.Student{Name: "student1", Age: 21, Created: time.Now()}
	// var student = &models.Student{Name: "student1", Age: 21, CompanyID: company.ID}
	//var student = &models.Student{Name: "student1", Age: 21}
	// var student = &models.Student{Name: "student1"}
	student.Company = *company
	// _ = append(company.Students, student)
	db.Create(student)

	//var tree = &models.Tree{Name: "tree1"}
	////_ = append(company.Students, student)
	//db.Create(tree)

}

func DebugTestSaveItem3a() {

	var db = helper.GetDb()

	var company models.Company
	db.First(&company, 1) // find product with integer primary key

	var student = &models.Student{Name: "student1a", Age: 21, Created: time.Now(), Company: company}
	//_ = append(company.Students, *student)
	db.Create(student)

}

func DebugTestLoadItem() {
	var db = helper.GetDb()

	//var student models.Student
	//db.First(&student, 1) // find product with integer primary key

	var company models.Company
	db.First(&company, 1) // find product with integer primary key
	//db.Preload(clause.Associations).Find(&users)
	db.Preload(clause.Associations).First(&company, 1)

}

func DebugTestLoadItem2() {
	var db = helper.GetDb()

	var student models.Student
	db.First(&student, 3) // find product with integer primary key

	var student2 models.Student
	db.Preload(clause.Associations).First(&student2, 7) // find product with integer primary key
}
