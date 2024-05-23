package main

import (
	"database/sql"
	"github.com/deniskarpenko/attendance-monitor/database"
	"log"
)

func main() {
	/*port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	err := router.Run(":" + port)
	database.DBInstance()

	if err == nil {
		panic(fmt.Sprintf("Router error - %s", err))
	}*/
	dbData := database.DBInstance()

	dbData.Ping()
	createTable(dbData)
	defer dbData.Close()
}

func createTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS punches (id SERIAL NOT NULL PRIMARY KEY, email varchar(100))`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
