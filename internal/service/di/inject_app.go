package di

import (
	"github.com/kaspermroz/go-message-backend/internal/go-messages/app"
	"github.com/kaspermroz/go-message-backend/internal/go-messages/app/commands"
	"github.com/kaspermroz/go-message-backend/internal/go-messages/app/queries"
)

func BuildApplication(chatRepository commands.ChatRepository, userRepository commands.UserRepository) *app.App {
	updateChat := commands.NewUpdateChatHandler(chatRepository)
	createChat := commands.NewCreateChatHandler(chatRepository, userRepository)

	getUserChats := queries.NewGetUserChatsHandler(chatRepository)

	return &app.App{
		Commands: app.Commands{UpdateChat: updateChat, CreateChat: createChat},
		Queries:  app.Queries{GetUserChats: getUserChats},
	}
}
