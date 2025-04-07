package services

import (
	"fmt"

	"github.com/ProgrammerPeasant/order-control/models"
	"github.com/ProgrammerPeasant/order-control/repositories"
)

type JoinRequestService interface {
	CreateJoinRequest(userID uint, companyID uint) error
	GetPendingJoinRequests(companyID uint) ([]models.JoinRequest, error)
	ApproveJoinRequest(userID uint, companyID uint) error
	RejectJoinRequest(userID uint, companyID uint) error
}

type joinRequestService struct {
	joinRequestRepo repositories.JoinRequestRepository
	companyRepo     repositories.CompanyRepository
	userRepo        repositories.UserRepository
}

func NewJoinRequestService(joinRequestRepo repositories.JoinRequestRepository, companyRepo repositories.CompanyRepository, userRepo repositories.UserRepository) *joinRequestService {
	return &joinRequestService{joinRequestRepo: joinRequestRepo, companyRepo: companyRepo, userRepo: userRepo}
}

func (s *joinRequestService) CreateJoinRequest(userID uint, companyID uint) error {
	// Проверка, существует ли уже такой запрос
	existingRequest, err := s.joinRequestRepo.FindByUserAndCompanyID(userID, companyID)
	if err == nil && existingRequest != nil {
		return fmt.Errorf("запрос на присоединение от данного пользователя к этой компании уже существует")
	}
	if err != nil && err.Error() != "record not found" {
		return err
	}

	joinRequest := &models.JoinRequest{
		UserID:    userID,
		CompanyID: companyID,
	}
	return s.joinRequestRepo.Create(joinRequest)
}

func (s *joinRequestService) GetPendingJoinRequests(companyID uint) ([]models.JoinRequest, error) {
	return s.joinRequestRepo.GetPendingByCompanyID(companyID)
}

func (s *joinRequestService) ApproveJoinRequest(userID uint, companyID uint) error {
	request, err := s.joinRequestRepo.FindByUserAndCompanyID(userID, companyID)
	if err != nil {
		return fmt.Errorf("запрос на присоединение не найден")
	}
	if request.Status != "pending" {
		return fmt.Errorf("запрос на присоединение не находится в статусе ожидания")
	}

	tx := s.companyRepo.Begin()
	defer tx.Rollback()

	if err := s.userRepo.UpdateCompanyID(userID, companyID, tx); err != nil {
		return fmt.Errorf("ошибка при обновлении CompanyID пользователя: %w", err)
	}

	if err := s.joinRequestRepo.UpdateStatus(userID, companyID, "approved", tx); err != nil {
		return fmt.Errorf("ошибка при обновлении статуса запроса: %w", err)
	}
	return tx.Commit().Error
}

func (s *joinRequestService) RejectJoinRequest(userID uint, companyID uint) error {
	request, err := s.joinRequestRepo.FindByUserAndCompanyID(userID, companyID)
	if err != nil {
		return fmt.Errorf("запрос на присоединение не найден")
	}
	if request.Status != "pending" {
		return fmt.Errorf("запрос на присоединение не находится в статусе ожидания")
	}

	tx := s.companyRepo.Begin()
	defer tx.Rollback()

	if err := s.joinRequestRepo.UpdateStatus(userID, companyID, "rejected", tx); err != nil {
		return fmt.Errorf("ошибка при обновлении статуса запроса: %w", err)
	}

	return tx.Commit().Error
}
