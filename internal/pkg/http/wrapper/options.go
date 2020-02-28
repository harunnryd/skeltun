package wrapper

import "github.com/go-chi/chi"

// Option ...
type Option func(*Wrapper)

// WithRouter ...
func WithRouter(router chi.Router) Option {
	return func(wrapper *Wrapper) {
		wrapper.Router = router
	}
}
