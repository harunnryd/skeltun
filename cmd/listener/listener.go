// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package listener

import (
	"github.com/harunnryd/skeltun/cmd/listener/provider"
	"github.com/harunnryd/skeltun/config"
	"github.com/harunnryd/skeltun/internal/app/repo"
	"github.com/harunnryd/skeltun/internal/app/usecase"
	"github.com/harunnryd/skeltun/internal/pkg"
	"github.com/harunnryd/skeltun/job"
	"os"
	"os/signal"

	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
)

// IListener ...
type IListener interface {
	// Start is used for starting the listener.
	Start()
}

// Listener ...
type Listener struct {
	config  config.IConfig
	redis   *redis.Pool
	repo    repo.IRepo
	usecase usecase.IUseCase
	pkg     pkg.IPkg
	job     job.IJob
	statement
}

type statement struct {
	workerPool *work.WorkerPool
}

// New ...
func New(opts ...Option) IListener {
	listener := new(Listener)
	for _, opt := range opts {
		opt(listener)
	}
	return listener
}

// Start is used for starting the listener.
func (listener *Listener) Start() {
	// Make a new statement.workerPool. Arguments:
	// Context{} is a struct that will be the context for the request.
	// 10 is the max concurrency
	// "app.name" is the Redis namespace
	// redis is a Redis pool
	if listener.statement.workerPool == nil {
		listener.statement.workerPool = work.NewWorkerPool(provider.Provider{}, 10, listener.config.GetString("app.name"), listener.redis)
	}

	var iProvider = provider.New(
		provider.WithConfig(listener.config),
		provider.WithRedis(listener.redis),
		provider.WithRepo(listener.repo),
		provider.WithUseCase(listener.usecase),
		provider.WithPkg(listener.pkg),
		provider.WithJob(listener.job),
	)

	// Add middleware that will be executed for each job
	listener.statement.workerPool.Middleware(iProvider.Log)

	// Example: Map the name of jobs to handler functions
	listener.statement.workerPool.Job("do_send_notification", iProvider.DoSendNotification)
	listener.statement.workerPool.Job("hcheck", iProvider.Hcheck)

	// Example: Customize options:
	listener.statement.workerPool.JobWithOptions("export", work.JobOptions{Priority: 10, MaxFails: 1}, iProvider.Export)

	// Start processing jobs
	listener.statement.workerPool.Start()

	// Wait for a signal to quit:
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)
	<-signalChan

	// Stop the pool
	listener.statement.workerPool.Stop()
}
