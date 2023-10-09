package models

type Parties struct {
	Id        int             `json:"id"gorm:"primary_key"`
	Name      string          `json:"name"gorm:"varchar(255)"`
	PaslonsID int             `json:"paslons_id"`
	Paslons   PaslonsResponse `json:"paslons"gorm:"foreignKey:PaslonsID"`
}

type PartiesResponse struct {
	Id        int             `json:"id"`
	Name      string          `json:"name"`
	PaslonsID int             `json:"-"`
}

func (PartiesResponse) TableName() string {
	return "parties"
}