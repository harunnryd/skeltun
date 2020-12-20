// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package osignal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/harunnryd/skeltun/config"
	"github.com/harunnryd/skeltun/internal/pkg/osignal/param"
	"github.com/harunnryd/skeltun/internal/pkg/osignal/transporter"
	"net/http"
	"time"
)

// IOsignal is an interface that stores the methods that Osignal struct will use.
type IOsignal interface {
	// DoSendNotification is used for sending notifications with Onesignal.
	// It returns transporter and any errors written.
	DoSendNotification(params param.DoSendNotification) (transporter transporter.DoSendNotification, err error)
}

// Osignal is an struct that implements IOsignal methods.
type Osignal struct {
	config      config.IConfig
	contentType string
	netClient   *http.Client
	statement
}

type statement struct {
	request          *http.Request
	response         *http.Response
	currentTimestamp time.Time
}

// New it returns instance of Osignal that implements IOsignal methods.
func New(opts ...Option) IOsignal {
	osignal := new(Osignal)
	for _, opt := range opts {
		opt(osignal)
	}
	return osignal
}

func (osignal *Osignal) getPayload(v interface{}) (jsonMarshal []byte) {
	jsonMarshal, _ = json.Marshal(v)
	return
}

// DoSendNotification is used for sending notifications with Onesignal.
// It returns transporter and any errors written.
func (osignal *Osignal) DoSendNotification(params param.DoSendNotification) (transporter transporter.DoSendNotification, err error) {
	osignal.statement.request, err = http.NewRequest("POST", osignal.config.GetString("onesignal.uri.create"), bytes.NewBuffer(osignal.getPayload(params)))
	if err != nil {
		return
	}

	osignal.statement.request.Header.Set("Authorization", fmt.Sprint("Basic ", osignal.config.GetString("onesignal.api.key")))
	osignal.statement.request.Header.Set("Content-Type", "application/json")

	osignal.statement.response, err = osignal.netClient.Do(osignal.statement.request)
	if err != nil {
		return
	}

	defer osignal.statement.response.Body.Close()

	err = json.NewDecoder(osignal.statement.response.Body).Decode(&transporter)
	if err != nil {
		return
	}

	return
}
