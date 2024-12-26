package utils

import "github.com/golang-jwt/jwt/v5"

var SecretKey = []byte("arash_developer1380")

func GenerateJWT(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	webToken, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return webToken, err
}
