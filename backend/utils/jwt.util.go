package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

var secret_key = "secret_key"

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	webToken, err := token.SignedString([]byte(secret_key))

	if err != nil {
		return "", err
	}

	return webToken, nil
}

func verifyToken(tokenString string) (*jwt.Token, error) {
	tokenJwt, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, isValid := token.Method.(*jwt.SigningMethodHMAC)
		if !isValid {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret_key), nil
	})

	if err != nil {
		return nil, err
	}

	return tokenJwt, err
}

func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	token, err := verifyToken(tokenString)
	if err != nil {
		return nil, err
	}

	claims, isOk := token.Claims.(jwt.MapClaims)
	if isOk && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
