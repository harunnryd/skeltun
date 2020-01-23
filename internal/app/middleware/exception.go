package middleware

import (
	"net/http"
	"skeltun/internal/pkg/http/presenter"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// URLNotFound ...
func URLNotFound(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())

		// Temporary routing context to look-ahead before routing the request
		tctx := chi.NewRouteContext()

		// Attempt to find a handler for the routing path, if not found,
		// throw the ErrURLNotFound as a response.
		if !rctx.Routes.Match(tctx, r.Method, r.URL.Path) {
			psenter := presenter.ErrURLNotFound

			render.Status(r, psenter.HTTPStatus)
			render.JSON(w, r, psenter)
			return
		}
		next.ServeHTTP(w, r)
	})
}
