package main

import (
	"go-paslons-crud/config"
	"go-paslons-crud/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectToDB()

	routes.Routes(r.Group("/api/v1"))

	r.Run()
}