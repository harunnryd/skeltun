// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type IToken interface {
	GetImplicitToken(payload Payload) (tokenStr string, expiresAt int64, err error)
	GetNewToken(payload Payload, refreshToken string) (tokenStr string, expiresAt int64, err error)
	GetRefreshToken(payload Payload) (tokenStr string, expiresAt int64, err error)
	IsTokenValid(tokenString string) (valid bool, err error)
	GetClaims(tokenString string) (claims jwt.MapClaims, err error)
}

// Token define struct for token
type Token struct {
	hmac   []byte
	issuer string
	ttl
}

type ttl struct {
	accessToken  uint64
	refreshToken uint64
}

// Init initialized token
func New(opts ...Options) IToken {
	token := new(Token)

	for _, opt := range opts {
		opt(token)
	}

	return token
}

// GetImplicitToken generate new access token for implicit grant
func (t *Token) GetImplicitToken(payload Payload) (tokenStr string, expiresAt int64, err error) {
	tokenStr, expiresAt, err = t.createToken(payload, t.ttl.refreshToken)
	return
}

// GetNewToken generate new access token
func (t *Token) GetNewToken(payload Payload, refreshToken string) (tokenStr string, expiresAt int64, err error) {
	valid, err := t.IsTokenValid(refreshToken)
	if err != nil {
		return
	}

	if !valid {
		err = errors.New("refresh token is not valid")
		return
	}

	tokenStr, expiresAt, err = t.createToken(payload, t.ttl.accessToken)
	return
}

// GetRefreshToken generate new refresh token
func (t *Token) GetRefreshToken(payload Payload) (tokenStr string, expiresAt int64, err error) {
	tokenStr, expiresAt, err = t.createToken(payload, t.ttl.refreshToken)
	return
}

func (t *Token) createToken(payload Payload, ttl uint64) (tokenStr string, expiresAt int64, err error) {
	jwtStandardClaims := jwt.StandardClaims{
		IssuedAt: time.Now().Unix(),
		Issuer:   t.issuer,
	}

	if ttl > 0 {
		expiresAt = time.Now().Add(time.Duration(ttl) * time.Minute).Unix()
		jwtStandardClaims.ExpiresAt = expiresAt
	}

	claims := Claims{
		Payload: Payload{
			UserID: payload.UserID,
		},
		StandardClaims: jwtStandardClaims,
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err = jwtToken.SignedString(t.hmac)
	return
}

// IsTokenValid check whether token still valid
func (t *Token) IsTokenValid(tokenString string) (valid bool, err error) {
	token, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", jwtToken.Header["alg"])
		}

		return t.hmac, nil
	})

	if err != nil {
		return
	}

	valid = token.Valid
	return
}

// GetClaims get token claims
func (t *Token) GetClaims(tokenString string) (claims jwt.MapClaims, err error) {
	token, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", jwtToken.Header["alg"])
		}

		return t.hmac, nil
	})

	if err != nil {
		return
	}

	claims = token.Claims.(jwt.MapClaims)

	return
}
