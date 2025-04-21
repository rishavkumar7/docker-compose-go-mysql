package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

func ConnectDb() *sql.DB {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v\n", err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging the database: %v\n", err)
		return nil
	}

	return db
}
