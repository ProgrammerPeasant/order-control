package models

import "github.com/jinzhu/gorm"

type EstimateItem struct {
	gorm.Model
	EstimateID  uint
	ProductName string
	Quantity    uint
	UnitPrice   uint
	TotalPrice  uint
}

type Estimate struct {
	gorm.Model
	Title       string
	TotalAmount float64
	CompanyID   uint
	Company     Company `gorm:"foreignkey:CompanyID"`
	CreatedByID uint
	CreatedBy   User `gorm:"foreignkey:CreatedByID"`
	Items       []EstimateItem
}
