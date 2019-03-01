package user

import (
	"golang.org/x/crypto/bcrypt"
)

type (
	AuthInfo struct {
		UserID       string
		IsLocked     bool
		PasswordHash []byte
	}

	User struct {
		Username  string
		FirstName string
		LastName  string
	}

	Registration struct {
		Email        string
		Username     string
		PasswordHash []byte
	}
)

func (a *AuthInfo) IsSamePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword(a.PasswordHash, []byte(password))
	return err == nil
}

func PasswordHash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
