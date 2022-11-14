package facades

import (
	"github.com/gorilla/mux"
	"microservice-golang/src/main/routes"
)

func StartRoutingSystem(app *mux.Router) {
	routes.SetProductsRoutes(app)
}
