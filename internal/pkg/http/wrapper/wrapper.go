package wrapper

import (
	"net/http"
	"skeltun/internal/pkg/http/rest"

	"github.com/go-chi/chi"
)

// IWrapper ...
type IWrapper interface {
	chi.Router
	Action(...rest.HTTPAction)
}

// Wrapper wraps chi.Router to pre-process http.Handler and add support for ActionHandler
type Wrapper struct {
	chi.Router
	PrepareHandler rest.Handler
}

// New ...
func New(router chi.Router) IWrapper {
	return &Wrapper{
		Router: router,
	}
}

var _ chi.Router = &Wrapper{}

func (r *Wrapper) copy(router chi.Router) IWrapper {
	return &Wrapper{
		Router:         router,
		PrepareHandler: r.PrepareHandler,
	}
}

// With adds inline middlewares for an endpoint handler
func (r *Wrapper) With(middlewares ...func(http.Handler) http.Handler) chi.Router {
	return r.copy(r.Router.With(middlewares...))
}

// Group adds a new inline-Router along the current routing
// path, with a fresh middleware stack for the inline-Router
func (r *Wrapper) Group(fn func(r chi.Router)) chi.Router {
	im := r.copy(r.With())
	if fn != nil {
		fn(im)
	}
	return im
}

// Route mounts a sub-Router along a `pattern`` string.
func (r *Wrapper) Route(pattern string, fn func(r chi.Router)) chi.Router {
	subRouter := r.copy(chi.NewRouter())
	if fn != nil {
		fn(subRouter)
	}
	r.Mount(pattern, subRouter)
	return subRouter
}

// Mount attaches another http.Handler along ./pattern/*
func (r *Wrapper) Mount(pattern string, handler http.Handler) {
	r.Router.Mount(pattern, handler)
}

// Handle adds routes for `pattern` that matches all HTTP methods
func (r *Wrapper) Handle(pattern string, handler http.Handler) {
	r.Router.Handle(pattern, handler)
}

// Method adds routes for `pattern` that matches the `method` HTTP method
func (r *Wrapper) Method(method, pattern string, handler http.Handler) {
	r.Router.Method(method, pattern, handler)
}

// Action adds one or more HTTPAction for `h.Pattern()` that matches the `h.HTTPMethod()` HTTP method
func (r *Wrapper) Action(handlers ...rest.HTTPAction) {
	for _, handler := range handlers {
		r.Router.Method(handler.HTTPMethod, handler.Pattern, handler.H)
	}
}
