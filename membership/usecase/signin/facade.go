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
	authInfo, err := domain.GetFactory().UserRepository().FindAuthInfo(username)
	if err != nil || !authInfo.IsSamePassword(req.Password) {
		return nil, invalidCredentials
	}
	if authInfo.IsLocked {
		return nil, userLocked
	}
	resp := Response{
		Username: username,
	}
	return &resp, nil
}
