package apiserver

import "github.com/Brotiger/https_rest_api_go.git/internal/app/store"

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: ":debug",
		Store: store.NewConfig(),
	}
}