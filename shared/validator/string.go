package validator

import (
	"fmt"
	"regexp"

	"github.com/akornatskyy/sample-blog-api-go/shared/errorstate"
)

type StringValidatorBuilder interface {
	Required() StringValidatorBuilder
	Min(min int) StringValidatorBuilder
	Max(max int) StringValidatorBuilder
	Pattern(pattern string, message string) StringValidatorBuilder

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

func (v *stringValidator) Min(min int) StringValidatorBuilder {
	msg := fmt.Sprintf(msgMinLength, min)
	v.validators = append(v.validators, func(e errorstate.State, value string) bool {
		l := len(value)
		if l != 0 && l < min {
			e.Add(&errorstate.Detail{
				Type:     "field",
				Location: v.location,
				Reason:   "min length",
				Message:  msg,
			})
			return false
		}
		return true
	})
	return v
}

func (v *stringValidator) Max(max int) StringValidatorBuilder {
	msg := fmt.Sprintf(msgMaxLength, max)
	v.validators = append(v.validators, func(e errorstate.State, value string) bool {
		if len(value) > max {
			e.Add(&errorstate.Detail{
				Type:     "field",
				Location: v.location,
				Reason:   "max length",
				Message:  msg,
			})
			return false
		}
		return true
	})
	return v
}

func (v *stringValidator) Pattern(pattern string, message string) StringValidatorBuilder {
	r := regexp.MustCompile(pattern)
	v.validators = append(v.validators, func(e errorstate.State, value string) bool {
		if value != "" && !r.MatchString(value) {
			e.Add(&errorstate.Detail{
				Type:     "field",
				Location: v.location,
				Reason:   "pattern",
				Message:  message,
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
