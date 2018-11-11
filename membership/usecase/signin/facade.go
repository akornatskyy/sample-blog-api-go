package signin

import (
	"strings"
)

func Process(req *Request) (*Response, error) {
	username := strings.ToLower(req.Username)
	resp := Response{
		Username: username,
	}
	return &resp, nil
}
