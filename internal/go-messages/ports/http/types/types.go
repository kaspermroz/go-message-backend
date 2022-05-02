package types

import "net/http"

// ErrorMessage defines model for ErrorMessage.
type ErrorMessage struct {
	// Application-level error message, for debugging
	Error *string `json:"error,omitempty"`

	// User-level status message
	Status string `json:"status"`
}

type Test struct {
	Body string `json:"body"`
}

func (t Test) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
