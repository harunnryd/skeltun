// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package success

import (
	"reflect"
)

// Option ...
type Option func(*Success)

// WithHTTPStatus ...
func WithHTTPStatus(HTTPStatus int) Option {
	return func(success *Success) {
		success.HTTPStatus = HTTPStatus
	}
}

// WithResponseCode ...
func WithResponseCode(responseCode string) Option {
	return func(success *Success) {
		success.ResponseCode = responseCode
	}
}

// WithResponseDesc ...
func WithResponseDesc(responseDesc string) Option {
	return func(success *Success) {
		success.ResponseDesc = responseDesc
	}
}

// WithData ...
func WithData(data interface{}) Option {
	return func(success *Success) {
		val := reflect.ValueOf(data)
		if val.Kind() == reflect.Slice && val.Len() < 1 {
			success.Data = []string{}
			return
		}

		if val.Kind() == reflect.Struct && val.IsZero() {
			success.Data = nil
			return
		}

		success.Data = data
	}
}
