package models

import "github.com/jinzhu/gorm"

type EstimateItem struct {
	gorm.Model
	EstimateID  uint    `json:"estimate_id"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
	TotalPrice  float64 `json:"total_price"`
}

type Estimate struct {
	gorm.Model
	Title       string         `json:"title"`
	TotalAmount float64        `json:"total_amount"`
	CompanyID   uint           `json:"company_id"`
	Company     Company        `gorm:"foreignkey:CompanyID"`
	CreatedByID uint           `json:"created_by_id"`
	CreatedBy   User           `gorm:"foreignkey:CreatedByID"`
	Items       []EstimateItem `json:"items"`
}
