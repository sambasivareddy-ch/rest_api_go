package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("hivngssrch")

// Generates a token using passed email address
func CreateToken(userEmail string, userId int64) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email":  userEmail,
			"userId": userId,
			"exp":    time.Now().Add(time.Hour * 6).Unix(),
		},
	)

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Verifies the user passed token's validify
func VerifyToken(tokenString string) (int64, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if token.Valid && ok {
		userId := int64(claims["userId"].(float64))
		return userId, nil
	}

	return 0, errors.New("invalid authorization token")
}
