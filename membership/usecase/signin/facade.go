package signin

import (
	"strings"
)

func Process(req *Request) (*Response, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	username := strings.ToLower(req.Username)
	resp := Response{
		Username: username,
	}
	return &resp, nil
}
