package config

import (
	"github.com/spf13/viper"
)

// IConfig ...
type IConfig interface {
	GetString(string) string
	GetInt(string) int
	GetBool(string) bool
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
