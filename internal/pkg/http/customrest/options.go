// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package customrest

// Option ...
type Option func(*CustomRest)

// WithHTTPMethod ...
func WithHTTPMethod(httpMethod string) Option {
	return func(rest *CustomRest) {
		rest.HTTPMethod = httpMethod
	}
}

// WithPattern ...
func WithPattern(pattern string) Option {
	return func(rest *CustomRest) {
		rest.Pattern = pattern
	}
}

// WithHandler ...
func WithHandler(h Handler) Option {
	return func(rest *CustomRest) {
		rest.H = h
	}
}
