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

type SendMessageRequest struct {
	Message Message `json:"message"`
}

type CreateChatRequest struct {
	Name    string   `json:"name"`
	UserIDs []string `json:"user_ids"` // array of User uuids
}

type AllChatsUpdatedRequest struct {
	UserID string `json:"user_id"`
}

type User struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
}

type AllChatsProjectionChat struct {
	ChatID        string `json:"chat_id"`
	Title         string `json:"title"`
	MessagesCount int    `json:"messages_count"`
	Users         []User `json:"users"`
}

type AllChatsResponse struct {
	Chats []AllChatsProjectionChat `json:"chats"`
}
