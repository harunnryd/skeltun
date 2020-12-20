// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package middleware

import (
	"context"
	"net/http"
	"time"
)

// TimeoutContext ...
func TimeoutContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), (120 * time.Second))
		defer func() {
			cancel()
		}()

		// This gives you a copy of the request with a the request context
		// changed to the new context with the 5-second timeout created
		// above.
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
