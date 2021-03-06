package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/sirupsen/logrus"
)

func NewHTTPRouter(
	logger *logrus.Entry,
	handlers RouterHandlers,
) *http.Server {
	r := chi.NewRouter()
	r.Route("/api/v1", routes(handlers))
	return &http.Server{Addr: "0.0.0.0:8080", Handler: r}
}

func routes(
	handlers RouterHandlers,
) func(r chi.Router) {
	return func(r chi.Router) {
		r.Use(cors.Handler(cors.Options{
			AllowedOrigins:   []string{"http://localhost:3000", "https://go-message.vercel.app"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "User-ID"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		}))

		r.Mount("/chats/{chat_id}", getChatUpdatedRoute(handlers.ChatUpdated))
		r.Mount("/chats/{chat_id}/send", getSendMessageRoute(handlers.SendMessage))
		r.Mount("/chats", getCreateChatRoute(handlers.CreateChat, handlers.AllChatsUpdated))
	}

}

func getChatUpdatedRoute(chatUpdated ChatUpdatedSSEHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", chatUpdated.ChatUpdated)

	return r
}

func getSendMessageRoute(sendMessage SendMessageHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", sendMessage.Handle)

	return r
}

func getCreateChatRoute(createChat CreateChatHandler, allChatsUpdated AllChatsUpdatedSSEHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", createChat.Handle)
	r.Get("/", allChatsUpdated.AllChatsUpdated)
	return r
}
