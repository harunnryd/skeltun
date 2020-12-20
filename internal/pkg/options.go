// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pkg

import (
	"github.com/harunnryd/skeltun/config"
	"github.com/harunnryd/skeltun/internal/pkg/osignal"
	"github.com/harunnryd/skeltun/internal/pkg/token"
	"net"
	"net/http"
	"time"
)

// Option ...
type Option func(p *Pkg)

// WithDependency ...
func WithDependency(config config.IConfig) Option {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 5 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 5 * time.Second,
		},
	}
	return func(p *Pkg) {
		p.token = token.New(
			token.WithIssuer("PT.Bangun Sistem Digital"),
			token.WithSecretKey("THE-P0W3RRAN63R"),
		)
		p.osignal = osignal.New(
			osignal.WithConfig(config),
			osignal.WithNetClient(httpClient),
		)
	}
}
