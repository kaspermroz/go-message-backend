package http

// ErrorMessage defines model for ErrorMessage.
type ErrorMessage struct {
	// Application-level error message, for debugging
	Error *string `json:"error,omitempty"`

	// User-level status message
	Status string `json:"status"`
}

type Chat struct {
	UUID     string    `json:"uuid"`
	Title    string    `json:"title"`
	Messages []Message `json:"messages"`
}

type Message struct {
	AuthorId string `json:"author_id"`
	Text     string `json:"text"`
}
