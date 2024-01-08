package database

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func (u *User) HashPassword(password string) (string, error) {
	if len(password) < 12 {
		return "", errors.New("Password must be at least 12 characters long")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}
