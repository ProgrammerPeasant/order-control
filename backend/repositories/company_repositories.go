package repositories

import (
	"github.com/jinzhu/gorm"
	"order-control/models"
)

type CompanyRepository interface {
	CreateCompany(company *models.Company) error
	GetCompanyByID(id uint) (*models.Company, error)
	UpdateCompany(company *models.Company) error
	DeleteCompany(company *models.Company) error
	// ...
}

type companyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) CompanyRepository {
	return &companyRepository{db: db}
}

func (r *companyRepository) CreateCompany(company *models.Company) error {
	return r.db.Create(company).Error
}

func (r *companyRepository) GetCompanyByID(id uint) (*models.Company, error) {
	var company models.Company
	err := r.db.First(&company, id).Error
	if err != nil {
		return nil, err
	}
	return &company, nil
}

func (r *companyRepository) UpdateCompany(company *models.Company) error {
	return r.db.Save(company).Error
}

func (r *companyRepository) DeleteCompany(company *models.Company) error {
	return r.db.Delete(company).Error
}
