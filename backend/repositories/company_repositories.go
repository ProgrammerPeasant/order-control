package repositories

import (
	"github.com/ProgrammerPeasant/order-control/models"
	"github.com/jinzhu/gorm"
)

type CompanyRepository interface {
	CreateCompany(company *models.Company) error
	GetCompanyByID(id uint) (*models.Company, error)
	UpdateCompany(company *models.Company) error
	DeleteCompany(company *models.Company) error
	UpdateUserCompanyID(userID uint, companyID uint) error
	UpdateJoinRequestStatus(userID uint, companyID uint, status string) error
	Begin() *gorm.DB
	Commit(tx *gorm.DB) error
	Rollback(tx *gorm.DB) error
}

type companyRepository struct {
	db *gorm.DB
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

func (r *companyRepository) UpdateUserCompanyID(userID uint, companyID uint) error {
	return r.db.Model(&models.User{}).Where("id = ?", userID).Update("CompanyID", companyID).Error
}

func (r *companyRepository) UpdateJoinRequestStatus(userID uint, companyID uint, status string) error {
	return r.db.Model(&models.JoinRequest{}).Where("user_id = ? AND company_id = ?", userID, companyID).Update("Status", status).Error
}

func (r *companyRepository) Begin() *gorm.DB {
	return r.db.Begin()
}

func (r *companyRepository) Commit(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (r *companyRepository) Rollback(tx *gorm.DB) error {
	return tx.Rollback().Error
}

func NewCompanyRepository(db *gorm.DB) CompanyRepository {
	return &companyRepository{db: db}
}
