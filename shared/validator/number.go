package validator

import (
	"fmt"

	"github.com/akornatskyy/sample-blog-api-go/shared/errorstate"
)

type NumberValidatorBuilder interface {
	Min(min int) NumberValidatorBuilder
	Max(max int) NumberValidatorBuilder

	Build() NumberValidator
}

// NumberValidator validates a numeric value and adds any errors into
// error state.
type NumberValidator interface {
	Validate(e *errorstate.ErrorState, value int) bool
}

// Number creates number validator builder to setup validation rules.
func Number(location string) NumberValidatorBuilder {
	return &numberValidator{
		location: location,
	}
}

type numberValidator struct {
	location   string
	validators []func(*errorstate.ErrorState, int) bool
}

func (v *numberValidator) Min(min int) NumberValidatorBuilder {
	msg := fmt.Sprintf(msgMinRange, min)
	v.validators = append(v.validators, func(e *errorstate.ErrorState, value int) bool {
		if value < min {
			e.Add(&errorstate.Detail{
				Domain:   e.Domain,
				Type:     "field",
				Location: v.location,
				Reason:   "min range",
				Message:  msg,
			})
			return false
		}
		return true
	})
	return v
}

func (v *numberValidator) Max(max int) NumberValidatorBuilder {
	msg := fmt.Sprintf(msgMaxRange, max)
	v.validators = append(v.validators, func(e *errorstate.ErrorState, value int) bool {
		if value > max {
			e.Add(&errorstate.Detail{
				Domain:   e.Domain,
				Type:     "field",
				Location: v.location,
				Reason:   "max range",
				Message:  msg,
			})
			return false
		}
		return true
	})
	return v
}

func (v *numberValidator) Build() NumberValidator {
	return v
}

func (v *numberValidator) Validate(s *errorstate.ErrorState, value int) bool {
	for _, validator := range v.validators {
		if !validator(s, value) {
			return false
		}
	}
	return true
}
