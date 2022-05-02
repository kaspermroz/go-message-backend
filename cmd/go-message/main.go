package main

import (
	"context"

	"github.com/kaspermroz/go-message-backend/internal/service/di"
)

func main() {
	ctx := context.Background()
	s, err := di.BuildService(ctx)
	if err != nil {
		panic(err.(interface{}))
	}

	if err := s.Run(ctx); err != nil {
		panic(err.(interface{}))
	}

}
