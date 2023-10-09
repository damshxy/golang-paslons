package migrations

import (
	"go-paslons-crud/config"
	"go-paslons-crud/models"
)

func MigrationsDB() {
	config.DB.AutoMigrate(
		&models.Paslons{},
		&models.Parties{},
		&models.Votes{},
	)
}