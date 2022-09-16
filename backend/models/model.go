package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostTasks struct {
	ID   uint    `gorm:"primary key:autoIncrement" json:"id"`
	Data *string `json:"data"`
}

var DB *gorm.DB

func ConnecttoDatabase() {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "host=localhost user=postgres dbname=todo password=Gandhi@123 sslmode=disable",
	}))
	if err != nil {
		panic("Error:Failed to connect to database!")
	}

	db.AutoMigrate(&PostTasks{})

	DB = db
}
