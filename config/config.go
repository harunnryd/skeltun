// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"time"

	"github.com/spf13/viper"
)

// IConfig ...
type IConfig interface {
	GetString(string) string
	GetInt(string) int
	GetBool(string) bool
	GetDuration(string) time.Duration
	GetFloat64(string) float64
}

// Config ...
type Config struct {
	env *viper.Viper
}

// New ...
func New(opts ...Option) IConfig {
	config := new(Config)
	for _, opt := range opts {
		opt(config)
	}
	return config
}

// GetString ...
func (cfg *Config) GetString(k string) string {
	return cfg.env.GetString(k)
}

// GetInt ...
func (cfg *Config) GetInt(k string) int {
	return cfg.env.GetInt(k)
}

// GetBool ...
func (cfg *Config) GetBool(k string) bool {
	return cfg.env.GetBool(k)
}

// GetDuration ...
func (cfg *Config) GetDuration(k string) time.Duration {
	return cfg.env.GetDuration(k)
}

// GetFloat64 ...
func (cfg *Config) GetFloat64(k string) float64 {
	return cfg.env.GetFloat64(k)
}
