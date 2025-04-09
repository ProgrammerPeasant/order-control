package services

import (
	"fmt"
	"time"

	"github.com/ProgrammerPeasant/order-control/models"
	"github.com/ProgrammerPeasant/order-control/repositories"
	"github.com/ProgrammerPeasant/order-control/utils"
)

type JoinRequestService interface {
	CreateJoinRequest(userID uint, companyID uint, email string) error
	GetPendingJoinRequests(companyID uint) ([]models.JoinRequest, error)
	ApproveJoinRequest(userID uint, companyID uint) error
	RejectJoinRequest(userID uint, companyID uint) error
}

type joinRequestService struct {
	joinRequestRepo repositories.JoinRequestRepository
	companyRepo     repositories.CompanyRepository
	userRepo        repositories.UserRepository
	metrics         *utils.Metrics
}

func NewJoinRequestService(
	joinRequestRepo repositories.JoinRequestRepository,
	companyRepo repositories.CompanyRepository,
	userRepo repositories.UserRepository,
	metrics *utils.Metrics,
) *joinRequestService {
	return &joinRequestService{
		joinRequestRepo: joinRequestRepo,
		companyRepo:     companyRepo,
		userRepo:        userRepo,
		metrics:         metrics,
	}
}

func (s *joinRequestService) CreateJoinRequest(userID uint, companyID uint, email string) error {
	startTime := time.Now()
	defer func() {
		duration := time.Since(startTime).Seconds()
		s.metrics.JoinRequestDuration.WithLabelValues("create").Observe(duration)
	}()

	s.metrics.JoinRequestOperationsTotal.WithLabelValues("create", "attempt").Inc()

	existingRequest, err := s.joinRequestRepo.FindByUserAndCompanyID(userID, companyID)
	if err == nil && existingRequest != nil {
		s.metrics.JoinRequestOperationsTotal.WithLabelValues("create", "duplicate").Inc()
		s.metrics.RegisterError("join_request_error", "duplicate_request")
		return fmt.Errorf("запрос на присоединение от данного пользователя к этой компании уже существует")
	}
	if err != nil && err.Error() != "record not found" {
		s.metrics.JoinRequestOperationsTotal.WithLabelValues("create", "error").Inc()
		s.metrics.RegisterError("join_request_error", "find_error")
		return err
	}

	joinRequest := &models.JoinRequest{
		UserID:    userID,
		CompanyID: companyID,
		Email:     email,
	}

	err = s.joinRequestRepo.Create(joinRequest)
	if err != nil {
		s.metrics.JoinRequestOperationsTotal.WithLabelValues("create", "error").Inc()
		s.metrics.RegisterError("join_request_error", "create_error")
		return err
	}

	s.metrics.JoinRequestOperationsTotal.WithLabelValues("create", "success").Inc()
	s.metrics.JoinRequestTotal.WithLabelValues("pending").Inc()

	return nil
}

func (s *joinRequestService) GetPendingJoinRequests(companyID uint) ([]models.JoinRequest, error) {
	startTime := time.Now()
	defer func() {
		duration := time.Since(startTime).Seconds()
		s.metrics.JoinRequestDuration.WithLabelValues("get_pending").Observe(duration)
	}()

	s.metrics.JoinRequestOperationsTotal.WithLabelValues("get_pending", "attempt").Inc()

	requests, err := s.joinRequestRepo.GetPendingByCompanyID(companyID)
	if err != nil {
		s.metrics.JoinRequestOperationsTotal.WithLabelValues("get_pending", "error").Inc()
		s.metrics.RegisterError("join_request_error", "get_pending_error")
		return nil, err
	}

	s.metrics.JoinRequestOperationsTotal.WithLabelValues("get_pending", "success").Inc()

	s.metrics.PendingJoinRequestsGauge.WithLabelValues(fmt.Sprintf("%d", companyID)).Set(float64(len(requests)))

	return requests, nil
}

