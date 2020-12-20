// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package middleware

import (
	"errors"
	customerror "github.com/harunnryd/skeltun/internal/pkg/errors"
	"github.com/harunnryd/skeltun/internal/pkg/http/customrest/customwriter"
	"net/http"

	"github.com/go-chi/chi"
)

// URLNotFound ...
func URLNotFound(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())

		// Temporary routing context to look-ahead before routing the request
		tctx := chi.NewRouteContext()

		// Attempt to find a handler for the routing path, if not found,
		// throw the ErrURLNotFound as a response.
		if !rctx.Routes.Match(tctx, r.Method, r.URL.Path) {
			customwriter.New().WriteError(w, r, &customerror.URLNotFoundError{Err: errors.New("URL not found")})
			return
		}

		next.ServeHTTP(w, r)
	})
}
