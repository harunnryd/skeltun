// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Option ...
type Option func(*Config)

// WithEnvSetup ...
func WithEnvSetup() Option {
	return func(cfg *Config) {
		cfg.env = viper.New()
		cfg.env.AddConfigPath(".")
		cfg.env.AddConfigPath("params")
		cfg.env.SetConfigName("env")
		cfg.env.SetConfigType("yaml")

		// Check read process
		if err := cfg.env.ReadInConfig(); err != nil {
			panic(fmt.Errorf("Config error: %s", err))
		}

		fmt.Printf("=> config file: %s\n", cfg.env.ConfigFileUsed())
	}
}
