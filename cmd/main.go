package main

import (
	"fmt"
	"log"

	"wallet-simulator-go/internal/db"
)

func main() {
	// Connect to SQLite database
	database, err := db.ConnectSQLite("wallet.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	fmt.Println("Database connection established successfully")
}
