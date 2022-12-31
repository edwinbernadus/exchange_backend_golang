package helper

import (
	"fiff_golang_draft/models"
	"os"

	"regexp"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var TestingMode = false
var _instanceDb *gorm.DB = nil

func GetDbSqlite() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	var _ = db.AutoMigrate(&models.Product{}, &models.Student{}, &models.Company{}, &models.Bird{})
	var _ = db.AutoMigrate(&models.Tree{})
	var _ = db.AutoMigrate(&models.Order{})
	var _ = db.AutoMigrate(&models.TableLastPrice{})

	return db
}

func GetDb() *gorm.DB {

	if _instanceDb != nil {
		return _instanceDb
	}

	//dsn := "host=localhost user=postgres password=Testing1 dbname=golang_awesome port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var dsn = GetConnString()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	var _ = db.AutoMigrate(&models.Product{}, &models.Student{}, &models.Company{}, &models.Bird{})
	var _ = db.AutoMigrate(&models.Tree{})
	var _ = db.AutoMigrate(&models.Order{})
	var _ = db.AutoMigrate(&models.TableLastPrice{})

	_instanceDb = db
	return db
}

func GetEnvRailway() string {
	pgdatabase, _ := os.LookupEnv("PGDATABASE")
	return pgdatabase
}

func IsProdMode() string {
	var s = GetEnvRailway()
	if s == "" {
		return "false"
	} else {
		return "true"
	}
}

func GetConnString() string {
	if TestingMode {
		var dsn1 = "host=localhost user=edwin password=Testing1 dbname=golang_awesome_test port=5432 sslmode=disable"
		return dsn1
	}
	dsn := "host=localhost user=edwin password=Testing1 dbname=golang_awesome2 port=5432 sslmode=disable"

	var env1 = os.Getenv("CONN_STRING")
	if env1 != "" {
		dsn = env1
	}

	if (IsProdMode()) == "true" {
		// postgresql://postgres:0kFon0WFzPdavh1FymY7@containers-us-west-118.railway.app:6156/railway
		return "host=containers-us-west-118.railway.app user=postgres password=0kFon0WFzPdavh1FymY7 dbname=railway port=6156 sslmode=disable"
	}
	return dsn
}

func obfuscate(s string) string {
	// Compile the regular expressions.
	rUser, err := regexp.Compile(`user=\w+`)
	if err != nil {
		panic(err)
	}
	rPassword, err := regexp.Compile(`password=\w+`)
	if err != nil {
		panic(err)
	}

	// Replace the values of the user and password fields.
	modifiedString := rUser.ReplaceAllString(s, "user=[REDACTED]")
	modifiedString = rPassword.ReplaceAllString(modifiedString, "password=[REDACTED]")

	return modifiedString
}

func GetConnStringObfuscate() string {
	item1 := GetConnString()
	output := obfuscate(item1)
	return output

}
