package utils

import (
	"hih-yadebi-backend/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var cfgJWT = config.LoadJWT()

func GenerateJWT(userID int, username string) (string, error) {
	jwtSecret := cfgJWT.JWTSecret
	jwtTimeout := time.Minute * time.Duration(cfgJWT.JWTTimeout)
	// jwtTimeout, _ := cfg.JWTTimeout * time.Minute
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"timeout":  time.Now().Add(jwtTimeout).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}
