package models

type Votes struct {
	ID        int             `json:"id"gorm:"primary_key not null`
	VoterName string          `json:"voter_name"gorm:"varchar(255)"`
	PaslonsID int             `json:"paslons_id"`
	Paslons   PaslonsResponse `json:"paslons"gorm:"foreignKey:PaslonsID"`
}

type VotesResponse struct {
	Id        int    `json:"id"`
	VoterName string `json:"voter_name"`
	PaslonsId int    `json:"-"`
}

func (VotesResponse) TableName() string {
	return "votes"
}