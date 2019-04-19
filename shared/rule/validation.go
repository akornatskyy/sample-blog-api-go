// Package rule contains validation rules that can be reused.
package rule

import (
	"github.com/akornatskyy/goext/validator"
)

var (
	Email    = validator.String("email").Required().Min(6).Max(50).Email().Build()
	Username = validator.String("username").Required().Min(2).Max(20).Build()
	Password = validator.String("password").Required().Min(8).Max(12).Build()

	Query = validator.String("q").Max(20).Build()
	Page  = validator.Number("page").Min(0).Max(9).Build()

	Slug   = validator.String("slug").Required().Min(2).Max(35).Build()
	Fields = validator.String("fields").Max(20).Pattern(
		"^(comments|permissions)(,(comments|permissions))?$",
		"Required to match valid options.",
	).Build()

	Message = validator.String("message").Required().Min(2).Max(250).Build()
)
