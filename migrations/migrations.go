package main

import (
	"go-paslons-crud/config"
	"go-paslons-crud/models"
)

func init() {
	config.ConnectToDB()
}

func main() {
	config.DB.AutoMigrate(&models.Paslons{})
}