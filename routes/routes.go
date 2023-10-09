package routes

import "github.com/gin-gonic/gin"

func Routes(c *gin.RouterGroup) {
	PaslonsRoutes(c)
	VotesRoutes(c)
	PartiesRoutes(c)
}