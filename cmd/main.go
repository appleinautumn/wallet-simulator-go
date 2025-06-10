package main

import (
	"fmt"
	"log"
	"os"

	"wallet-simulator-go/internal/db"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, relying on environment variables")
	}

	// Connect to database
	dbFile := os.Getenv("DB_FILE")
	if dbFile == "" {
		log.Fatal("Environment variable DB_FILE is not set")
	}

	// Connect to database
	database, err := db.NewConnection(dbFile)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	fmt.Println("Database connection established successfully")
}
