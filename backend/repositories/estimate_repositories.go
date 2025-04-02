package repositories

import (
	"errors"
	"github.com/ProgrammerPeasant/order-control/models"
	"github.com/jinzhu/gorm"
)

// EstimateRepository определяет контракт для работы с репозиторием смет.
type EstimateRepository interface {
	Create(estimate *models.Estimate) error
	Update(estimate *models.Estimate) error
	Delete(estimate *models.Estimate) error
	GetByID(estimateID int64) (*models.Estimate, error)
	GetByCompanyID(companyID uint) ([]*models.Estimate, error)
}

// EstimateRepository представляет реализацию репозитория смет с использованием GORM.
type EstimateRepositoryImpl struct { // Переименовано в EstimateRepositoryImpl
	db *gorm.DB
}

// NewEstimateRepository создает новый экземпляр EstimateRepositoryImpl.
func NewEstimateRepository(db *gorm.DB) *EstimateRepositoryImpl { // Возвращаем *EstimateRepositoryImpl
	return &EstimateRepositoryImpl{db}
}

// Create создает новую смету в базе данных.
func (r *EstimateRepositoryImpl) Create(estimate *models.Estimate) error { // Receiver type - *EstimateRepositoryImpl
	return r.db.Create(estimate).Error
}

// Update обновляет существующую смету в базе данных.
func (r *EstimateRepositoryImpl) Update(estimate *models.Estimate) error { // Receiver type - *EstimateRepositoryImpl
	return r.db.Save(estimate).Error
}

// Delete удаляет смету из базы данных.
func (r *EstimateRepositoryImpl) Delete(estimate *models.Estimate) error { // Receiver type - *EstimateRepositoryImpl
	return r.db.Delete(estimate).Error
}

// GetByID получает смету по ее ID, включая связанные Items.
func (r *EstimateRepositoryImpl) GetByID(estimateID int64) (*models.Estimate, error) { // Receiver type - *EstimateRepositoryImpl
	var estimate models.Estimate

	if err := r.db.Preload("Items").First(&estimate, estimateID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &estimate, nil
}

// GetByCompanyID получает все сметы, принадлежащие указанной компании.
func (r *EstimateRepositoryImpl) GetByCompanyID(companyID uint) ([]*models.Estimate, error) { // Receiver type - *EstimateRepositoryImpl
	var estimates []*models.Estimate

	if err := r.db.Where("company_id = ?", companyID).Find(&estimates).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return estimates, nil
}
