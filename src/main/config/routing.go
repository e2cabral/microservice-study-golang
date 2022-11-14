package config

import (
	"github.com/gorilla/mux"
	"microservice-golang/src/infra/facades"
)

func StartRouting(router *mux.Router) {
	facades.StartRoutingSystem(router)
}
