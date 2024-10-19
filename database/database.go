package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "postgresql://book_management_db_adxq_user:8ew8XhWJ4XlTDAEhFllNK9TPDeYTvoIn@dpg-cs9lujrqf0us739iedsg-a.singapore-postgres.render.com/book_management_db_adxq"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connected successfully!")
}