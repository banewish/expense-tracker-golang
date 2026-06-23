package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func expenseDatabase() {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
		return
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		log.Println("Error connecting to the database:", err)
		return
	}

	// Ping the database to verify connection
	if err := db.Ping(); err != nil {
		log.Println("Unable to reach the database:", err)
		return
	}
	log.Println("Successfully connected and pinged the database!")

	// defer DB.Close() to close the database connection when the function exits
}
