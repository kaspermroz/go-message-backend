package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kaspermroz/go-message-backend/internal/go-messages/ports/http/types"
	"net/http"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	watermillHttp "github.com/ThreeDotsLabs/watermill-http/pkg/http"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/pkg/errors"
)

type Service struct {
	sseRouter  *watermillHttp.SSERouter
	httpRouter *http.Server
	publisher  *gochannel.GoChannel
}

func NewService(
	sseRouter *watermillHttp.SSERouter,
	router *http.Server,
	publisher *gochannel.GoChannel,
) (*Service, error) {
	return &Service{
		sseRouter:  sseRouter,
		httpRouter: router,
		publisher:  publisher,
	}, nil
}

func (s Service) Run(ctx context.Context) error {
	var serviceErrors = make(chan error)

	go func() {
		serviceErrors <- s.sseRouter.Run(ctx)
	}()

	go func() {
		serviceErrors <- s.httpRouter.ListenAndServe()
	}()

	// Test SSE
	messages, err := s.publisher.Subscribe(ctx, "test")
	if err != nil {
		return err
	}
	for i := 0; i < 5; i++ {
		go s.publishTestMessage()
	}

	for _ = range messages {
		time.Sleep(2 * time.Second)
		go s.publishTestMessage()
	}

	if err := <-serviceErrors; err != nil {
		return errors.Wrap(err, "could not run service")
	}

	return nil
}

func (s Service) publishTestMessage() {
	test := types.Test{Body: "testing!"}
	payload, err := json.Marshal(test)
	if err != nil {
		panic(err.(interface{}))
	}
	msg := message.NewMessage(watermill.NewUUID(), payload)
	err = s.publisher.Publish("test", msg)
	if err != nil {
		panic(err.(interface{}))
	}
	fmt.Println("message published")
}
