package services

import (
	"errors"
	"github.com/ProgrammerPeasant/order-control/models"
	"github.com/ProgrammerPeasant/order-control/repositories"
	"github.com/ProgrammerPeasant/order-control/utils"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type UserService interface {
	Register(username, email, password string, role string, companyID uint) (*models.User, error)
	Login(username, password string) (*models.User, string, error)
	CreateJoinRequest(userID uint, companyID uint, email string) error
	// ...
}

type userService struct {
	userRepo           repositories.UserRepository
	joinRequestService JoinRequestService // Необходимо указать тип
	metrics            *utils.Metrics
}

func NewUserService(userRepo repositories.UserRepository, joinRequestService JoinRequestService, metrics *utils.Metrics) UserService {
	return &userService{
		userRepo:           userRepo,
		joinRequestService: joinRequestService,
		metrics:            metrics,
	}
}

// Регистрация нового пользователя
//func (s *userService) Register(username, email, password string, role string, companyID uint) error {
//	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
//	if err != nil {
//		return err
//	}
//
//	user := &models.User{
//		Username:  username,
//		Email:     strings.ToLower(email),
//		Password:  string(hashed),
//		Role:      role,
//		CompanyID: companyID,
//	}
//	return s.userRepo.CreateUser(user)
//}

func (s *userService) Register(username, email, password string, role string, companyID uint) (*models.User, error) {
	existingUser, err := s.userRepo.GetUserByUsername(username)
	if err == nil && existingUser != nil {
		s.metrics.RegisterError("registration_error", "Username already exists")
		return nil, errors.New("пользователь с таким именем уже существует")
	}

	existingEmail, err := s.userRepo.FindByEmail(email)
	if err == nil && existingEmail != nil {
		s.metrics.RegisterError("registration_error", "Email already exists")
		return nil, errors.New("пользователь с таким email уже существует")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err

	}

	user := &models.User{
		Username:  username,
		Email:     strings.ToLower(email),
		Password:  string(hashed),
		Role:      role,
		CompanyID: companyID,
	}

	createdUser, err := s.userRepo.CreateUser(user)
	if err != nil {
		s.metrics.RegisterDBError("create", "users")
		return nil, err // Возвращаем nil для пользователя и ошибку
	}

	return createdUser, nil // Возвращаем созданного пользователя и nil для ошибки
}

// Логин: проверяем хеш пароля и, если всё ок, выдаём JWT и модель пользователя
func (s *userService) Login(username, password string) (*models.User, string, error) {
	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		s.metrics.RegisterError("authentication_error", "User not found")
		return nil, "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		s.metrics.RegisterError("authentication_error", "Invalid password")
		return nil, "", err
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		s.metrics.RegisterError("token_generation_error", err.Error())
		return nil, "", err
	}
	return user, token, nil
}

func (s *userService) CreateJoinRequest(userID uint, companyID uint, email string) error {
	return s.joinRequestService.CreateJoinRequest(userID, companyID, email)
}
