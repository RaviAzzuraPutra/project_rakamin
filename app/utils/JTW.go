package utils

import (
	"last-project/app/config/app_config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(app_config.JWT)

func GenerateJWT(userID string, isAdmin bool) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  userID,
		"is_admin": isAdmin,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	return token.SignedString(jwtSecret)
}
