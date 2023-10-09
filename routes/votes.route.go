package routes

import (
	"go-paslons-crud/handlers"

	"github.com/gin-gonic/gin"
)

func VotesRoutes(c *gin.RouterGroup)  {
	c.POST("/vote", handlers.CreateVotes)
	c.GET("/votes", handlers.GetVotes)
	c.GET("/vote/:id", handlers.GetVotesById)
	c.PATCH("/vote/:id", handlers.UpadatePaslon)
	c.DELETE("/vote/:id", handlers.DeleteVotes)
}