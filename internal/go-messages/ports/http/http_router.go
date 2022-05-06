package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/sirupsen/logrus"

	"github.com/kaspermroz/go-message-backend/internal/go-messages/ports/http/test"
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
		r.Mount("/test", getTestRoute(handlers.Test))
	}

}

func getTestRoute(test test.SSEHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Get("/", test.TestHandler)

	return r
}
