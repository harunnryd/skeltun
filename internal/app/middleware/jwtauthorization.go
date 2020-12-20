// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package middleware

import (
	"context"
	"github.com/harunnryd/skeltun/internal/pkg/http/customrest/customwriter"
	"github.com/harunnryd/skeltun/internal/pkg/token"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// JWTAuthorization ...
func JWTAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			customwriter.New().WriteError(w, r, jwt.NewValidationError("Invalid Authorization Header", jwt.ValidationErrorUnverifiable))
			return
		}

		tokenStr := strings.Replace(authorizationHeader, "Bearer ", "", -1)
		claims, err := token.New(token.WithSecretKey("THE-P0W3RRAN63R")).GetClaims(tokenStr)
		if err != nil {
			customwriter.New().WriteError(w, r, err)
			return
		}

		ctx := context.WithValue(r.Context(), "claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
