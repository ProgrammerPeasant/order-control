package repositories

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

func (m *MockEstimateRepository) GetByID(id int64) (*models.Estimate, error) {
	args := m.Called(id)
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

func TestEstimateRepositoryImpl(t *testing.T) {
	t.Run("Create_Success", func(t *testing.T) {
		mockRepo := new(MockEstimateRepository)
		repo := mockRepo
		estimate := &models.Estimate{Title: "Test Estimate"}

		mockRepo.On("Create", estimate).Return(nil)

		err := repo.Create(estimate)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Create_Error", func(t *testing.T) {
		mockRepo := new(MockEstimateRepository)
		repo := mockRepo
		estimate := &models.Estimate{Title: "Test Estimate"}

		mockRepo.On("Create", estimate).Return(errors.New("create error"))

		err := repo.Create(estimate)

		assert.Error(t, err)
		assert.Equal(t, "create error", err.Error())
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetByID_Success", func(t *testing.T) {
		mockRepo := new(MockEstimateRepository)
		repo := mockRepo
		expectedEstimate := &models.Estimate{Title: "Test Estimate"}

		mockRepo.On("GetByID", int64(1)).Return(expectedEstimate, nil)

		estimate, err := repo.GetByID(1)

		assert.NoError(t, err)
		assert.Equal(t, expectedEstimate, estimate)
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetByID_NotFound", func(t *testing.T) {
		mockRepo := new(MockEstimateRepository)
		repo := mockRepo

		mockRepo.On("GetByID", int64(1)).Return(nil, nil)

		estimate, err := repo.GetByID(1)

		assert.NoError(t, err)
		assert.Nil(t, estimate)
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetByCompanyID_Success", func(t *testing.T) {
		mockRepo := new(MockEstimateRepository)
		repo := mockRepo
		expectedEstimates := []*models.Estimate{
			{Title: "Estimate 1", CompanyID: 10},
			{Title: "Estimate 2", CompanyID: 10},
		}

		mockRepo.On("GetByCompanyID", uint(10)).Return(expectedEstimates, nil)

		estimates, err := repo.GetByCompanyID(10)

		assert.NoError(t, err)
		assert.Equal(t, expectedEstimates, estimates)
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetByCompanyID_Error", func(t *testing.T) {
		mockRepo := new(MockEstimateRepository)
		repo := mockRepo

		mockRepo.On("GetByCompanyID", uint(10)).Return(nil, errors.New("find error"))

		estimates, err := repo.GetByCompanyID(10)

		assert.Error(t, err)
		assert.Nil(t, estimates)
		assert.Equal(t, "find error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}
