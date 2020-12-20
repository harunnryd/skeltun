// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package usecase

import (
	"github.com/harunnryd/skeltun/internal/app/usecase/hcheck"
)

// IUseCase ...
type IUseCase interface {
	// GetHcheck it returns instance of Hcheck that implements IHcheck methods.
	GetHcheck() hcheck.IHcheck
}

// UseCase ...
type UseCase struct {
	hcheck hcheck.IHcheck
}

// New ...
func New(opts ...Option) IUseCase {
	usecase := new(UseCase)
	for _, opt := range opts {
		opt(usecase)
	}
	return usecase
}

// GetHcheck it returns instance of Hcheck that implements IHcheck methods.
func (usecase *UseCase) GetHcheck() hcheck.IHcheck {
	return usecase.hcheck
}
