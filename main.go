package main

import (
	"log"
	"lab07.com/database"
	"lab07.com/models"
	"lab07.com/routes"
)

func main() {
	database.Connect()

	database.DB.AutoMigrate(&models.Book{})

	router := routes.SetupRouter()

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}