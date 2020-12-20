// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hcheck

import (
	"github.com/harunnryd/skeltun/config"
	"github.com/harunnryd/skeltun/internal/app/usecase"
)

// Option ...
type Option func(*Hcheck)

// WithConfig ...
func WithConfig(config config.IConfig) Option {
	return func(hcheck *Hcheck) {
		hcheck.config = config
	}
}

// WithUseCase ...
func WithUseCase(usecase usecase.IUseCase) Option {
	return func(hcheck *Hcheck) {
		hcheck.usecase = usecase
	}
}
