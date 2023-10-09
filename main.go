package main

import (
	"go-paslons-crud/config"
	"go-paslons-crud/config/migrations"
	"go-paslons-crud/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectToDB()
	migrations.MigrationsDB()


	routes.Routes(r.Group("/api/v1"))

	r.Run()
}