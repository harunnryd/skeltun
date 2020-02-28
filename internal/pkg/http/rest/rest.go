package rest

import (
	"net/http"
	"skeltun/internal/pkg/http/rest/customwriter"
)

// Handler ...
type Handler func(w http.ResponseWriter, r *http.Request) (interface{}, error)

// IRest ...
type IRest interface {
	GetHTTPMethod() string
	GetHandler() Handler
	GetPattern() string
}

// Rest ...
type Rest struct {
	HTTPMethod string
	Pattern    string
	H          Handler
}

// New ...
func New(opts ...Option) IRest {
	rest := new(Rest)
	for _, opt := range opts {
		opt(rest)
	}
	return rest
}

// ServeHTTP ...
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := h(w, r)

	// When err not nil throw the respondwithError method. however
	// throw the respondwithSuccess method if err equal nil.
	if err != nil {
		customwriter.New().WriteError(w, r, err)
		return
	}
	customwriter.New().WriteSuccess(w, r, data)
}

// GetHandler ...
func (rest *Rest) GetHandler() Handler {
	return rest.H
}

// GetHTTPMethod ...
func (rest *Rest) GetHTTPMethod() string {
	return rest.HTTPMethod
}

// GetPattern ...
func (rest *Rest) GetPattern() string {
	return rest.Pattern
}
