package mock

import (
	"errors"
	"log"

	"github.com/akornatskyy/sample-blog-api-go/membership/domain/user"
	"github.com/akornatskyy/sample-blog-api-go/shared/iojson"
)

type userRepository struct {
}

type userInfo struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
	IsLocked     bool   `json:"is_locked"`
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
		ID:           u.ID,
		IsLocked:     u.IsLocked,
		PasswordHash: []byte(u.PasswordHash),
	}
	return &m, nil
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
