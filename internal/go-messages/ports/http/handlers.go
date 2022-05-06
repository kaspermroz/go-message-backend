package http

import (
	watermillHttp "github.com/ThreeDotsLabs/watermill-http/pkg/http"
	"github.com/kaspermroz/go-message-backend/internal/go-messages/ports/http/test"
)

type RouterHandlers struct {
	Test test.SSEHandler
}

func NewHandlers(sseRouter *watermillHttp.SSERouter) RouterHandlers {
	return RouterHandlers{
		Test: test.NewSSEHandler(sseRouter),
	}
}
