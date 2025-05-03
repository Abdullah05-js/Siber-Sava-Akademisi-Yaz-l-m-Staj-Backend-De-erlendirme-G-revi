package utils

import (
	"Todo-list/models"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

func GenerateToken(user *models.User) (string, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET"))
	method := jwt.SigningMethodHS256
	claims := jwt.MapClaims{
		"userID": user.ID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // 24 saat
	}

	token, err := jwt.NewWithClaims(method, claims).SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return token, nil
}

