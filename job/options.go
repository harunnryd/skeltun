// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package job

import (
	"github.com/harunnryd/skeltun/config"

	"github.com/gomodule/redigo/redis"
)

// Option ...
type Option func(j *Job)

// WithRedis ...
func WithRedis(redis *redis.Pool) Option {
	return func(j *Job) {
		j.redis = redis
	}
}

// WithConfig ...
func WithConfig(config config.IConfig) Option {
	return func(j *Job) {
		j.config = config
	}
}
