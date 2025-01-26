package services

import (
	"golang.org/x/crypto/bcrypt"
	"order-control/models"
	"order-control/repositories"
	"order-control/utils"
	"strings"
	_ "strings"

	_ "golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(username, email, password string, role models.Role) error
	Login(username, password string) (string, error)
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
func (s *userService) Register(username, email, password string, role models.Role) error {
	// Хешируем пароль
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Сохраняем пользователя
	user := &models.User{
		Username: username,
		Email:    strings.ToLower(email),
		Password: string(hashed),
		Role:     role,
	}
	return s.userRepo.CreateUser(user)
}

// Логин: проверяем хеш пароля и, если всё ок, выдаём JWT
func (s *userService) Login(username, password string) (string, error) {
	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		return "", err
	}

	// Сравниваем хехированный пароль
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	// Генерируем токен
	token, err := utils.GenerateJWT(user)
	if err != nil {
		return "", err
	}
	return token, nil
}
