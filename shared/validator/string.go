package validator

import (
	"github.com/akornatskyy/sample-blog-api-go/shared/errorstate"
)

type StringValidatorBuilder interface {
	Required() StringValidatorBuilder

	Build() StringValidator
}

// StringValidator validates a string value and adds any errors into
// error state.
type StringValidator interface {
	Validate(s errorstate.State, value string) bool
}

// String creates string validator builder to setup validation rules.
func String(location string) StringValidatorBuilder {
	return &stringValidator{
		location:   location,
		validators: []func(errorstate.State, string) bool{},
	}
}

type stringValidator struct {
	location   string
	validators []func(errorstate.State, string) bool
}

func (v *stringValidator) Required() StringValidatorBuilder {
	v.validators = append(v.validators, func(e errorstate.State, value string) bool {
		if value == "" {
			e.Add(&errorstate.Detail{
				Type:     "field",
				Location: v.location,
				Reason:   "required",
				Message:  msgRequiredField,
			})
			return false
		}
		return true
	})
	return v
}

func (v *stringValidator) Build() StringValidator {
	return v
}

func (v *stringValidator) Validate(s errorstate.State, value string) bool {
	for _, validator := range v.validators {
		if !validator(s, value) {
			return false
		}
	}
	return true
}
