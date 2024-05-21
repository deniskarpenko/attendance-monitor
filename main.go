package main

import (
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

	log.Printf(dbData.GetConnection())
	log.Printf("Connection:%s Port:%s Username:%s Password:%s Database:%s ", dbData.GetConnection(), dbData.GetPort(), dbData.GetUsername(), dbData.GetPassword())

}
