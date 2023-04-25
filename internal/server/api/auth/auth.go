package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"gitlab.com/texm/shokku/internal/env"
	"gitlab.com/texm/shokku/internal/server/auth"
	"net/http"
)

func HandleLogout(e *env.Env, c echo.Context) error {
	e.Auth.ClearTokenCookies(c)
	return c.NoContent(http.StatusOK)
}

func HandleRefreshAuth(e *env.Env, c echo.Context) error {
	user, err := e.Auth.GetUserFromContext(c)
	if err != nil {
		log.Error().Msg("failed to parse user from context")
		return echo.ErrInternalServerError
	}

	claims := auth.UserClaims{
		Name: user.Name,
	}
	token, err := e.Auth.NewToken(claims)
	if err != nil {
		log.Error().Err(err).Msg("failed to create jwt")
		return echo.ErrInternalServerError
	}

	e.Auth.SetTokenCookies(c, token)

	return c.NoContent(http.StatusOK)
}

func HandleGetDetails(e *env.Env, c echo.Context) error {
	user, err := e.Auth.GetUserFromContext(c)
	if err != nil {
		log.Error().Msg("failed to parse user from context")
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, echo.Map{
		"username": user.Name,
	})
}
