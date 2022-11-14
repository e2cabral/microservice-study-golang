package main

import (
	"github.com/gorilla/mux"
	"log"
	"microservice-golang/src/infra/database"
	"microservice-golang/src/main/config"
	"net/http"
)

func main() {
	app := mux.NewRouter()

	config.StartRouting(app)

	if err := database.InitializeDatabaseConnection(); err != nil {
		log.Fatal(err.Error())
	}

	if err := http.ListenAndServe(":3000", app); err != nil {
		log.Fatal(err)
	}
}
