package util

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	byteHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(byteHash), err
}

func ComparePasswords(actualHashedPassword string, userGivenPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(actualHashedPassword), []byte(userGivenPassword))
	return err == nil
}
