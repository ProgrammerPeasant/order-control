package repositories

import (
	"testing"

	"github.com/ProgrammerPeasant/order-control/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(user *models.User) (*models.User, error) {
	args := m.Called(user)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) GetUserByUsername(username string) (*models.User, error) {
	args := m.Called(username)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) FindByEmail(email string) (*models.User, error) {
	args := m.Called(email)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) FindByID(id uint) (*models.User, error) {
	args := m.Called(id)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) UpdateCompanyID(userID uint, companyID uint, tx interface{}) error {
	args := m.Called(userID, companyID, tx)
	return args.Error(0)
}

func TestUserRepository(t *testing.T) {
	t.Run("CreateUser_Success", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		repo := mockRepo
		user := &models.User{Username: "testuser", Email: "test@example.com"}

		mockRepo.On("CreateUser", user).Return(user, nil)

		createdUser, err := repo.CreateUser(user)

		assert.NoError(t, err)
		assert.Equal(t, user, createdUser)
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetUserByUsername_Success", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		repo := mockRepo
		expectedUser := &models.User{Username: "testuser"}

		mockRepo.On("GetUserByUsername", "testuser").Return(expectedUser, nil)

		user, err := repo.GetUserByUsername("testuser")

		assert.NoError(t, err)
		assert.Equal(t, expectedUser, user)
		mockRepo.AssertExpectations(t)
	})

	t.Run("FindByEmail_Success", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		repo := mockRepo
		expectedUser := &models.User{Email: "test@example.com"}

		mockRepo.On("FindByEmail", "test@example.com").Return(expectedUser, nil)

		user, err := repo.FindByEmail("test@example.com")

		assert.NoError(t, err)
		assert.Equal(t, expectedUser, user)
		mockRepo.AssertExpectations(t)
	})

	t.Run("FindByID_Success", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		repo := mockRepo
		expectedUser := &models.User{}

		mockRepo.On("FindByID", uint(1)).Return(expectedUser, nil)

		user, err := repo.FindByID(1)

		assert.NoError(t, err)
		assert.Equal(t, expectedUser, user)
		mockRepo.AssertExpectations(t)
	})

	t.Run("UpdateCompanyID_Success", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		repo := mockRepo

		mockRepo.On("UpdateCompanyID", uint(1), uint(2), mock.Anything).Return(nil)

		err := repo.UpdateCompanyID(1, 2, nil)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}
