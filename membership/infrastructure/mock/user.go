package mock

import (
	"errors"

	"github.com/akornatskyy/sample-blog-api-go/membership/domain/user"
)

type userRepository struct {
}

func NewUserRepository() user.Repository {
	return userRepository{}
}

func (r userRepository) FindAuthInfo(username string) (*user.AuthInfo, error) {
	return nil, errors.New("not found")
}
