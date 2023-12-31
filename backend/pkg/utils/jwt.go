package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/luckyss6/cc_server/pkg/constant"
)

func CreateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"name":  "John Doe",
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	s, err := token.SignedString([]byte(constant.JwtSecret))
	if err != nil {
		return "", err
	}
	return s, nil
}
