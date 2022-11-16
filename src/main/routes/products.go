package routes

import (
	"github.com/gorilla/mux"
	"microservice-golang/src/infra/helpers"
	"microservice-golang/src/presentation/controllers"
)

func SetProductsRoutes(app *mux.Router) {
	router := helpers.NewRouter("/products")

	sub := app.PathPrefix(router.Prefix).Subrouter()

	sub.HandleFunc("", controllers.Find).Methods("GET")
	sub.HandleFunc("", controllers.Create).Methods("POST")
	sub.HandleFunc("/{id:[0-9]+}", controllers.Update).Methods("PATCH")
	sub.HandleFunc("/{id:[0-9]+}", controllers.FindOne).Methods("GET")
}
