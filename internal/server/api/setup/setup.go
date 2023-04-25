package setup

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"gitlab.com/texm/shokku/internal/env"
	"gitlab.com/texm/shokku/internal/models"
	"gitlab.com/texm/shokku/internal/server/auth"
	"gitlab.com/texm/shokku/internal/server/dto"
	"net/http"
)

func GetStatus(e *env.Env, c echo.Context) error {
	method := string(e.Auth.GetMethod())
	log.Debug().
		Bool("is_setup", e.SetupCompleted).
		Str("method", method).
		Msg("get setup status")
	return c.JSON(http.StatusOK, dto.GetSetupStatusResponse{
		IsSetup: e.SetupCompleted,
		Method:  method,
	})
}

func GetSetupKeyValid(e *env.Env, c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func setupServerWithAuthMethod(e *env.Env, method auth.Method) error {
	var state models.Server
	e.DB.FirstOrCreate(&state)

	state.IsSetup = true
	state.AuthMethod = method
	if err := e.DB.Save(&state).Error; err != nil {
		log.Error().Err(err).Msg("failed to save setup state")
		return echo.ErrInternalServerError
	}

	newAuth, authErr := createAuthenticator(e, method)
	if authErr != nil {
		log.Error().Err(authErr).Msg("failed to init new authenticator")
		return echo.ErrInternalServerError
	}
	e.Auth = newAuth
	e.SetupCompleted = true

	return nil
}

func createAuthenticator(e *env.Env, method auth.Method) (auth.Authenticator, error) {
	config := auth.Config{
		SigningKey:    e.Auth.GetSigningKey(),
		CookieDomain:  e.Auth.GetCookieDomain(),
		TokenLifetime: e.Auth.GetTokenLifetime(),
	}
	switch method {
	case auth.MethodGithub:
		return auth.NewGithubAuthenticator(config)
	case auth.MethodPassword:
		return auth.NewPasswordAuthenticator(config, auth.DefaultBCryptCost)
	}
	return nil, fmt.Errorf("unknown method %s", method)
}
