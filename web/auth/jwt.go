package auth

import (
	"GoToKube/config"
	"GoToKube/logger"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID uint
	jwt.RegisteredClaims
}

func GenerateJWT(userID uint) (string, error) {
	jwtKey := []byte(config.Data.Auth.JwtSecret)
	if len(jwtKey) == 0 {
		err := errors.New("JWT secret key is not set in environment variables")
		logger.GlobalLogger.Error(err.Error())
		return "", err
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// 使用 HS512 签名方法创建 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString(jwtKey)
}

func ParseJWT(tokenString string) (*Claims, error) {
	jwtKey := []byte(config.Data.Auth.JwtSecret)
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
