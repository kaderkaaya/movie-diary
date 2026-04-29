package utils

import "golang.org/x/crypto/bcrypt"

const bcryptCost = 12

func HashPassword(plain string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plain), bcryptCost)
	return string(hashedPassword), err
}
