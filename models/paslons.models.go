package models

import "gorm.io/gorm"

type Paslons struct {
	gorm.Model
	Id uint `gorm:primaryKey`
	Name string
	Visi string
	Image string
}