package utils

import "golang.org/x/crypto/bcrypt"

func HashingPassword(password string) (string, error) {
	hashedByte, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(hashedByte), err
}

func CheckPassword(password string, hasPassword string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hasPassword),
		[]byte(password),
	)
	return err == nil
}
