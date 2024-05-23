package database

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type DbData struct {
	connection string
	port       string
	username   string
	password   string
	database   string
}

func DBInstance() *sql.DB {
	var data DbData

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	data.connection = os.Getenv("POSTGRES_HOST")
	data.port = os.Getenv("POSTGRES_PORT")
	data.username = os.Getenv("POSTGRES_USER")
	data.password = os.Getenv("POSTGRES_PASSWORD")
	data.database = os.Getenv("POSTGRES_DB")

	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", data.username, data.password, data.connection, data.database)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("INNER" + os.Getenv("POSTGRES_HOST"))

	return db
}

func (db *DbData) GetConnection() string {
	return db.connection
}

func (db *DbData) GetPort() string {
	return db.port
}

func (db *DbData) GetUsername() string {
	return db.username
}

func (db *DbData) GetPassword() string {
	return db.password
}
