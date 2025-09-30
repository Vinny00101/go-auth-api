package utils

import (
	"go-api/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(config.Env_Config.SECRET_KEY)

func JwtGeneration(userID int, email string) (string, error) {
	claims := jwt.MapClaims{
		"email":   email,
		"user_id": userID,
		"exp":     time.Now().Add(time.Minute * 30).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
