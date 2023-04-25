package server

import (
	"fmt"
	"gitlab.com/texm/shokku/internal/server/auth"
	"time"
)

type initAuthConfig struct {
	SigningKey    []byte
	TokenLifetime time.Duration
	Method        auth.Method
	DebugMode     bool
	IsSetup       bool
}

func initAuthenticator(cfg initAuthConfig) (auth.Authenticator, error) {
	authCfg := auth.Config{
		SigningKey:    cfg.SigningKey,
		TokenLifetime: cfg.TokenLifetime,
	}

	if !cfg.IsSetup {
		return auth.NewNoneAuthenticator(authCfg)
	}

	bCryptCost := 14
	if cfg.DebugMode {
		// make hashing faster in dev
		bCryptCost = 3
	}

	switch cfg.Method {
	case auth.MethodPassword:
		return auth.NewPasswordAuthenticator(authCfg, bCryptCost)
	case auth.MethodGithub:
		return auth.NewGithubAuthenticator(authCfg)
	case auth.MethodNone:
		return auth.NewNoneAuthenticator(authCfg)
	}

	return nil, fmt.Errorf("unsupported auth method '%s'", cfg.Method)
}
