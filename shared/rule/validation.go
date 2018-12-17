package rule

import (
	"github.com/akornatskyy/sample-blog-api-go/shared/validator"
)

var (
	Email    = validator.String("email").Required().Min(6).Max(50).Email().Build()
	Username = validator.String("username").Required().Min(2).Max(20).Build()
	Password = validator.String("password").Required().Min(8).Max(12).Build()
)
