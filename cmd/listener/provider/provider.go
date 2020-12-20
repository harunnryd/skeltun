// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package provider

import (
	"fmt"
	"github.com/harunnryd/skeltun/config"
	"github.com/harunnryd/skeltun/internal/app/repo"
	"github.com/harunnryd/skeltun/internal/app/usecase"
	"github.com/harunnryd/skeltun/internal/pkg"
	"github.com/harunnryd/skeltun/internal/pkg/osignal/param"
	"github.com/harunnryd/skeltun/internal/pkg/osignal/transporter"
	"github.com/harunnryd/skeltun/job"

	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
)

// IProvider is an interface that stores the methods that Provider struct will use.
type IProvider interface {
	// Log is used for printing the Log.
	// It returns any errors written.
	Log(job *work.Job, next work.NextMiddlewareFunc) error

	// Export is used for Checkins the queue.
	// It returns any errors written.
	Export(job *work.Job) error

	// DoSendNotification is used for sending notifications with Onesignal.
	// It returns any errors written.
	DoSendNotification(job *work.Job) (err error)

	// Hcheck is used for checking health of redis connection.
	// It returns any errors written.
	Hcheck(job *work.Job) (err error)
}

// Provider is an struct that implements IProvider methods.
type Provider struct {
	config  config.IConfig
	redis   *redis.Pool
	repo    repo.IRepo
	usecase usecase.IUseCase
	pkg     pkg.IPkg
	job     job.IJob
}

// New it returns instance of Provider that implements IProvider methods.
func New(opts ...Option) IProvider {
	provider := new(Provider)
	for _, opt := range opts {
		opt(provider)
	}
	return provider
}

// Log is used for printing the Log.
// It returns any errors written.
func (provider *Provider) Log(job *work.Job, next work.NextMiddlewareFunc) error {
	fmt.Println("Starting job: ", job.Name)
	return next()
}

// Export is used for Checkins the queue.
// It returns any errors written.
func (provider *Provider) Export(job *work.Job) error {
	return nil
}

// DoSendNotification is used for sending notifications with Onesignal.
// It returns any errors written.
func (provider *Provider) DoSendNotification(job *work.Job) (err error) {
	var includeExternalUserIDs = job.Args["include_external_user_ids"].([]interface{})
	var contents = job.Args["contents"].(map[string]interface{})
	if err = job.ArgError(); err != nil {
		return
	}

	var doSendNotificationParams = param.DoSendNotificationParam()
	doSendNotificationParams = param.DoSendNotification{
		AppID:                  provider.config.GetString("onesignal.api.app_id"),
		IncludeExternalUserIDs: includeExternalUserIDs,
		Contents:               contents,
	}

	var doSendNotificationResponse = transporter.DoSendNotificationTransporter()
	doSendNotificationResponse, err = provider.pkg.GetOsignal().DoSendNotification(doSendNotificationParams)
	if err != nil {
		return
	}

	fmt.Printf("%+v\n", doSendNotificationResponse)
	return
}

// Hcheck is used for checking health of redis connection.
// It returns any errors written.
func (provider *Provider) Hcheck(job *work.Job) (err error) {
	var responseCode = job.ArgString("response_code")
	var responseDesc = job.ArgString("response_desc")
	if err = job.ArgError(); err != nil {
		return
	}

	fmt.Printf("response_code: %s\n", responseCode)
	fmt.Printf("response_desc: %s\n", responseDesc)
	return
}
