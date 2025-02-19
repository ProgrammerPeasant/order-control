package services

import (
	"github.com/ProgrammerPeasant/order-control/models"
	"github.com/ProgrammerPeasant/order-control/repositories"
)

type EstimateService struct {
	estimateRepo *repositories.EstimateRepositories
}

func NewEstimateService(repo *repositories.EstimateRepositories) *EstimateService {
	return &EstimateService{estimateRepo: repo}
}

func (s *EstimateService) CreateEstimate(estimate *models.Estimate) error {
	for _, item := range estimate.Items {
		item.TotalPrice = float64(item.Quantity) * item.UnitPrice
		estimate.TotalAmount += item.TotalPrice
	}
	return s.estimateRepo.Create(estimate)
}

func (s *EstimateService) UpdateEstimate(estimate *models.Estimate) error {
	for _, item := range estimate.Items {
		item.TotalPrice = float64(item.Quantity) * item.UnitPrice
		estimate.TotalAmount += item.TotalPrice
	}
	return s.estimateRepo.Update(estimate)
}

func (s *EstimateService) DeleteEstimate(estimate *models.Estimate) error {
	return s.estimateRepo.Delete(estimate)
}

func (s *EstimateService) GetEstimateByID(estimateID int64) (*models.Estimate, error) {
	return s.estimateRepo.GetByID(estimateID)
}

func (s *EstimateService) GetEstimatesByCompanyID(companyID uint) ([]*models.Estimate, error) {
	return s.estimateRepo.GetByCompanyID(companyID)
}
