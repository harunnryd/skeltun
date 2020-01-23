package rest

import (
	"net/http"
)

// Handler ...
type Handler func(w http.ResponseWriter, r *http.Request) (interface{}, error)

// ServeHTTP ...
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := h(w, r)

	// When err not nil throw the respondwithError method. however 
	// throw the respondwithSuccess method if err equal nil.
	if err != nil {
		respondwithError(w, r, err)
		return
	}
	respondwithSuccess(w, r, data)
}

// HTTPAction ...
type HTTPAction struct {
	HTTPMethod string
	Pattern    string
	H          Handler
}

// New ...
func New(hmethod, pattern string, handler Handler) HTTPAction {
	return HTTPAction{
		HTTPMethod: hmethod,
		Pattern:    pattern,
		H:          handler,
	}
}
