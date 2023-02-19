package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Load environment variables from local.env file
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatal(err)
	}

	// Pull the database connection parameters from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")

	// Construct the database connection string
	dbInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=require",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// Connect to the database
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS budgets (
			id SERIAL PRIMARY KEY,
			budget_display_name TEXT,
			alert_threshold_exceeded FLOAT,
			cost_amount FLOAT,
			cost_interval_start TIMESTAMP,
			budget_amount FLOAT,
			budget_amount_type TEXT,
			currency_code TEXT
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Budgets table created successfully")
}
