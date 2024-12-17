package utils

import (
	"time"

	"github.com/Alfazal007/apiserver/internal/database"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(user database.User) (string, error) {
	jwtSecret := LoadEnvFiles().AccessTokenSecret

	secretKey := []byte(jwtSecret)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()
	claims["authorized"] = true
	claims["user_id"] = user.ID
	claims["username"] = user.Username

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
