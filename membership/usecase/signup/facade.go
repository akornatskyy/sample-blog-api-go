package signup

import (
	"strings"

	"github.com/akornatskyy/sample-blog-api-go/membership/domain"
	"github.com/akornatskyy/sample-blog-api-go/membership/domain/user"
)

func Process(req *Request) (*Response, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	username := strings.ToLower(strings.TrimSpace(req.Username))
	r := domain.UserRepository()
	if ok, err := r.HasAccount(username); err != nil || ok {
		return nil, errUsernameTaken
	}
	hash, err := user.PasswordHash(req.Password)
	if err != nil {
		return nil, err
	}
	reg := req.ToRegistration()
	reg.Username = username
	reg.PasswordHash = hash
	if ok, err := r.CreateAccount(reg); err != nil || !ok {
		return nil, errCreateFailed
	}
	resp := Response{}
	return &resp, nil
}
