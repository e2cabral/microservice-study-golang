package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"microservice-golang/src/main/config"
)

func main() {
	app := gin.Default()

	config.StartRouting(app)

	if err := app.Run(":3000"); err != nil {
		log.Fatal(err)
	}
}
