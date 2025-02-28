package models

import (
	"gorm.io/gorm"
)

type Genre struct {
	ID  	uint    	`gorm:"primaryKey;autoIncrement" json:"id"`
	Name 	*string 	`json:"name"`
}

type Book struct{
	ID			uint		`gorm:"primary key;autoIncrement" json:"id"`
	Author		*string		`json:"author"`
	Title		*string		`json:"title"`
	Publisher 	*string		`json:"publisher"`
	GenreID 	uint 		`json:"genre_id"`
	Genre		Genre		`gorm:"constraint:onUpdate:CASCADE,onDelete:SET NULL;"`
}

func MigrateBooks(db *gorm.DB) error{
	err := db.AutoMigrate(&Genre{}, &Book{})
	return err
}