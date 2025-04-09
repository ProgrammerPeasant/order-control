package repositories

import (
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
	if arg := args.Get(0); arg != nil {
		return arg.(*models.Company), args.Error(1)
	}
	return nil, args.Error(1)
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

func TestCompanyRepository(t *testing.T) {
	t.Run("CreateCompany_Success", func(t *testing.T) {
		mockRepo := new(MockCompanyRepository)
		repo := mockRepo
		company := &models.Company{Name: "Test Company"}

		mockRepo.On("CreateCompany", company).Return(nil)

		err := repo.CreateCompany(company)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetCompanyByID_Success", func(t *testing.T) {
		mockRepo := new(MockCompanyRepository)
		repo := mockRepo
		expectedCompany := &models.Company{Name: "Test Company"}

		mockRepo.On("GetCompanyByID", uint(1)).Return(expectedCompany, nil)

		company, err := repo.GetCompanyByID(1)

		assert.NoError(t, err)
		assert.Equal(t, expectedCompany, company)
		mockRepo.AssertExpectations(t)
	})

	t.Run("UpdateCompany_Success", func(t *testing.T) {
		mockRepo := new(MockCompanyRepository)
		repo := mockRepo
		company := &models.Company{Name: "Updated Company"}

		mockRepo.On("UpdateCompany", company).Return(nil)

		err := repo.UpdateCompany(company)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("DeleteCompany_Success", func(t *testing.T) {
		mockRepo := new(MockCompanyRepository)
		repo := mockRepo
		company := &models.Company{}

		mockRepo.On("DeleteCompany", company).Return(nil)

		err := repo.DeleteCompany(company)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("UpdateUserCompanyID_Success", func(t *testing.T) {
		mockRepo := new(MockCompanyRepository)
		repo := mockRepo

		mockRepo.On("UpdateUserCompanyID", uint(1), uint(2)).Return(nil)

		err := repo.UpdateUserCompanyID(1, 2)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("UpdateJoinRequestStatus_Success", func(t *testing.T) {
		mockRepo := new(MockCompanyRepository)
		repo := mockRepo

		mockRepo.On("UpdateJoinRequestStatus", uint(1), uint(2), "approved").Return(nil)

		err := repo.UpdateJoinRequestStatus(1, 2, "approved")

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}
