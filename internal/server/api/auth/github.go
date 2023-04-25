package auth

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"gitlab.com/texm/shokku/internal/env"
	"gitlab.com/texm/shokku/internal/models"
	"gitlab.com/texm/shokku/internal/server/auth"
	"gitlab.com/texm/shokku/internal/server/dto"
	"gitlab.com/texm/shokku/internal/server/github"
	"net/http"
)

func GetGithubAuthInfo(e *env.Env, c echo.Context) error {
	var ghApp models.GithubApp
	if err := e.DB.First(&ghApp).Error; err != nil {
		log.Error().Err(err).Msg("no github app in db")
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, dto.GetGithubAuthInfoResponse{
		ClientID: ghApp.ClientId,
	})
}

func CompleteGithubAuth(e *env.Env, c echo.Context) error {
	var req dto.GithubAuthRequest
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}

	ctx := context.Background()
	params := github.CodeExchangeParams{
		Code:        req.Code,
		Scopes:      []string{},
		RedirectURL: req.RedirectURL,
	}
	client, err := github.ExchangeCode(ctx, e, params)
	if err != nil {
		log.Error().Err(err).Msg("failed to exchange code for client")
		return echo.ErrBadRequest
	}

	user, err := client.GetUser(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to get user")
		return echo.ErrBadRequest
	}

	var count int64
	dbUser := models.User{Name: user.GetLogin()}
	r := e.DB.Model(&dbUser).Where(&dbUser).Count(&count)
	if r.Error != nil {
		log.Error().
			Err(r.Error).
			Str("name", user.GetLogin()).
			Msg("user lookup error")
		return echo.ErrInternalServerError
	}
	if count == 0 {
		return echo.ErrForbidden
	}

	claims := auth.UserClaims{
		Name: user.GetLogin(),
	}
	token, err := e.Auth.NewToken(claims)
	if err != nil {
		log.Error().Err(err).Msg("failed to create jwt")
		return echo.ErrInternalServerError
	}
	e.Auth.SetTokenCookies(c, token)

	return c.NoContent(http.StatusOK)
}
