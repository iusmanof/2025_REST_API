package models

type Genre struct {
	ID   uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name *string `json:"name"`
}
