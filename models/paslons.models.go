package models

import "gorm.io/gorm"

type Paslons struct {
	gorm.Model
	ID int `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"type:varchar(255)"`
	Visi string `json:"visi" gorm:"type:varchar(255)"`
	Image string `json:"image" gorm:"type:varchar(255)"`
}