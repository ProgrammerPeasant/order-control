package repositories

import (
	"github.com/ProgrammerPeasant/order-control/models"
	"github.com/jinzhu/gorm"
)

type JoinRequestRepository interface {
	Create(joinRequest *models.JoinRequest) error
	GetPendingByCompanyID(companyID uint) ([]models.JoinRequest, error)
	FindByUserAndCompanyID(userID uint, companyID uint) (*models.JoinRequest, error)
	UpdateStatus(userID uint, companyID uint, status string, tx *gorm.DB) error
}

type joinRequestRepository struct {
	db *gorm.DB
}

func NewJoinRequestRepository(db *gorm.DB) *joinRequestRepository {
	return &joinRequestRepository{db: db}
}

func (r *joinRequestRepository) Create(joinRequest *models.JoinRequest) error {
	return r.db.Create(joinRequest).Error
}

func (r *joinRequestRepository) GetPendingByCompanyID(companyID uint) ([]models.JoinRequest, error) {
	var joinRequests []models.JoinRequest
	err := r.db.Where("company_id = ? AND status = ?", companyID, "pending").Find(&joinRequests).Error
	return joinRequests, err
}

func (r *joinRequestRepository) FindByUserAndCompanyID(userID uint, companyID uint) (*models.JoinRequest, error) {
	var joinRequest models.JoinRequest
	err := r.db.Where("user_id = ? AND company_id = ?", userID, companyID).First(&joinRequest).Error
	if err != nil {
		return nil, err
	}
	return &joinRequest, nil
}

func (r *joinRequestRepository) UpdateStatus(userID uint, companyID uint, status string, tx *gorm.DB) error {
	return tx.Model(&models.JoinRequest{}).Where("user_id = ? AND company_id = ?", userID, companyID).Update("Status", status).Error
}
