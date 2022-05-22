package http

import (
	"context"
	"github.com/ThreeDotsLabs/watermill"
	watermillHttp "github.com/ThreeDotsLabs/watermill-http/pkg/http"
	"github.com/ThreeDotsLabs/watermill/message"
	appctx "github.com/kaspermroz/go-message-backend/internal/go-messages/app"
	"github.com/kaspermroz/go-message-backend/internal/go-messages/domain"
	"github.com/kaspermroz/go-message-backend/internal/go-messages/ports/pubsub"
	"net/http"
)

type AllChatsUpdatedSSEHandler struct {
	AllChatsUpdated http.HandlerFunc
}

func NewAllChatsUpdatedSSEHandler(sseRouter *watermillHttp.SSERouter, logger watermill.LoggerAdapter, ctx context.Context) AllChatsUpdatedSSEHandler {
	streamAdapter := allChatsUpdatedStreamAdapter{
		ctxWithApp: ctx,
		logger:     logger,
	}

	return AllChatsUpdatedSSEHandler{AllChatsUpdated: sseRouter.AddHandler(pubsub.TopicChatUpdated, streamAdapter)}
}

type allChatsUpdatedStreamAdapter struct {
	ctxWithApp context.Context
	logger     watermill.LoggerAdapter
}

func (a allChatsUpdatedStreamAdapter) GetResponse(w http.ResponseWriter, r *http.Request) (response interface{}, ok bool) {
	app, err := appctx.GetApplication(a.ctxWithApp)
	if err != nil {
		a.logger.Error("could not get app from context", err, nil)
		return nil, false
	}

	req := AllChatsUpdatedRequest{UserID: r.Header.Get("User-ID")}

	if err != nil {
		a.logger.Error("could not parse request", err, nil)
		return nil, false
	}

	userID, err := domain.NewUUID(req.UserID)
	if err != nil {
		a.logger.Error("could not get user ID from request", err, nil)
		return nil, false
	}

	chats, err := app.Queries.GetUserChats.Handle(r.Context(), userID)
	if err != nil {
		a.logger.Error("could not get user chats for user", err, watermill.LogFields{
			"user_id": userID.String(),
		})
		return nil, false
	}

	projections := mapChatsToAllChatsProjection(chats)
	res := AllChatsResponse{Chats: projections}

	return res, true
}

func (a allChatsUpdatedStreamAdapter) Validate(r *http.Request, msg *message.Message) (ok bool) {
	return true
}
