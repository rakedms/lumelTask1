package main

import (
	"log"
	"lumelTask1/database"
	"lumelTask1/loader"
	"lumelTask1/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load("./config/config.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	err := database.InitMongoDB()
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	csvFilePath := "salesData/sales.csv"

	err = loader.LoadData(csvFilePath)
	if err != nil {
		log.Fatalf("Failed to load CSV data: %v", err)
	}

	router := gin.Default()

	routes.SetupRoutes(router)

	router.Run(":8080")
}
