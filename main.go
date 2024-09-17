package main

import (
	"heintzz/ecommerce/apps"
	"heintzz/ecommerce/external/database"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("error when load env file with error", err.Error())
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	
	db, err := database.ConnectPostgres(host, port, user, password, dbname)
	if err != nil {
			log.Fatalf("Error connecting to database: %v", err)
	}

	if db != nil {
		log.Println("DB connected!")
	}

	
	appPort := ":8080"
	apps.RunServer(appPort)
}

