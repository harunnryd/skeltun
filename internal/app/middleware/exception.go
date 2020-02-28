package middleware

import (
	"net/http"
	"skeltun/internal/pkg/http/presenter"
	"skeltun/internal/pkg/http/rest/customwriter"

	"github.com/go-chi/chi"
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
			customwriter.New().WriteError(w, r, presenter.ErrURLNotFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}
