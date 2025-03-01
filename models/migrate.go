package models

import (
	"gorm.io/gorm"
)

func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&Genre{}, &Book{})
	return err
}
