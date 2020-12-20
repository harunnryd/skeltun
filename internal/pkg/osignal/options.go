// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package osignal

import (
	"github.com/harunnryd/skeltun/config"
	"net/http"
)

// Option ...
type Option func(osignal *Osignal)

// WithNetClient ...
func WithNetClient(netClient *http.Client) Option {
	return func(osignal *Osignal) {
		osignal.netClient = netClient
	}
}

// WithConfig ...
func WithConfig(config config.IConfig) Option {
	return func(osignal *Osignal) {
		osignal.config = config
	}
}
