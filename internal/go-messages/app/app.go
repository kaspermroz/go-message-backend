package app

import (
	"github.com/kaspermroz/go-message-backend/internal/go-messages/app/commands"
)

type App struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	SendMessage commands.SendMessageHandler
}

type Queries struct{}
