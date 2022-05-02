package http

import "github.com/kaspermroz/go-message-backend/internal/go-messages/ports/http/test"

type RouterHandlers struct {
	Test test.QueryHandler
}

func NewHandlers() RouterHandlers {
	return RouterHandlers{
		Test: test.QueryHandler{},
	}
}
