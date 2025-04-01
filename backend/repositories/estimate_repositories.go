package repositories

import (
	"errors"
	"github.com/ProgrammerPeasant/order-control/models"
	"github.com/jinzhu/gorm"
)

type EstimateRepositories struct {
	db *gorm.DB
}

func NewEstimateRepositories(db *gorm.DB) *EstimateRepositories {
	return &EstimateRepositories{db}
}

func (r *EstimateRepositories) Create(estimate *models.Estimate) error {
	return r.db.Create(estimate).Error
}

func (r *EstimateRepositories) Update(estimate *models.Estimate) error {
	return r.db.Save(estimate).Error
}

func (r *EstimateRepositories) Delete(estimate *models.Estimate) error {
	return r.db.Delete(estimate).Error
}

func (r *EstimateRepositories) GetByID(estimateID int64) (*models.Estimate, error) {
	var estimate models.Estimate

	if err := r.db.Preload("Items").First(&estimate, estimateID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &estimate, nil
}

func (r *EstimateRepositories) GetByCompanyID(companyID uint) ([]*models.Estimate, error) { // Изменено название функции и тип параметра
	var estimates []*models.Estimate

	if err := r.db.Where("company_id = ?", companyID).Find(&estimates).Error; err != nil { // Используем company_id
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return estimates, nil
}
