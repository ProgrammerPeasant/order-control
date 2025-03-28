package services

import (
	"fmt"
	"github.com/ProgrammerPeasant/order-control/models"
	"github.com/ProgrammerPeasant/order-control/repositories"
	"github.com/xuri/excelize/v2"
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

func (s *EstimateService) ExportEstimateToExcelByID(estimateID int64) (*excelize.File, error) {
	estimate, err := s.estimateRepo.GetByID(estimateID)
	if err != nil {
		return nil, err
	}
	if estimate == nil {
		return nil, fmt.Errorf("смета с ID %d не найдена", estimateID)
	}

	f := excelize.NewFile()
	index, err := f.NewSheet("Смета")
	if err != nil {
		return nil, err
	}

	f.SetCellValue("Смета", "A1", "Название сметы:")
	f.SetCellValue("Смета", "B1", estimate.Title)

	f.SetCellValue("Смета", "A2", "Общая сумма:")
	f.SetCellValue("Смета", "B2", estimate.TotalAmount)

	f.SetCellValue("Смета", "A4", "Продукт")
	f.SetCellValue("Смета", "B4", "Количество")
	f.SetCellValue("Смета", "C4", "Цена за единицу")
	f.SetCellValue("Смета", "D4", "Итого")

	for i, item := range estimate.Items {
		row := i + 5
		f.SetCellValue("Смета", fmt.Sprintf("A%d", row), item.ProductName)
		f.SetCellValue("Смета", fmt.Sprintf("B%d", row), item.Quantity)
		f.SetCellValue("Смета", fmt.Sprintf("C%d", row), item.UnitPrice)
		f.SetCellValue("Смета", fmt.Sprintf("D%d", row), item.TotalPrice)
	}

	f.SetActiveSheet(index)
	return f, nil
}
