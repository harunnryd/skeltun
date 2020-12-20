// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package success

// ISuccess ...
type ISuccess interface {
	GetResponseCode() string
	GetResponseDesc() string
	GetData() interface{}
	GetHTTPStatus() int
}

// Success ...
type Success struct {
	ResponseCode string      `json:"response_code"`
	ResponseDesc string      `json:"response_desc"`
	Data         interface{} `json:"data,omitempty"`
	HTTPStatus   int         `json:"-"`
}

// New ...
func New(opts ...Option) ISuccess {
	success := new(Success)
	for _, opt := range opts {
		opt(success)
	}
	return success
}

// GetResponseCode ...
func (success *Success) GetResponseCode() string {
	return success.ResponseCode
}

// GetResponseDesc ...
func (success *Success) GetResponseDesc() string {
	return success.ResponseDesc
}

// GetData ...
func (success *Success) GetData() interface{} {
	return success.Data
}

// GetHTTPStatus ...
func (success *Success) GetHTTPStatus() int {
	return success.HTTPStatus
}
