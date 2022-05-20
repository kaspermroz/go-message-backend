package di

import (
	"github.com/kaspermroz/go-message-backend/internal/go-messages/app"
	"github.com/kaspermroz/go-message-backend/internal/go-messages/app/commands"
)

func BuildApplication(chatRepository commands.ChatRepository) *app.App {
	updateChat := commands.NewUpdateChatHandler(chatRepository)
	createChat := commands.NewCreateChatHandler(chatRepository)

	return &app.App{
		Commands: app.Commands{UpdateChat: updateChat, CreateChat: createChat},
		Queries:  app.Queries{},
	}
}
