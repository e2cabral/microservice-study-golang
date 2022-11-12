package facades

import (
	"github.com/gin-gonic/gin"
	"microservice-golang/src/main/routes"
)

func StartRoutingSystem(app *gin.Engine) {
	routes.SetProductsRoutes(app)
}
