package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte("arash_developer1380")

func GenerateJWT(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	webToken, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return webToken, err
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("UnExpected signing method %v", token.Header["alg"])
		}

		return []byte(SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	token, err := VerifyToken(tokenString)
	if err != nil {
		return nil, err
	}

	claims, isOk := token.Claims.(jwt.MapClaims)
	if isOk && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid toekn")
}
