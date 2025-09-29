package utils

import (
	"go-api/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(config.Env_Config.SECRET_KEY)

func JwtGeneration(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Minute * 30).Unix(),
	}

	token := jwt.NewWithClaims(&jwt.SigningMethodRSA{}, claims)
	return token.SignedString(jwtSecret)
}
