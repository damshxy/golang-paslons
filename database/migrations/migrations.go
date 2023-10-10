package migrations

import (
	"go-paslons-crud/database"
	"go-paslons-crud/models"
)

func MigrationsDB() {
	database.DB.AutoMigrate(
		&models.Parties{},
		&models.Paslons{},
		&models.Votes{},
	)
}