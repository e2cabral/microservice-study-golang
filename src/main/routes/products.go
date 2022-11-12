package routes

import (
	"github.com/gin-gonic/gin"
	"microservice-golang/src/infra/helpers"
	"net/http"
)

func SetProductsRoutes(app *gin.Engine) {
	router := helpers.NewRouter("/products")

	products := app.Group(router.Group)
	{
		products.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello, World!",
				"test":    "Testing my terminal",
			})
		})
	}
}
