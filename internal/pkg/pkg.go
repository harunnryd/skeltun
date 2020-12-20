// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pkg

import (
	"github.com/harunnryd/skeltun/internal/pkg/osignal"
	"github.com/harunnryd/skeltun/internal/pkg/token"
)

// IPkg ...
type IPkg interface {
	// GetToken it returns instance of token.Token that implements token.IToken methods.
	GetToken() token.IToken

	// GetOsignal it returns instance of osignal.Osignal that implements osignal.IOsignal methods.
	GetOsignal() osignal.IOsignal
}

// Pkg ...
type Pkg struct {
	token   token.IToken
	osignal osignal.IOsignal
}

// New it returns instance of Pkg that implements IPkg methods.
func New(opts ...Option) IPkg {
	p := new(Pkg)
	for _, opt := range opts {
		opt(p)
	}
	return p
}

// GetToken it returns instance of token.Token that implements token.IToken methods.
func (p *Pkg) GetToken() token.IToken {
	return p.token
}

// GetOsignal it returns instance of osignal.Osignal that implements osignal.IOsignal methods.
func (p *Pkg) GetOsignal() osignal.IOsignal {
	return p.osignal
}
