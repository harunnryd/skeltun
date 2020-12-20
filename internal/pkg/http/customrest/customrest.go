// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package customrest

import (
	"github.com/harunnryd/skeltun/internal/pkg/http/customrest/customwriter"
	"net/http"
)

// Handler ...
type Handler func(w http.ResponseWriter, r *http.Request) (interface{}, error)

// ICustomRest ...
type ICustomRest interface {
	GetHTTPMethod() string
	GetHandler() Handler
	GetPattern() string
}

// CustomRest ...
type CustomRest struct {
	HTTPMethod string
	Pattern    string
	H          Handler
}

// New ...
func New(opts ...Option) ICustomRest {
	rest := new(CustomRest)
	for _, opt := range opts {
		opt(rest)
	}
	return rest
}

// ServeHTTP ...
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := h(w, r)

	// When err not nil throw the WriteError method. however
	// throw the WriteSuccess method if err equal nil.
	if err != nil {
		customwriter.New().WriteError(w, r, err)
		return
	}
	customwriter.New().WriteSuccess(w, r, data)
}

// GetHandler ...
func (rest *CustomRest) GetHandler() Handler {
	return rest.H
}

// GetHTTPMethod ...
func (rest *CustomRest) GetHTTPMethod() string {
	return rest.HTTPMethod
}

// GetPattern ...
func (rest *CustomRest) GetPattern() string {
	return rest.Pattern
}
