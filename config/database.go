package config

import (
	"database/sql"
	"fmt"
	"log"
)

func LoadDatabase() *sql.DB {
	dbHost := "localhost"
	dbPort := "5432"
	dbUser := "postgres"
	dbPassword := "meongberem"
	dbName := "test-db"
	sslMode := "disable"

	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, dbPort, dbUser, dbPassword, dbName, sslMode)
	db, err := sql.Open("postgres", dataSourceName)

	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	} else {
		log.Println("Database Successfully Connected")
	}

	return db
}
