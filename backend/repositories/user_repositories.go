package repositories

import (
	"github.com/ProgrammerPeasant/order-control/models"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindByID(id uint) (*models.User, error)
	UpdateCompanyID(userID uint, companyID uint, tx *gorm.DB) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(user *models.User) (*models.User, error) {
	result := r.db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *userRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateCompanyID(userID uint, companyID uint, tx *gorm.DB) error {
	return tx.Model(&models.User{}).Where("id = ?", userID).Update("company_id", companyID).Error
}
