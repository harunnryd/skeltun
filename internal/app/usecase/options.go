// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package usecase

import (
	"github.com/harunnryd/skeltun/config"
	"github.com/harunnryd/skeltun/internal/app/repo"
	"github.com/harunnryd/skeltun/internal/app/usecase/hcheck"
	"github.com/harunnryd/skeltun/internal/pkg"
	"github.com/harunnryd/skeltun/job"

	"github.com/gomodule/redigo/redis"
)

// Option ...
type Option func(*UseCase)

// WithDependency ...
func WithDependency(config config.IConfig) Option {
	var iRepo = repo.New(repo.WithDependency(config))
	var iPkg = pkg.New(pkg.WithDependency(config))
	var redisPool = &redis.Pool{
		MaxActive: 5,
		MaxIdle:   5,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			return redis.DialURL(config.GetString("database.redis.hosts"))
		},
	}
	var iJob = job.New(job.WithConfig(config), job.WithRedis(redisPool))

	return func(usecase *UseCase) {
		// Inject all your UseCase's in here.
		// Example :
		// usecase.user = user.New(
		//    user.WithConfig(config),
		//    user.WithRepo(repo),
		// )
		usecase.hcheck = hcheck.New(
			hcheck.WithConfig(config),
			hcheck.WithRepo(iRepo),
			hcheck.WithPkg(iPkg),
			hcheck.WithJob(iJob),
		)
	}
}
