package models

import "github.com/jinzhu/gorm"

// Company — пример модели компании
type Company struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string
	Address     string
	// ...
}
