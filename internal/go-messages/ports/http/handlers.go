package http

import (
	"github.com/ThreeDotsLabs/watermill"

	watermillHttp "github.com/ThreeDotsLabs/watermill-http/pkg/http"
)

type RouterHandlers struct {
	ChatUpdated ChatUpdatedSSEHandler
}

func NewHandlers(sseRouter *watermillHttp.SSERouter, repository ChatRepository, logger watermill.LoggerAdapter) RouterHandlers {
	return RouterHandlers{
		ChatUpdated: NewChatUpdatedSSEHandler(sseRouter, repository, logger),
	}
}
