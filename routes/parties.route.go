package routes

import (
	"go-paslons-crud/handlers"

	"github.com/gin-gonic/gin"
)

func PartiesRoutes(c *gin.RouterGroup) {
	c.POST("/partie", handlers.CreateParties)
	c.GET("/parties", handlers.GetParties)
	c.GET("/partie/:id", handlers.GetPartiesById)
	c.PATCH("/partie/:id", handlers.UpdateParties)
	c.DELETE("/partie/:id", handlers.DeleteParties)
}