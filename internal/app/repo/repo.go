// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package repo

import (
	"github.com/harunnryd/skeltun/internal/app/repo/hcheck"
)

// IRepo ...
type IRepo interface {
	// GetHcheck it returns instance of Hcheck that implements methods.
	GetHcheck() hcheck.IHcheck
}

// Repo ...
type Repo struct {
	hcheck hcheck.IHcheck
}

// New ...
func New(opts ...Option) IRepo {
	repo := new(Repo)
	for _, opt := range opts {
		opt(repo)
	}
	return repo
}

// GetHcheck it returns instance of Hcheck that implements methods.
func (repo *Repo) GetHcheck() hcheck.IHcheck {
	return repo.hcheck
}
