package server

import (
	"net/http"
	"skeltun/internal/app/handler"
	"skeltun/internal/app/middleware"
	"skeltun/internal/pkg/http/rest"
	"skeltun/internal/pkg/http/wrapper"

	"github.com/go-chi/chi"
)

// Router ...
func (s *Server) Router(handler handler.IHandler) (w wrapper.IWrapper) {
	w = wrapper.New(wrapper.WithRouter(chi.NewRouter()))
	w.Use(middleware.URLNotFound)
	w.Route("/v1", func(r chi.Router) {
		router := r.(wrapper.IWrapper)
		router.Action(
			rest.New(
				rest.WithHTTPMethod(http.MethodGet),
				rest.WithPattern("/hc"),
				rest.WithHandler(handler.Hcheck().HealthCheck),
			),
		)
		router.Action(
			rest.New(
				rest.WithHTTPMethod(http.MethodGet),
				rest.WithPattern("/attendance-books"),
				rest.WithHandler(handler.Attendance().AttendanceBook),
			),
		)
	})
	return
}
