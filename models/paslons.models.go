package models

type Paslons struct {
	Id      int               `json:"id"gorm:"primary_key"`
	Name    string            `json:"name"gorm:"varchar(255) not null"`
	Visi    string            `json:"visi"gorm:"varchar(255) not null"`
	Image   string            `json:"image"gorm:"varchar(255) not null"`
	Parties	[]PartiesResponse	`json:"parties"`
	Votes   []VotesResponse   `json:"votes"`
}

type PaslonsResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Visi  string `json:"visi"`
	Image string `json:"image"`
}

func (PaslonsResponse) TableName() string {
	return "paslons"
}