package render

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/kaspermroz/go-message-backend/internal/go-messages/ports/http/types"
)

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	types.ErrorMessage
}

func (e *ErrResponse) Render(_ http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		ErrorMessage: types.ErrorMessage{
			Status: "Invalid request",
			Error:  strToStrPtr(err.Error()),
		},
	}
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusUnprocessableEntity,
		ErrorMessage: types.ErrorMessage{
			Status: "Error rendering response",
			Error:  strToStrPtr(err.Error()),
		},
	}
}

func ErrInternalServerError(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusInternalServerError,
		ErrorMessage: types.ErrorMessage{
			Status: "internal server error",
			Error:  strToStrPtr(err.Error()),
		},
	}
}

func ErrBadRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		ErrorMessage: types.ErrorMessage{
			Status: "bad request",
			Error:  strToStrPtr(err.Error()),
		},
	}
}

func ErrUnauthorized(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusUnauthorized,
		ErrorMessage: types.ErrorMessage{
			Status: "unauthorized",
			Error:  strToStrPtr(err.Error()),
		},
	}
}

func ErrForbidden(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusForbidden,
		ErrorMessage: types.ErrorMessage{
			Status: "forbidden",
			Error:  strToStrPtr(err.Error()),
		},
	}
}

func ErrNotFound(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusNotFound,
		ErrorMessage: types.ErrorMessage{
			Status: "not found",
			Error:  strToStrPtr(err.Error()),
		},
	}
}

func strToStrPtr(in string) *string {
	if in == "" {
		return nil
	}
	return &in
}
