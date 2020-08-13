package apiserver

import "github.com/shuhrat-shokirov/http-rest-api/internal/app/store"

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8181",
		LogLevel: "debug",
		Store: store.NewConfig(),
	}
}
