package services

import (
	"fmt"
	"github.com/ProgrammerPeasant/order-control/models"
	"github.com/ProgrammerPeasant/order-control/repositories"
	"github.com/xuri/excelize/v2"
	"strconv"
)

type EstimateService struct {
	estimateRepo repositories.EstimateRepository // Используем интерфейс
}

func NewEstimateService(repo repositories.EstimateRepository) *EstimateService { // Принимаем интерфейс
	return &EstimateService{estimateRepo: repo}
}

func (s *EstimateService) CreateEstimate(estimate *models.Estimate) error {
	estimate.TotalAmount = 0

	for i := range estimate.Items {
		item := &estimate.Items[i]
		item.TotalPrice = float64(item.Quantity) * item.UnitPrice
		if item.DiscountPercent > 0 {
			discountAmount := item.TotalPrice * (item.DiscountPercent / 100)
			item.TotalPrice -= discountAmount
		}
		estimate.TotalAmount += item.TotalPrice
	}

	if estimate.OverallDiscountPercent > 0 {
		discountAmount := estimate.TotalAmount * (estimate.OverallDiscountPercent / 100)
		estimate.TotalAmount -= discountAmount
	}
	return s.estimateRepo.Create(estimate)
}

func (s *EstimateService) UpdateEstimate(estimate *models.Estimate) error {
	estimate.TotalAmount = 0

	for i := range estimate.Items {
		item := &estimate.Items[i]
		item.TotalPrice = float64(item.Quantity) * item.UnitPrice
		if item.DiscountPercent > 0 {
			discountAmount := item.TotalPrice * (item.DiscountPercent / 100)
			item.TotalPrice -= discountAmount
		}
		estimate.TotalAmount += item.TotalPrice
	}

	if estimate.OverallDiscountPercent > 0 {
		discountAmount := estimate.TotalAmount * (estimate.OverallDiscountPercent / 100)
		estimate.TotalAmount -= discountAmount
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

	err = f.SetCellValue("Смета", "A1", "Название сметы:")
	if err != nil {
		return nil, err
	}
	err = f.SetCellValue("Смета", "B1", estimate.Title)
	if err != nil {
		return nil, err
	}

	err = f.SetCellValue("Смета", "A2", "Общая сумма:")
	if err != nil {
		return nil, err
	}
	err = f.SetCellValue("Смета", "B2", estimate.TotalAmount)
	if err != nil {
		return nil, err
	}

	err = f.SetCellValue("Смета", "A4", "Продукт")
	if err != nil {
		return nil, err
	}
	err = f.SetCellValue("Смета", "B4", "Количество")
	if err != nil {
		return nil, err
	}
	err = f.SetCellValue("Смета", "C4", "Цена за единицу")
	if err != nil {
		return nil, err
	}
	err = f.SetCellValue("Смета", "D4", "Скидка (%)")
	if err != nil {
		return nil, err
	}
	err = f.SetCellValue("Смета", "E4", "Итого")
	if err != nil {
		return nil, err
	}

	for i, item := range estimate.Items {
		row := i + 5
		err := f.SetCellValue("Смета", fmt.Sprintf("A%d", row), item.ProductName)
		if err != nil {
			return nil, err
		}
		err = f.SetCellValue("Смета", fmt.Sprintf("B%d", row), item.Quantity)
		if err != nil {
			return nil, err
		}
		err = f.SetCellValue("Смета", fmt.Sprintf("C%d", row), item.UnitPrice)
		if err != nil {
			return nil, err
		}
		err = f.SetCellValue("Смета", fmt.Sprintf("D%d", row), item.DiscountPercent)
		if err != nil {
			return nil, err
		}
		err = f.SetCellValue("Смета", fmt.Sprintf("E%d", row), item.TotalPrice)
		if err != nil {
			return nil, err
		}
	}

	err = f.SetCellValue("Смета", "A"+strconv.Itoa(len(estimate.Items)+6), "Общая скидка (%):")
	if err != nil {
		return nil, err
	}
	err = f.SetCellValue("Смета", "B"+strconv.Itoa(len(estimate.Items)+6), estimate.OverallDiscountPercent)
	if err != nil {
		return nil, err
	}

	f.SetActiveSheet(index)
	return f, nil
}
