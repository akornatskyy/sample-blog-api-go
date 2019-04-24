package sql

import (
	"database/sql"

	"github.com/akornatskyy/sample-blog-api-go/membership/domain/user"
)

type userRepository struct {
}

func NewUserRepository(db *sql.DB) user.Repository {
	r := &userRepository{}
	return r
}

func (r *userRepository) FindAuthInfo(username string) (*user.AuthInfo, error) {
	return nil, nil
}

func (r *userRepository) FindUserByID(id string) (*user.User, error) {
	return nil, nil
}

func (r *userRepository) HasAccount(username string) (bool, error) {
	return true, nil
}

func (r *userRepository) CreateAccount(reg *user.Registration) (bool, error) {
	return false, nil
}
