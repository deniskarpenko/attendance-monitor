package database

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type DbData struct {
	connection string
	port       string
	username   string
	password   string
}

func DBInstance() DbData {
	var data DbData
	
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	data.connection = os.Getenv("POSTGRES_HOST")
	data.port = os.Getenv("POSTGRES_PORT")
	data.username = os.Getenv("POSTGRES_USER")
	data.password = os.Getenv("POSTGRES_PASSWORD")

	log.Println("INNER" + os.Getenv("POSTGRES_HOST"))

	return data
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
