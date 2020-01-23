package handler

import "skeltun/internal/app/handler/hcheck"

// IHandler ...
type IHandler interface {
	Hcheck() hcheck.IHcheck
}

// Handler ...
type Handler struct {
	hcheck hcheck.IHcheck
}

// New ...
func New(opts ...Option) IHandler {
	handler := new(Handler)
	for _, opt := range opts {
		opt(handler)
	}
	return handler
}

// Hcheck ...
func (handler *Handler) Hcheck() hcheck.IHcheck {
	return handler.hcheck
}
