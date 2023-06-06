package token

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secretKey"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secretKey"), nil
	})
	if err != nil {
		return "", fmt.Errorf("failed to parse token: %w", err)
	}
	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid claims")
	}
	userID, ok := claims["userID"].(string)
	if !ok {
		return "", fmt.Errorf("invalid userID claim")
	}
	exp := claims["exp"].(float64)
	if time.Now().Unix() > int64(exp) {
		return "", fmt.Errorf("token has expired")
	}
	return userID, nil
}