func (s *joinRequestService) ApproveJoinRequest(userID uint, companyID uint) error {
	startTime := time.Now()
	defer func() {
		duration := time.Since(startTime).Seconds()
		s.metrics.JoinRequestDuration.WithLabelValues("approve").Observe(duration)
	}()

	s.metrics.JoinRequestOperationsTotal.WithLabelValues("approve", "attempt").Inc()

	request, err := s.joinRequestRepo.FindByUserAndCompanyID(userID, companyID)
	if err != nil {
		s.metrics.JoinRequestOperationsTotal.WithLabelValues("approve", "error").Inc()
		s.metrics.RegisterError("join_request_error", "request_not_found")
		return fmt.Errorf("запрос на присоединение не найден")
	}
	if request.Status != "pending" {
		s.metrics.JoinRequestOperationsTotal.WithLabelValues("approve", "invalid_status").Inc()
		s.metrics.RegisterError("join_request_error", "invalid_status")
		return fmt.Errorf("запрос на присоединение не находится в статусе ожидания")
	}

	tx := s.companyRepo.Begin()
	defer tx.Rollback()

	if err := s.userRepo.UpdateCompanyID(userID, companyID, tx); err != nil {
		s.metrics.JoinRequestOperationsTotal.WithLabelValues("approve", "update_user_error").Inc()
		s.metrics.RegisterError("join_request_error", "update_user_error")
		return fmt.Errorf("ошибка при обновлении CompanyID пользователя: %w", err)
	}

	if err := s.joinRequestRepo.UpdateStatus(userID, companyID, "approved", tx); err != nil {
		s.metrics.JoinRequestOperationsTotal.WithLabelValues("approve", "update_status_error").Inc()
		s.metrics.RegisterError("join_request_error", "update_status_error")
		return fmt.Errorf("ошибка при обновлении статуса запроса: %w", err)
	}

	err = tx.Commit().Error
	if err != nil {
		s.metrics.JoinRequestOperationsTotal.WithLabelValues("approve", "commit_error").Inc()
		s.metrics.RegisterError("join_request_error", "commit_error")
		return err
	}

	// Обновляем метрики
	s.metrics.JoinRequestOperationsTotal.WithLabelValues("approve", "success").Inc()
	s.metrics.JoinRequestTotal.WithLabelValues("approved").Inc()
	s.metrics.JoinRequestTotal.WithLabelValues("pending").Dec()

	return nil
}

func (s *joinRequestService) RejectJoinRequest(userID uint, companyID uint) error {
	startTime := time.Now()
	defer func() {
		// Измеряем время выполнения операции
		duration := time.Since(startTime).Seconds()
		s.metrics.JoinRequestDuration.WithLabelValues("reject").Observe(duration)
	}()

	s.metrics.JoinRequestOperationsTotal.WithLabelValues("reject", "attempt").Inc()

	request, err := s.joinRequestRepo.FindByUserAndCompanyID(userID, companyID)
	if err != nil {
		s.metrics.JoinRequestOperationsTotal.WithLabelValues("reject", "error").Inc()
		s.metrics.RegisterError("join_request_error", "request_not_found")
		return fmt.Errorf("запрос на присоединение не найден")
	}
	if request.Status != "pending" {
		s.metrics.JoinRequestOperationsTotal.WithLabelValues("reject", "invalid_status").Inc()
		s.metrics.RegisterError("join_request_error", "invalid_status")
		return fmt.Errorf("запрос на присоединение не находится в статусе ожидания")
	}

	tx := s.companyRepo.Begin()
	defer tx.Rollback()

	if err := s.joinRequestRepo.UpdateStatus(userID, companyID, "rejected", tx); err != nil {
		s.metrics.JoinRequestOperationsTotal.WithLabelValues("reject", "update_status_error").Inc()
		s.metrics.RegisterError("join_request_error", "update_status_error")
		return fmt.Errorf("ошибка при обновлении статуса запроса: %w", err)
	}

	err = tx.Commit().Error
	if err != nil {
		s.metrics.JoinRequestOperationsTotal.WithLabelValues("reject", "commit_error").Inc()
		s.metrics.RegisterError("join_request_error", "commit_error")
		return err
	}

	s.metrics.JoinRequestOperationsTotal.WithLabelValues("reject", "success").Inc()
	s.metrics.JoinRequestTotal.WithLabelValues("rejected").Inc()
	s.metrics.JoinRequestTotal.WithLabelValues("pending").Dec()

	return nil
}
