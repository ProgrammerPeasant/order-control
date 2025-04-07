package utils

import (
	"fmt"
	"github.com/ProgrammerPeasant/order-control/models"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found")
	}
}

var jwtSecret = []byte(os.Getenv("JWTOKEN"))

type Claims struct {
	UserID    uint   `json:"user_id"`
	Role      string `json:"role"`
	CompanyID uint   `json:"company_id"`
	Email     string `json:"email"` // Добавьте это поле
	jwt.StandardClaims
}

func GenerateJWT(user *models.User) (string, error) {
	now := time.Now()
	expireTime := now.Add(24 * time.Hour)

	claims := &Claims{
		UserID:    user.ID,
		Role:      user.Role,
		CompanyID: user.CompanyID,
		Email:     user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  now.Unix(),
			Issuer:    "myapp",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateJWT(tokenStr string) (*Claims, error) {
	log.Printf("Начало валидации токена: %s", tokenStr)

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		LoadEnv()
		log.Printf("JWT Secret: %s", string(jwtSecret))
		return jwtSecret, nil
	})

	if err != nil {
		log.Printf("Ошибка при парсинге токена: %v", err)
		return nil, err
	}

	log.Printf("Токен успешно распарсен: %+v", token)

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		log.Printf("Claims успешно извлечены и токен валиден: %+v", claims)
		userID := claims.UserID
		role := claims.Role
		companyID := claims.CompanyID
		email := claims.Email

		log.Printf("Извлеченные данные из claims: UserID=%d, Role=%s, CompanyID=%d", userID, role, companyID)

		return &Claims{
			UserID:         userID,
			Role:           role,
			CompanyID:      companyID,
			Email:          email,
			StandardClaims: jwt.StandardClaims{},
		}, nil
	}

	log.Println("Токен невалиден")
	return nil, fmt.Errorf("invalid token")
}
