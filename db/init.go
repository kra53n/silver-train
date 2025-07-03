package db

import (
	// "fmt"

	"gorm.io/gorm"
	//"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"

	"silver-train/model/auth"
)

var DB *gorm.DB

func Connect() {
	// var dsn string
	var err error
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable"
	// db, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	DB, _ = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(&authModel.RefreshToken{})
}
