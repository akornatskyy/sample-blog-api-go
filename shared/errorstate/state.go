package errorstate

import "fmt"

// State interface to keep track of errors.
type State interface {
	error

	Add(d *Detail) State

	OrNil() error
}

// Detail contains information about error specifics.
type Detail struct {
	Domain   string `json:"domain,omitempty"`
	Type     string `json:"type,omitempty"`
	Location string `json:"location,omitempty"`
	Reason   string `json:"reason,omitempty"` // Unique identifier for this error.
	Message  string `json:"message,omitempty"`
}

// New returns error state for given domain. Domain is unique identifier for
// the use case raising this error. This helps distinguish usecase-specific
// errors.
func New(domain string) State {
	return &errorState{
		Domain: domain,
		Errors: []*Detail{},
	}
}

type errorState struct {
	Domain string    `json:"-"`
	Errors []*Detail `json:"errors,omitempty"`
}

// Add detail related to this error.
func (e *errorState) Add(d *Detail) State {
	if d.Domain == "" {
		d.Domain = e.Domain
	}
	e.Errors = append(e.Errors, d)
	return e
}

// OrNil returns nil if there is no errors added to this state, otherwise
// returns error.
func (e *errorState) OrNil() error {
	if len(e.Errors) != 0 {
		return e
	}
	return nil
}

func (e *errorState) Error() string {
	return fmt.Sprintf("%#v", e)
}
