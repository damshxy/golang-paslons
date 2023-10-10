package main

import (
	"go-paslons-crud/database"
	"go-paslons-crud/database/migrations"
	"go-paslons-crud/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.ConnectToDB()
	migrations.MigrationsDB()


	routes.Routes(r.Group("/api/v1"))

	r.Run()
}