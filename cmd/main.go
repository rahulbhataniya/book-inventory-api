package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/rahulbhataniya/book-inventory-api/internal/config"
	"github.com/rahulbhataniya/book-inventory-api/internal/model"
	"github.com/rahulbhataniya/book-inventory-api/internal/routes"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	config.ConnectDB()
	config.DB.AutoMigrate(&model.Book{})

	r := routes.SetupRouter()

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed:", err)
	}
}
