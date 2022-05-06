package test

import (
	watermillHttp "github.com/ThreeDotsLabs/watermill-http/pkg/http"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/kaspermroz/go-message-backend/internal/go-messages/ports/http/types"
	"net/http"
)

type SSEHandler struct {
	TestHandler http.HandlerFunc
}

func NewSSEHandler(sseRouter *watermillHttp.SSERouter) SSEHandler {
	streamAdapter := streamAdapter{}

	return SSEHandler{
		TestHandler: sseRouter.AddHandler("test", streamAdapter),
	}
}

type streamAdapter struct{}

func (a streamAdapter) GetResponse(w http.ResponseWriter, r *http.Request) (response interface{}, ok bool) {
	return &types.Test{Body: "testing!"}, true
}

func (a streamAdapter) Validate(r *http.Request, msg *message.Message) (ok bool) {
	return true
}
