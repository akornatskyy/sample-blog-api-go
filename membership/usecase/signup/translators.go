package signup

import (
	"github.com/akornatskyy/sample-blog-api-go/membership/domain/user"
)

func (req *Request) ToRegistration() *user.Registration {
	return &user.Registration{
		Email: req.Email,
	}
}
