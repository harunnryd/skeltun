// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package errors

// ValidationError ...
type ValidationError struct {
	Err error
}

// Error ...
func (r *ValidationError) Error() string {
	return r.Err.Error()
}

// TimeoutError ...
type TimeoutError struct {
	Err error
}

// Error ...
func (r *TimeoutError) Error() string {
	return r.Err.Error()
}

// URLNotFoundError ...
type URLNotFoundError struct {
	Err error
}

// Error ...
func (r *URLNotFoundError) Error() string {
	return r.Err.Error()
}
