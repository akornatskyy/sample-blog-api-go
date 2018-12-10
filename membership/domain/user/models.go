package user

import (
	"golang.org/x/crypto/bcrypt"
)

type (
	AuthInfo struct {
		UserID       int
		IsLocked     bool
		PasswordHash []byte
	}
)

func (a *AuthInfo) IsSamePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword(a.PasswordHash, []byte(password))
	return err == nil
}
