package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"todo-backend-go/models"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// AutoMigrate: สร้างตารางตามโมเดล
	err = DB.AutoMigrate(&models.Todo{})
	if err != nil {
		panic("Failed to migrate database")
	}
}
