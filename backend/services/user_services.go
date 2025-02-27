package services

import (
	"github.com/ProgrammerPeasant/order-control/models"
	"github.com/ProgrammerPeasant/order-control/repositories"
	"github.com/ProgrammerPeasant/order-control/utils"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type UserService interface {
	Register(username, email, password string, role string, companyID uint) error
	Login(username, password string) (*models.User, string, error)
	// ...
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(ur repositories.UserRepository) UserService {
	return &userService{
		userRepo: ur,
	}
}

// Регистрация нового пользователя
func (s *userService) Register(username, email, password string, role string, companyID uint) error { // Role is string, CompanyID is uint
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.User{
		Username:  username,
		Email:     strings.ToLower(email),
		Password:  string(hashed),
		Role:      role,
		CompanyID: companyID,
	}
	return s.userRepo.CreateUser(user)
}

// Логин: проверяем хеш пароля и, если всё ок, выдаём JWT и модель пользователя
func (s *userService) Login(username, password string) (*models.User, string, error) {
	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		return nil, "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, "", err
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		return nil, "", err
	}
	return user, token, nil
}
