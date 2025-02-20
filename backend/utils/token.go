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

var jwtSecret = []byte(os.Getenv("JWTOKEN")) // Обычно храним в env

type Claims struct {
	UserID uint        `json:"user_id"`
	Role   models.Role `json:"role"`
	jwt.StandardClaims
}

// Генерация JWT
func GenerateJWT(user *models.User) (string, error) {
	now := time.Now()
	expireTime := now.Add(24 * time.Hour)

	claims := &Claims{
		UserID: user.ID,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  now.Unix(),
			Issuer:    "myapp",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// Валидация JWT
func ValidateJWT(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
