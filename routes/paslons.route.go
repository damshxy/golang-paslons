package routes

import (
	"go-paslons-crud/handlers"

	"github.com/gin-gonic/gin"
)

func PaslonsRoutes(c *gin.RouterGroup) {
	c.POST("/paslon", handlers.CreatePaslons)
	c.GET("/paslons", handlers.GetPaslons)
	c.GET("/paslon/:id", handlers.GetPaslonById)
	c.PATCH("/paslon/:id", handlers.UpadatePaslon)
	c.DELETE("/paslon/:id", handlers.DeletePaslon)
}