package models

import "github.com/jinzhu/gorm"

// Company — пример модели компании
type Company struct {
	gorm.Model
	Name           string `json:"name"`
	Description    string `json:"desc"`
	Address        string `json:"address"`
	LogoURL        string `json:"logo_url"`        // урл логотипа
	ColorPrimary   string `json:"color_primary"`   // Основной цвет (hex формат)
	ColorSecondary string `json:"color_secondary"` // Вторичный цвет (hex формат)
	ColorAccent    string `json:"color_accent"`    // Акцентный цвет (hex формат)
	// ... other fields ...
}
