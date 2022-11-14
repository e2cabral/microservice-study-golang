package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"microservice-golang/src/infra/database"
	"microservice-golang/src/main/config"
)

func main() {
	if err := database.InitializeDatabaseConnection(); err != nil {
		log.Fatal(err.Error())
	}
	app := gin.Default()

	config.StartRouting(app)

	if err := app.Run(":3000"); err != nil {
		log.Fatal(err)
	}
}
