package app

import (
	"github.com/kaspermroz/go-message-backend/internal/go-messages/app/commands"
)

type App struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	UpdateChat commands.UpdateChatHandler
	CreateChat commands.CreateChatHandler
}

type Queries struct{}
