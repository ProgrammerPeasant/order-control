package repositories

import (
	"testing"

	"github.com/ProgrammerPeasant/order-control/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockJoinRequestRepository struct {
	mock.Mock
}

func (m *MockJoinRequestRepository) Create(joinRequest *models.JoinRequest) error {
	args := m.Called(joinRequest)
	return args.Error(0)
}

func (m *MockJoinRequestRepository) GetPendingByCompanyID(companyID uint) ([]models.JoinRequest, error) {
	args := m.Called(companyID)
	return args.Get(0).([]models.JoinRequest), args.Error(1)
}

func (m *MockJoinRequestRepository) FindByUserAndCompanyID(userID uint, companyID uint) (*models.JoinRequest, error) {
	args := m.Called(userID, companyID)
	return args.Get(0).(*models.JoinRequest), args.Error(1)
}

func (m *MockJoinRequestRepository) UpdateStatus(userID uint, companyID uint, status string, tx interface{}) error {
	args := m.Called(userID, companyID, status, tx)
	return args.Error(0)
}

func TestJoinRequestRepository(t *testing.T) {
	t.Run("Create_Success", func(t *testing.T) {
		mockRepo := new(MockJoinRequestRepository)
		repo := mockRepo
		joinRequest := &models.JoinRequest{UserID: 1, CompanyID: 2, Status: "pending"}

		mockRepo.On("Create", joinRequest).Return(nil)

		err := repo.Create(joinRequest)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetPendingByCompanyID_Success", func(t *testing.T) {
		mockRepo := new(MockJoinRequestRepository)
		repo := mockRepo
		expectedRequests := []models.JoinRequest{
			{UserID: 1, CompanyID: 2, Status: "pending"},
		}

		mockRepo.On("GetPendingByCompanyID", uint(2)).Return(expectedRequests, nil)

		joinRequests, err := repo.GetPendingByCompanyID(2)

		assert.NoError(t, err)
		assert.Equal(t, expectedRequests, joinRequests)
		mockRepo.AssertExpectations(t)
	})

	t.Run("FindByUserAndCompanyID_Success", func(t *testing.T) {
		mockRepo := new(MockJoinRequestRepository)
		repo := mockRepo
		expectedRequest := &models.JoinRequest{UserID: 1, CompanyID: 2, Status: "pending"}

		mockRepo.On("FindByUserAndCompanyID", uint(1), uint(2)).Return(expectedRequest, nil)

		joinRequest, err := repo.FindByUserAndCompanyID(1, 2)

		assert.NoError(t, err)
		assert.Equal(t, expectedRequest, joinRequest)
		mockRepo.AssertExpectations(t)
	})

	t.Run("UpdateStatus_Success", func(t *testing.T) {
		mockRepo := new(MockJoinRequestRepository)
		repo := mockRepo

		mockRepo.On("UpdateStatus", uint(1), uint(2), "approved", mock.Anything).Return(nil)

		err := repo.UpdateStatus(1, 2, "approved", nil)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}
