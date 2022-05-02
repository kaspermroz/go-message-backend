package http

import (
	"github.com/go-chi/chi"
	"github.com/kaspermroz/go-message-backend/internal/go-messages/ports/http/test"
	"github.com/sirupsen/logrus"
)

func NewHTTPRouter(
	logger *logrus.Entry,
	handlers RouterHandlers,
) chi.Router {
	r := chi.NewRouter()
	r.Route("/api/v1", routes(handlers))
	return r
}

func routes(
	handlers RouterHandlers,
) func(r chi.Router) {
	return func(r chi.Router) {
		r.Mount("/test", getTestRoute(handlers.Test))
	}

}

func getTestRoute(test test.QueryHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", test.GetTest)

	return r
}
