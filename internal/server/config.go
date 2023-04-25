package server

import (
	"context"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	DebugMode bool   `env:"DEBUG_MODE,default=false"`
	Host      string `env:"HOST,default=0.0.0.0"`
	Port      string `env:"PORT,default=5330"`

	DokkuSSHHost string `env:"DOKKU_SSH_HOST,default=127.0.0.1"`
	DokkuSSHPort string `env:"DOKKU_SSH_PORT,default=22"`

	DBPath string `env:"DB_PATH,default=/data/shokku.db"`

	AuthTokenLifetimeMinutes int `env:"TOKEN_LIFETIME_MINS,default=15"`
}

func LoadConfig() (Config, error) {
	ctx := context.Background()
	var cfg Config
	return cfg, envconfig.Process(ctx, &cfg)
}
