package models

import "github.com/jinzhu/gorm"

// Company — пример модели компании
type Company struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"desc"`
	Address     string `json:"address"`
	// ...
}
