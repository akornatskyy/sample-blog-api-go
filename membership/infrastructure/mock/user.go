package mock

import (
	"errors"
	"log"

	"github.com/akornatskyy/sample-blog-api-go/membership/domain/user"
	"github.com/akornatskyy/sample-blog-api-go/shared/iojson"
	"github.com/google/uuid"
)

type userRepository struct {
}

type userInfo struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
	IsLocked     bool   `json:"is_locked"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}

var (
	users = loadUsers("user-samples.json")

	errNotFound = errors.New("not found")
)

func NewUserRepository() user.Repository {
	return userRepository{}
}

func (r userRepository) FindAuthInfo(username string) (*user.AuthInfo, error) {
	u, ok := users[username]
	if !ok {
		return nil, errNotFound
	}
	m := user.AuthInfo{
		UserID:       u.ID,
		IsLocked:     u.IsLocked,
		PasswordHash: []byte(u.PasswordHash),
	}
	return &m, nil
}

func (r userRepository) FindUserByID(id string) (*user.User, error) {
	for _, u := range users {
		if u.ID == id {
			m := user.User{
				Username:  u.Username,
				FirstName: u.FirstName,
				LastName:  u.LastName,
			}
			return &m, nil
		}
	}

	return nil, errNotFound
}

func (r userRepository) HasAccount(username string) (bool, error) {
	_, ok := users[username]
	return ok, nil
}

func (r userRepository) CreateAccount(reg *user.Registration) (bool, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return false, err
	}
	users[reg.Username] = userInfo{
		ID:           id.String(),
		Username:     reg.Username,
		PasswordHash: string(reg.PasswordHash),
		IsLocked:     false,
	}
	return true, nil
}

func loadUsers(filename string) map[string]userInfo {
	var r map[string][]userInfo
	if err := iojson.ReadFile(filename, &r); err != nil {
		panic(err)
	}

	users := map[string]userInfo{}
	for _, u := range r["users"] {
		users[u.Username] = u
	}

	log.Printf("loaded %d users", len(users))

	return users
}
