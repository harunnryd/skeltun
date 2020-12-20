// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import (
	"github.com/harunnryd/skeltun/internal/app/handler"
	"log"
	"time"
)

// Option ...
type Option func(*Server)

// WithDefault ...
func WithDefault(logger *log.Logger, addr string, handler handler.IHandler, readTimeout, writeTimeout, idleTimeout int) Option {
	return func(s *Server) {
		s.ErrorLog = logger
		s.Addr = addr
		s.Handler = handler
		s.ReadTimeout = time.Second * time.Duration(readTimeout)
		s.WriteTimeout = time.Second * time.Duration(writeTimeout)
		s.IdleTimeout = time.Second * time.Duration(idleTimeout)
	}
}
