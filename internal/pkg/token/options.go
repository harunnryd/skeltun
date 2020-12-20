// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package token

// Options define functional options for token
type Options func(*Token)

// WithSecretKey assign secret key to token
func WithSecretKey(secretKey string) Options {
	return func(token *Token) {
		token.hmac = []byte(secretKey)
	}
}

// WithAccessTokenTTL assign access token ttl
func WithAccessTokenTTL(ttl uint64) Options {
	return func(token *Token) {
		token.ttl.accessToken = ttl
	}
}

// WithRefreshTokenTTL assign refresh token ttl
func WithRefreshTokenTTL(ttl uint64) Options {
	return func(token *Token) {
		token.ttl.refreshToken = ttl
	}
}

// WithIssuer assign issuer
func WithIssuer(issue string) Options {
	return func(token *Token) {
		token.issuer = issue
	}
}
