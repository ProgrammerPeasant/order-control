package utils

import (
	"fmt"
	"github.com/ProgrammerPeasant/order-control/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("4c89213f58c5fa8b9ad4827f869a8b926e8b4c73ca75c1391cd150d375519cd33bc72dbe622a47310dc57b49faf6537214e0266b305b05cc704b77a6a7533d445a2b9218247973a8ebbec220fcc1b3bc1417377b33d33800efe40f0fee2bc1f714474bfd620bfb1dcfc1ac25b2badef3cf1a360c11fc103b006f58323eb102b3bdce7c8b3a9b4658fd3506f921df803aec3547ac38fce044a24b968202371a76ed6fa168e6ecde7917d8cf4e80876fba50667622bb5a6e86d6ea4b15316cc820bb076870aed751acc64f2faec605bda1920d183364f35b4826f5e800647f8f9b5fdace980806bff895bc5cc6c350eea9c2a587649c7eb840d45dac635dec4cc3") // Обычно храним в env

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
