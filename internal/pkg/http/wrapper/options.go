// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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
