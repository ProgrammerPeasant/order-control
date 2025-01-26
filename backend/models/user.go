package models

import "github.com/jinzhu/gorm"

type Role string

const (
	RoleAdmin   Role = "ADMIN"
	RoleManager Role = "MANAGER"
	RoleClient  Role = "CLIENT"
)

// User — пример модели пользователя
type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Role     Role   `gorm:"type:varchar(20);not null"`
}
