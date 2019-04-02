package mock

import (
	"errors"

	"github.com/akornatskyy/sample-blog-api-go/membership/domain/user"
	"github.com/akornatskyy/sample-blog-api-go/shared/mock"
	"github.com/google/uuid"
)

var (
	errNotFound = errors.New("not found")
)

type userRepository struct {
}

func NewUserRepository() user.Repository {
	return &userRepository{}
}

func (*userRepository) FindAuthInfo(username string) (*user.AuthInfo, error) {
	for _, u := range mock.DB.Users {
		if u.Username == username {
			m := &user.AuthInfo{
				UserID:       u.ID,
				IsLocked:     u.IsLocked,
				PasswordHash: []byte(u.PasswordHash),
			}
			return m, nil
		}
	}
	return nil, errNotFound
}

func (*userRepository) FindUserByID(id string) (*user.User, error) {
	u := mock.DB.UserByID[id]
	if u == nil {
		return nil, errNotFound
	}
	m := &user.User{
		Username:  u.Username,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}
	return m, nil
}

func (*userRepository) HasAccount(username string) (bool, error) {
	for _, u := range mock.DB.Users {
		if u.Username == username {
			return true, nil
		}
	}
	return false, nil
}

func (*userRepository) CreateAccount(reg *user.Registration) (bool, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return false, err
	}
	u := &mock.User{
		ID:           id.String(),
		Username:     reg.Username,
		PasswordHash: string(reg.PasswordHash),
		IsLocked:     false,
	}
	mock.DB.Users = append(mock.DB.Users, u)
	mock.DB.UserByID[u.ID] = u
	return true, nil
}
