package services

import (
	"errors"
	"github.com/ProgrammerPeasant/order-control/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockEstimateRepository struct {
	mock.Mock
}

func (m *MockEstimateRepository) Create(estimate *models.Estimate) error {
	args := m.Called(estimate)
	return args.Error(0)
}

func (m *MockEstimateRepository) GetByID(estimateID int64) (*models.Estimate, error) {
	args := m.Called(estimateID)
	if arg := args.Get(0); arg != nil {
		return arg.(*models.Estimate), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockEstimateRepository) GetByCompanyID(companyID uint) ([]*models.Estimate, error) {
	args := m.Called(companyID)
	if arg := args.Get(0); arg != nil {
		return arg.([]*models.Estimate), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockEstimateRepository) Update(estimate *models.Estimate) error {
	args := m.Called(estimate)
	return args.Error(0)
}

func (m *MockEstimateRepository) Delete(estimate *models.Estimate) error {
	args := m.Called(estimate)
	return args.Error(0)
}

func TestEstimateService_CreateEstimate(t *testing.T) {
	// Setup
	mockRepo := new(MockEstimateRepository)
	estimateService := NewEstimateService(mockRepo)
	estimate := &models.Estimate{
		Title: "Test Estimate",
		Items: []models.EstimateItem{
			{ProductName: "Product 1", Quantity: 2, UnitPrice: 10.0},
			{ProductName: "Product 2", Quantity: 1, UnitPrice: 25.0, DiscountPercent: 10.0},
		},
		OverallDiscountPercent: 5.0,
	}

	mockRepo.On("Create", mock.Anything).Return(nil) // Ожидаем успешное создание

	// Act
	err := estimateService.CreateEstimate(estimate)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 40.375, estimate.TotalAmount) // Проверяем правильность расчета TotalAmount
	assert.Equal(t, 20.0, estimate.Items[0].TotalPrice)
	assert.Equal(t, 22.5, estimate.Items[1].TotalPrice)
	mockRepo.AssertExpectations(t)
}

func TestEstimateService_CreateEstimate_Error(t *testing.T) {
	// Setup
	mockRepo := new(MockEstimateRepository)
	estimateService := NewEstimateService(mockRepo)
	estimate := &models.Estimate{Title: "Test Estimate"}
	expectedError := errors.New("failed to create estimate")

	mockRepo.On("Create", estimate).Return(expectedError)

	// Act
	err := estimateService.CreateEstimate(estimate)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestEstimateService_UpdateEstimate(t *testing.T) {
	mockRepo := new(MockEstimateRepository)
	estimateService := NewEstimateService(mockRepo)
	estimate := &models.Estimate{
		Title: "Updated Estimate",
		Items: []models.EstimateItem{
			{ProductName: "Product 3", Quantity: 3, UnitPrice: 15.0, DiscountPercent: 5.0},
		},
		OverallDiscountPercent: 10.0,
	}

	mockRepo.On("Update", mock.Anything).Return(nil)

	err := estimateService.UpdateEstimate(estimate)

	assert.NoError(t, err)
	assert.Equal(t, 38.475, estimate.TotalAmount)
	assert.Equal(t, 42.75, estimate.Items[0].TotalPrice)
	mockRepo.AssertExpectations(t)
}

func TestEstimateService_UpdateEstimate_RepoError(t *testing.T) {
	// Setup
	mockRepo := new(MockEstimateRepository)
	estimateService := NewEstimateService(mockRepo)
	estimate := &models.Estimate{Title: "Updated Estimate"}
	expectedError := errors.New("failed to update estimate")

	// Определяем ожидаемое поведение mock-репозитория (ошибка)
	mockRepo.On("Update", mock.Anything).Return(expectedError)

	// Act
	err := estimateService.UpdateEstimate(estimate)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestEstimateService_DeleteEstimate(t *testing.T) {
	// Setup
	mockRepo := new(MockEstimateRepository)
	estimateService := NewEstimateService(mockRepo)
	estimate := &models.Estimate{}

	// Определяем ожидаемое поведение mock-репозитория
	mockRepo.On("Delete", estimate).Return(nil)

	// Act
	err := estimateService.DeleteEstimate(estimate)

	// Assert
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestEstimateService_DeleteEstimate_RepoError(t *testing.T) {
	// Setup
	mockRepo := new(MockEstimateRepository)
	estimateService := NewEstimateService(mockRepo)
	estimate := &models.Estimate{}
	expectedError := errors.New("failed to delete estimate")

	// Определяем ожидаемое поведение mock-репозитория (ошибка)
	mockRepo.On("Delete", estimate).Return(expectedError)

	// Act
	err := estimateService.DeleteEstimate(estimate)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestEstimateService_GetEstimateByID_Found(t *testing.T) {
	// Setup
	mockRepo := new(MockEstimateRepository)
	estimateService := NewEstimateService(mockRepo)
	expectedEstimate := &models.Estimate{Title: "Found Estimate"}

	// Определяем ожидаемое поведение mock-репозитория
	mockRepo.On("GetByID", int64(1)).Return(expectedEstimate, nil)

	// Act
	estimate, err := estimateService.GetEstimateByID(1)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedEstimate, estimate)
	mockRepo.AssertExpectations(t)
}

func TestEstimateService_GetEstimateByID_NotFound(t *testing.T) {
	// Setup
	mockRepo := new(MockEstimateRepository)
	estimateService := NewEstimateService(mockRepo)

	// Определяем ожидаемое поведение mock-репозитория
	mockRepo.On("GetByID", int64(1)).Return(nil, nil)

	// Act
	estimate, err := estimateService.GetEstimateByID(1)

	// Assert
	assert.NoError(t, err)
	assert.Nil(t, estimate)
	mockRepo.AssertExpectations(t)
}

func TestEstimateService_GetEstimateByID_RepoError(t *testing.T) {
	// Setup
	mockRepo := new(MockEstimateRepository)
	estimateService := NewEstimateService(mockRepo)
	expectedError := errors.New("failed to get estimate")

	// Определяем ожидаемое поведение mock-репозитория (ошибка)
	mockRepo.On("GetByID", int64(1)).Return(nil, expectedError)

	// Act
	estimate, err := estimateService.GetEstimateByID(1)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, estimate)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestEstimateService_GetEstimatesByCompanyID_Found(t *testing.T) {
	// Setup
	mockRepo := new(MockEstimateRepository)
	estimateService := NewEstimateService(mockRepo)
	expectedEstimates := []*models.Estimate{
		{Title: "Estimate 1", CompanyID: 10},
		{Title: "Estimate 2", CompanyID: 10},
	}

	// Определяем ожидаемое поведение mock-репозитория
	mockRepo.On("GetByCompanyID", uint(10)).Return(expectedEstimates, nil)

	// Act
	estimates, err := estimateService.GetEstimatesByCompanyID(10)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedEstimates, estimates)
	mockRepo.AssertExpectations(t)
}

func TestEstimateService_GetEstimatesByCompanyID_NotFound(t *testing.T) {
	// Setup
	mockRepo := new(MockEstimateRepository)
	estimateService := NewEstimateService(mockRepo)
	expectedEstimates := []*models.Estimate{}

	// Определяем ожидаемое поведение mock-репозитория
	mockRepo.On("GetByCompanyID", uint(10)).Return(expectedEstimates, nil)

	// Act
	estimates, err := estimateService.GetEstimatesByCompanyID(10)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedEstimates, estimates)
	mockRepo.AssertExpectations(t)
}

func TestEstimateService_GetEstimatesByCompanyID_RepoError(t *testing.T) {
	// Setup
	mockRepo := new(MockEstimateRepository)
	estimateService := NewEstimateService(mockRepo)
	expectedError := errors.New("failed to get estimates")

	// Определяем ожидаемое поведение mock-репозитория (ошибка)
	mockRepo.On("GetByCompanyID", uint(10)).Return(nil, expectedError)

	// Act
	estimates, err := estimateService.GetEstimatesByCompanyID(10)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, estimates)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}
