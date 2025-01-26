package services

import (
	"order-control/models"
	"order-control/repositories"
)

type CompanyService interface {
	Create(name, desc, address string) (*models.Company, error)
	GetByID(id uint) (*models.Company, error)
	Update(company *models.Company) error
	Delete(id uint) error
}

type companyService struct {
	companyRepo repositories.CompanyRepository
}

func NewCompanyService(cr repositories.CompanyRepository) CompanyService {
	return &companyService{companyRepo: cr}
}

func (s *companyService) Create(name, desc, address string) (*models.Company, error) {
	c := &models.Company{
		Name:        name,
		Description: desc,
		Address:     address,
	}
	return c, s.companyRepo.CreateCompany(c)
}

func (s *companyService) GetByID(id uint) (*models.Company, error) {
	return s.companyRepo.GetCompanyByID(id)
}

func (s *companyService) Update(company *models.Company) error {
	return s.companyRepo.UpdateCompany(company)
}

func (s *companyService) Delete(id uint) error {
	company, err := s.companyRepo.GetCompanyByID(id)
	if err != nil {
		return err
	}
	return s.companyRepo.DeleteCompany(company)
}
