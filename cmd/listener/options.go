// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package listener

import (
	"github.com/harunnryd/skeltun/config"
	"github.com/harunnryd/skeltun/internal/app/repo"
	"github.com/harunnryd/skeltun/internal/app/usecase"
	"github.com/harunnryd/skeltun/internal/pkg"
	"github.com/harunnryd/skeltun/job"

	"github.com/gomodule/redigo/redis"
)

// Option ...
type Option func(listener *Listener)

// WithConfig ...
func WithConfig(config config.IConfig) Option {
	return func(listener *Listener) {
		listener.config = config
	}
}

// WithRedis ...
func WithRedis(redis *redis.Pool) Option {
	return func(listener *Listener) {
		listener.redis = redis
	}
}

// WithRepo ...
func WithRepo(repo repo.IRepo) Option {
	return func(listener *Listener) {
		listener.repo = repo
	}
}

// WithUseCase ...
func WithUseCase(usecase usecase.IUseCase) Option {
	return func(listener *Listener) {
		listener.usecase = usecase
	}
}

// WithPkg ...
func WithPkg(pkg pkg.IPkg) Option {
	return func(listener *Listener) {
		listener.pkg = pkg
	}
}

// WithJob ...
func WithJob(job job.IJob) Option {
	return func(listener *Listener) {
		listener.job = job
	}
}
