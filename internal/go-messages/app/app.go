package app

import (
	"github.com/kaspermroz/go-message-backend/internal/go-messages/app/commands"
	"github.com/kaspermroz/go-message-backend/internal/go-messages/app/queries"
)

type App struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	UpdateChat commands.UpdateChatHandler
	CreateChat commands.CreateChatHandler
}

type Queries struct {
	GetUserChats queries.GetUserChatsHandler
}
