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
	w = wrapper.New(chi.NewRouter())
	w.Use(middleware.URLNotFound)
	w.Route("/v1", func(r chi.Router) {
		router := r.(wrapper.IWrapper)
		router.Action(rest.New(http.MethodGet, "/hc", handler.Hcheck().HealthCheck))
	})

	return
}
