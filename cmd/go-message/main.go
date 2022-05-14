package main

import (
	"context"

	"github.com/kaspermroz/go-message-backend/internal/service/di"
)

func main() {
	ctx := context.Background()
	s, appctx, err := di.BuildService(ctx)
	if err != nil {
		panic(err.(interface{}))
	}

	if err := s.Run(appctx); err != nil {
		panic(err.(interface{}))
	}

}
