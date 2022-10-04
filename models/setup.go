package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(postgres.Open("host=localhost user=postgres password=123 dbname=api port=5432"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Schedule{}, &Activity{})

	DB = database
}
