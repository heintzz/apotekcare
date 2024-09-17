package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func ConnectPostgres(host, port, user, password, dbname string) (*sql.DB, error) {
	db, err := getPostgres(host, port, user, password, dbname)
	if err != nil {
			return nil, err
	}

	err = db.Ping()
	if err != nil {
			return nil, err
	}

	return db, nil
}

func getPostgres(host, port, user, password, dbname string) (*sql.DB, error) {	
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := createConnection(connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func createConnection(connectionString string) (*sql.DB, error) {	
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(25)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}

