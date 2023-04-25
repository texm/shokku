package env

import (
	"github.com/labstack/echo/v4"
	"github.com/texm/dokku-go"
	"gitlab.com/texm/shokku/internal/server/auth"
	"gorm.io/gorm"
)

const Version = "0.0.1"

type Env struct {
	Version string
	Router  *echo.Echo
	Dokku   *dokku.SSHClient
	DB      *gorm.DB
	Auth    auth.Authenticator

	DebugMode      bool
	SetupCompleted bool
}

func New(debugMode bool) *Env {
	return &Env{
		DebugMode: debugMode,
		Version:   Version,
	}
}

type handlerFunc func(env *Env, c echo.Context) error

func (e *Env) H(f handlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error { return f(e, c) }
}

func (e *Env) Shutdown() error {
	if err := e.Dokku.Close(); err != nil {
		return err
	}
	return nil
}
