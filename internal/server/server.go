package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gitlab.com/texm/shokku/internal/env"
	"gitlab.com/texm/shokku/internal/server/api"
	"gitlab.com/texm/shokku/internal/server/commands"
	"gitlab.com/texm/shokku/internal/server/db"
	"gitlab.com/texm/shokku/internal/server/dokku"
	"gitlab.com/texm/shokku/internal/server/middleware"
	"io/fs"
	"time"
)

func New(cfg Config, appFS fs.FS) (*env.Env, error) {
	e := env.New(cfg.DebugMode)
	e.Router = initRouter(cfg.DebugMode, appFS)

	dbSess, dbErr := db.Init(cfg.DBPath)
	if dbErr != nil {
		return nil, fmt.Errorf("failed to init db: %w", dbErr)
	}
	e.DB = dbSess

	serverSecrets, secretsErr := getServerSecrets(e.DB)
	if secretsErr != nil {
		return nil, fmt.Errorf("failed to get server secrets: %w", secretsErr)
	}

	dokkuCfg := dokku.Config{
		DebugMode:  cfg.DebugMode,
		PrivateKey: serverSecrets.privKey,
		Host:       cfg.DokkuSSHHost,
		Port:       cfg.DokkuSSHPort,
	}
	dokkuClient, dokkuErr := dokku.Init(dokkuCfg)
	if dokkuErr != nil {
		return nil, fmt.Errorf("failed to init dokku client: %w", dokkuErr)
	}
	e.Dokku = dokkuClient

	serverState, stateErr := initServerState(e)
	if stateErr != nil {
		return nil, fmt.Errorf("failed to get setup state: %w", stateErr)
	}

	authCfg := initAuthConfig{
		SigningKey:    serverSecrets.signingKey,
		TokenLifetime: time.Minute * time.Duration(cfg.AuthTokenLifetimeMinutes),
		Method:        serverState.AuthMethod,
		DebugMode:     cfg.DebugMode,
		IsSetup:       serverState.IsSetup,
	}
	authn, authErr := initAuthenticator(authCfg)
	if authErr != nil {
		return nil, fmt.Errorf("failed to init authenticator: %w", authErr)
	}
	e.Auth = authn

	if !e.SetupCompleted {
		e.Router.Use(middleware.ServerSetupBlocker(e, serverState.SetupKey))
	}

	api.RegisterRoutes(e, []echo.MiddlewareFunc{
		middleware.TokenAuth(e),
		middleware.ProvideUserContext(e),
	})

	dokku.SyncState(e)

	go commands.PollStatuses()

	return e, nil
}
