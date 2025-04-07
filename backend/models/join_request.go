package models

import "github.com/jinzhu/gorm"

type JoinRequest struct {
	gorm.Model
	UserID    uint    `json:"user_id"`
	CompanyID uint    `json:"company_id"`
	Status    string  `json:"status" gorm:"type:varchar(20);default:'pending'"` // pending, approved, rejected
	User      User    `gorm:"foreignkey:UserID"`
	Company   Company `gorm:"foreignkey:CompanyID"`
}
