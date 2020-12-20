// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package job

import (
	"github.com/harunnryd/skeltun/config"
	"log"

	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
)

// IJob ...
type IJob interface {
	Dispatch(jobName string, secondsFromNow int64, args work.Q)
	Queue(jobName string, args work.Q)
}

// Job ...
type Job struct {
	config config.IConfig
	redis  *redis.Pool
	statement
}

type statement struct {
	enqueuer *work.Enqueuer
}

// New ...
func New(opts ...Option) IJob {
	j := new(Job)
	for _, opt := range opts {
		opt(j)
	}
	return j
}

// Queue ...
// Make a redis pool
// &redis.Pool{
// 	MaxActive: 5,
// 	MaxIdle:   5,
// 	Wait:      true,
// 	Dial: func() (redis.Conn, error) {
// 		return redis.Dial("tcp", ":6379", redis.DialPassword("powerrangers"))
// 	},
// }
func (j *Job) Queue(jobName string, args work.Q) {
	if j.statement.enqueuer == nil {
		j.statement.enqueuer = work.NewEnqueuer(j.config.GetString("app.name"), j.redis)
	}

	_, err := j.statement.enqueuer.Enqueue(jobName, args)
	if err != nil {
		log.Fatal(err)
	}
}

// Dispatch ...
// Make a redis pool
// &redis.Pool{
// 	MaxActive: 5,
// 	MaxIdle:   5,
// 	Wait:      true,
// 	Dial: func() (redis.Conn, error) {
// 		return redis.Dial("tcp", ":6379", redis.DialPassword("powerrangers"))
// 	},
// }
func (j *Job) Dispatch(jobName string, secondsFromNow int64, args work.Q) {
	if j.statement.enqueuer == nil {
		j.statement.enqueuer = work.NewEnqueuer(j.config.GetString("app.name"), j.redis)
	}

	_, err := j.statement.enqueuer.EnqueueIn(jobName, secondsFromNow, args)
	if err != nil {
		log.Fatal(err)
	}
}
