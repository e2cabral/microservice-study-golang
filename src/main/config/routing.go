package config

import (
	"github.com/gin-gonic/gin"
	"microservice-golang/src/infra/facades"
)

func StartRouting(router *gin.Engine) {
	facades.StartRoutingSystem(router)
}
