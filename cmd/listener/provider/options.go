// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package provider

import (
	"github.com/harunnryd/skeltun/config"
	"github.com/harunnryd/skeltun/internal/app/repo"
	"github.com/harunnryd/skeltun/internal/app/usecase"
	"github.com/harunnryd/skeltun/internal/pkg"
	"github.com/harunnryd/skeltun/job"

	"github.com/gomodule/redigo/redis"
)

// Option ...
type Option func(provider *Provider)

// WithConfig ...
func WithConfig(config config.IConfig) Option {
	return func(provider *Provider) {
		provider.config = config
	}
}

// WithRedis ...
func WithRedis(redis *redis.Pool) Option {
	return func(provider *Provider) {
		provider.redis = redis
	}
}

// WithRepo ...
func WithRepo(repo repo.IRepo) Option {
	return func(provider *Provider) {
		provider.repo = repo
	}
}

// WithUseCase ...
func WithUseCase(usecase usecase.IUseCase) Option {
	return func(provider *Provider) {
		provider.usecase = usecase
	}
}

// WithPkg ...
func WithPkg(pkg pkg.IPkg) Option {
	return func(provider *Provider) {
		provider.pkg = pkg
	}
}

// WithJob ...
func WithJob(job job.IJob) Option {
	return func(provider *Provider) {
		provider.job = job
	}
}
