package main

import (
	"heintzz/apotekcare/apps"
	"heintzz/apotekcare/external/database"
	"heintzz/apotekcare/internal/utils"
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

	utils.InitToken(os.Getenv("JWT_SECRET_KEY"), 60) 
	
	db, err := database.ConnectPostgres(host, port, user, password, dbname)
	if err != nil {
			log.Fatalf("Error connecting to database: %v", err)
	}

	if db != nil {
		log.Println("DB connected!")
	}

	
	appPort := ":8080"
	apps.RunServer(appPort, db)
}

