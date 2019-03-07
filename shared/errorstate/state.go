// Package errorstate contains structs to keep multiple error details.
package errorstate

// ErrorState contains information about domain raising this error and
// any accociated error details.
type ErrorState struct {
	// Domain is unique identifier for the use case raising this error.
	Domain string `json:"-"`
	// Error details accociated with this error.
	Errors []*Detail `json:"errors,omitempty"`
}

// Detail contains information about error specifics.
type Detail struct {
	// Domain is unique identifier for the use case raising this error.
	// This helps distinguish usecase-specific errors.
	Domain string `json:"domain,omitempty"`
	// Indicates how the location property should be interpreted.
	Type string `json:"type,omitempty"`
	// The location of the error (the interpretation of its value depends on type).
	Location string `json:"location,omitempty"`
	// Unique identifier for this error.
	Reason string `json:"reason,omitempty"`
	// A human readable message providing more details about the error.
	Message string `json:"message,omitempty"`
}

// Single constructs ErrorState with a single Detail.
func Single(d *Detail) *ErrorState {
	return &ErrorState{
		Errors: []*Detail{d},
	}
}

// Add detail related to this error.
func (e *ErrorState) Add(d *Detail) *ErrorState {
	e.Errors = append(e.Errors, d)
	return e
}

// OrNil returns nil if there is no errors added to this state, otherwise
// returns error.
func (e *ErrorState) OrNil() error {
	if e.Errors != nil {
		return e
	}
	return nil
}

// Error returns a string representation.
func (e *ErrorState) Error() string {
	return "errorstate: multiple errors"
}
