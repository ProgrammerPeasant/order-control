package models

import "github.com/jinzhu/gorm"

const (
	RoleAdmin   = "ADMIN"
	RoleManager = "MANAGER"
	RoleClient  = "CLIENT"
)

// User — пример модели пользователя
type User struct {
	gorm.Model
	Username  string  `gorm:"unique;not null"`
	Password  string  `gorm:"not null"`
	Email     string  `gorm:"unique;not null"`
	Role      string  `gorm:"type:varchar(20);not null"` // Тип Role теперь string
	CompanyID uint    `gorm:"index"`                     // Добавляем CompanyID как внешний ключ
	Company   Company `gorm:"foreignkey:CompanyID"`      // Связь с моделью Company
}
