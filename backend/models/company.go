package models

import "github.com/jinzhu/gorm"

// Company — пример модели компании
type Company struct {
	gorm.Model
	Name         string   `json:"name"`
	Description  string   `json:"desc"`
	Address      string   `json:"address"`
	LogoURL      string   `json:"logo_url"`                         // урл логотипа
	DesignColors []string `json:"design_colors" gorm:"type:text[]"` // три цвета в hex формате
	// ... other fields ...
}
