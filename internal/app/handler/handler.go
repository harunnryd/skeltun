// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package handler

import (
	"github.com/harunnryd/skeltun/internal/app/handler/hcheck"
)

// IHandler ...
type IHandler interface {
	// GetHcheck it returns instance of Hcheck that implements IHcheck methods.
	GetHcheck() hcheck.IHcheck
}

// Handler ...
type Handler struct {
	hcheck hcheck.IHcheck
}

// New ...
func New(opts ...Option) IHandler {
	handler := new(Handler)
	for _, opt := range opts {
		opt(handler)
	}
	return handler
}

// GetHcheck it returns instance of Hcheck that implements IHcheck methods.
func (handler *Handler) GetHcheck() hcheck.IHcheck {
	return handler.hcheck
}
