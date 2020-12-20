// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package handler

import (
	"github.com/harunnryd/skeltun/config"
	"github.com/harunnryd/skeltun/internal/app/handler/hcheck"
	"github.com/harunnryd/skeltun/internal/app/usecase"
)

// Option ...
type Option func(*Handler)

// WithHandler ...
func WithHandler(config config.IConfig) Option {
	iUsecase := usecase.New(usecase.WithDependency(config))
	return func(handler *Handler) {
		// Inject all your handler's in here.
		// Example :
		// handler.user = user.New(
		//     user.WithConfig(config),
		//     user.WithUseCase(iUsecase),
		// )
		handler.hcheck = hcheck.New(
			hcheck.WithConfig(config),
			hcheck.WithUseCase(iUsecase),
		)
	}
}
