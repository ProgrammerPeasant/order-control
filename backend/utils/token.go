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
	jwt.StandardClaims
}

func GenerateJWT(user *models.User) (string, error) {
	now := time.Now()
	expireTime := now.Add(24 * time.Hour)

	claims := &Claims{
		UserID:    user.ID,
		Role:      user.Role,
		CompanyID: user.CompanyID,
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
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		LoadEnv()
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		userID := claims.UserID
		role := claims.Role
		companyID := claims.CompanyID

		return &Claims{
			UserID:         userID,
			Role:           role,
			CompanyID:      companyID,
			StandardClaims: jwt.StandardClaims{},
		}, nil
	}

	return nil, fmt.Errorf("invalid token")
}
