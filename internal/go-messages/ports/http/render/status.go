package render

import (
	"net/http"

	"github.com/go-chi/render"
)

type StatusResponse struct {
	HTTPStatusCode int `json:"-"` // http response status code
}

func (e *StatusResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

var StatusNoContent = &StatusResponse{HTTPStatusCode: http.StatusNoContent}
