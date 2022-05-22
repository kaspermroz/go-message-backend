package di

import (
	"github.com/kaspermroz/go-message-backend/internal/go-messages/app"
	"github.com/kaspermroz/go-message-backend/internal/go-messages/app/commands"
)

func BuildApplication(chatRepository commands.ChatRepository, userRepository commands.UserRepository) *app.App {
	updateChat := commands.NewUpdateChatHandler(chatRepository)
	createChat := commands.NewCreateChatHandler(chatRepository, userRepository)

	return &app.App{
		Commands: app.Commands{UpdateChat: updateChat, CreateChat: createChat},
		Queries:  app.Queries{},
	}
}
