// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package token

import (
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/uuid"
)

// Payload define payload body for token
type Payload struct {
	UserID uuid.UUID
}

// Claims define custom jwt claims
type Claims struct {
	Payload
	jwt.StandardClaims
}
