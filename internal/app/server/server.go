// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import (
	"context"
	"github.com/harunnryd/skeltun/internal/app/handler"
	"github.com/harunnryd/skeltun/internal/pkg/http/wrapper"
	"log"
	"net/http"
	"os"
	"time"
)

// IServer ...
type IServer interface {
	Router(handler handler.IHandler) (w wrapper.IWrapper)
	GetHTTPServer() *http.Server
	GracefullShutdown(server *http.Server, logger *log.Logger, quit <-chan os.Signal, done chan<- bool)
}

// Server will create a http.Server from the Go standard library
type Server struct {
	ErrorLog     *log.Logger
	Addr         string
	Handler      handler.IHandler
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

// New ...
func New(opts ...Option) IServer {
	srv := new(Server)
	for _, opt := range opts {
		opt(srv)
	}
	return srv
}

// GetHTTPServer ...
func (s *Server) GetHTTPServer() *http.Server {
	return &http.Server{
		ErrorLog:     s.ErrorLog,
		Addr:         s.Addr,
		Handler:      s.Router(s.Handler),
		ReadTimeout:  s.ReadTimeout,
		WriteTimeout: s.WriteTimeout,
		IdleTimeout:  s.IdleTimeout,
	}
}

// GracefullShutdown ...
func (s *Server) GracefullShutdown(server *http.Server, logger *log.Logger, quit <-chan os.Signal, done chan<- bool) {
	<-quit
	logger.Println("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	server.SetKeepAlivesEnabled(false)
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalf("Could not gracefully shutdown the server: %v\n", err)
	}
	close(done)
}
