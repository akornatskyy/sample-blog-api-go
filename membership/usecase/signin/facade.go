package signin

import (
	"strings"

	"github.com/akornatskyy/sample-blog-api-go/membership/domain"
)

func Process(req *Request) (*Response, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	username := strings.ToLower(req.Username)
	authInfo, err := domain.UserRepository().FindAuthInfo(username)
	if err != nil || !authInfo.IsSamePassword(req.Password) {
		return nil, errInvalidCredentials
	}
	if authInfo.IsLocked {
		return nil, errUserLocked
	}
	resp := Response{
		UserID: authInfo.UserID,
	}
	return &resp, nil
}
