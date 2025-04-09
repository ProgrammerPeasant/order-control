package services

import (
	"errors"
	"github.com/jinzhu/gorm"
	"testing"

	"github.com/ProgrammerPeasant/order-control/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCompanyRepository struct {
	mock.Mock
}

func (m *MockCompanyRepository) CreateCompany(company *models.Company) error {
	args := m.Called(company)
	return args.Error(0)
}

func (m *MockCompanyRepository) GetCompanyByID(id uint) (*models.Company, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Company), args.Error(1)
}

func (m *MockCompanyRepository) UpdateCompany(company *models.Company) error {
	args := m.Called(company)
	return args.Error(0)
}

func (m *MockCompanyRepository) DeleteCompany(company *models.Company) error {
	args := m.Called(company)
	return args.Error(0)
}

func (m *MockCompanyRepository) UpdateUserCompanyID(userID uint, companyID uint) error {
	args := m.Called(userID, companyID)
	return args.Error(0)
}

func (m *MockCompanyRepository) UpdateJoinRequestStatus(userID uint, companyID uint, status string) error {
	args := m.Called(userID, companyID, status)
	return args.Error(0)
}

func (m *MockCompanyRepository) Begin() *gorm.DB {
	args := m.Called()
	return args.Get(0).(*gorm.DB)
}

func (m *MockCompanyRepository) Commit(tx *gorm.DB) error {
	args := m.Called(tx)
	return args.Error(0)
}

func (m *MockCompanyRepository) Rollback(tx *gorm.DB) error {
	args := m.Called(tx)
	return args.Error(0)
}

func TestCompanyService_Create(t *testing.T) {
	// Setup
	mockRepo := new(MockCompanyRepository)
	companyService := NewCompanyService(mockRepo)
	company := &models.Company{Name: "Test Company", Description: "Test Desc", Address: "Test Address"}

	mockRepo.On("CreateCompany", company).Return(nil) // Ожидаем успешное создание

	// Act
	createdCompany, err := companyService.Create("Test Company", "Test Desc", "Test Address")

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, createdCompany)
	assert.Equal(t, "Test Company", createdCompany.Name)
	assert.Equal(t, "Test Desc", createdCompany.Description)
	assert.Equal(t, "Test Address", createdCompany.Address)
	mockRepo.AssertExpectations(t) // Проверяем, что mock-репозиторий был вызван с ожидаемыми аргументами
}

func TestCompanyService_Create_Error(t *testing.T) {
	// Setup
	mockRepo := new(MockCompanyRepository)
	companyService := NewCompanyService(mockRepo)
	company := &models.Company{Name: "Test Company", Description: "Test Desc", Address: "Test Address"}
	expectedError := errors.New("failed to create company")

	mockRepo.On("CreateCompany", company).Return(expectedError)

	// Act
	createdCompany, err := companyService.Create("Test Company", "Test Desc", "Test Address")

	// Assert
	assert.Error(t, err)
	assert.Nil(t, createdCompany)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}
