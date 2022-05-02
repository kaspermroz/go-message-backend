package test

import (
	"github.com/kaspermroz/go-message-backend/internal/go-messages/ports/http/render"
	"github.com/kaspermroz/go-message-backend/internal/go-messages/ports/http/types"
	"net/http"
)

type QueryHandler struct{}

func (h QueryHandler) GetTest(w http.ResponseWriter, r *http.Request) {
	testReponse := &types.Test{Body: "testing!"}
	render.MustRender(w, r, testReponse)
}
