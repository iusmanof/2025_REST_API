package models

type Book struct {
	ID        uint    `gorm:"primary key;autoIncrement" json:"id"`
	Author    *string `json:"author"`
	Title     *string `json:"title"`
	Publisher *string `json:"publisher"`
	GenreID   uint    `json:"genre_id"`
	Genre     Genre   `gorm:"constraint:onUpdate:CASCADE,onDelete:SET NULL;"`
}
