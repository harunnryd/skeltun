// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hcheck

import (
	"github.com/harunnryd/skeltun/config"
	"github.com/harunnryd/skeltun/internal/app/repo"
	"github.com/harunnryd/skeltun/internal/pkg"
	"github.com/harunnryd/skeltun/job"
)

// Option ...
type Option func(*Hcheck)

// WithConfig ...
func WithConfig(config config.IConfig) Option {
	return func(hcheck *Hcheck) {
		hcheck.config = config
	}
}

// WithRepo ...
func WithRepo(repo repo.IRepo) Option {
	return func(hcheck *Hcheck) {
		hcheck.repo = repo
	}
}

// WithPkg ...
func WithPkg(pkg pkg.IPkg) Option {
	return func(hcheck *Hcheck) {
		hcheck.pkg = pkg
	}
}

// WithJob ...
func WithJob(job job.IJob) Option {
	return func(hcheck *Hcheck) {
		hcheck.job = job
	}
}
