package render

import (
	"net/http"

	"github.com/go-chi/render"
)

func MustRender(w http.ResponseWriter, r *http.Request, renderer render.Renderer) {
	err := render.Render(w, r, renderer)
	if err != nil {
		_ = render.Render(w, r, ErrRender(err))
		// doesn't make sense to handle this error, we assume it should render OK
	}
}
