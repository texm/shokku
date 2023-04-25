package setup

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"gitlab.com/texm/shokku/internal/env"
	"gitlab.com/texm/shokku/internal/models"
	"gitlab.com/texm/shokku/internal/server/auth"
	"gitlab.com/texm/shokku/internal/server/dto"
	"gitlab.com/texm/shokku/internal/server/github"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

const (
	githubAppInstallURL = "https://github.com/apps/%s/installations/new/permissions?target_id=%d"
)

func GetGithubSetupStatus(e *env.Env, c echo.Context) error {
	var ghApp models.GithubApp
	r := e.DB.Find(&ghApp)
	if r.Error != nil && !errors.Is(r.Error, sql.ErrNoRows) {
		log.Error().Err(r.Error).Msg("Failed to lookup github app")
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, dto.GetGithubSetupStatus{
		AppCreated: r.RowsAffected > 0,
	})
}

func CreateGithubApp(e *env.Env, c echo.Context) error {
	var req dto.CreateGithubAppRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	cfg, err := github.CompleteAppManifest(e, req.Code)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, dto.CreateGithubAppResponse{
		Slug: cfg.GetSlug(),
	})
}

func GetGithubAppInstallInfo(e *env.Env, c echo.Context) error {
	ctx := context.Background()

	client, clientErr := github.GetAppClient(e)
	if clientErr != nil {
		log.Error().Err(clientErr).Msg("failed to get app client")
		return echo.NewHTTPError(http.StatusBadRequest, clientErr)
	}

	app, appErr := client.GetApp(ctx)
	if appErr != nil {
		log.Error().Err(appErr).Msg("failed to get github client app")
		return echo.NewHTTPError(http.StatusBadRequest, appErr)
	}
	url := fmt.Sprintf(githubAppInstallURL, app.GetSlug(), app.Owner.GetID())

	return c.JSON(http.StatusOK, dto.InstallGithubAppResponse{
		InstallURL: url,
	})
}

func CompleteGithubSetup(e *env.Env, c echo.Context) error {
	var req dto.CompleteGithubSetupRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	client, clientErr := github.GetAppClient(e)
	if clientErr != nil {
		log.Error().Err(clientErr).Msg("failed to get app client")
		return echo.ErrBadRequest
	}

	ctx := context.Background()
	inst, _, instErr := client.Apps.GetInstallation(ctx, req.InstallationId)
	if instErr != nil {
		log.Error().
			Err(instErr).
			Int64("id", req.InstallationId).
			Msg("failed to get installation")
		return echo.ErrBadRequest
	}

	if err := setupServerWithAuthMethod(e, auth.MethodGithub); err != nil {
		log.Error().Err(err).Msg("failed to setup github auth")
		return echo.ErrInternalServerError
	}

	go func() {
		if err := github.SyncUsersToDB(e); err != nil {
			log.Error().Err(err).Msg("failed to sync github users")
		}
	}()

	log.Debug().
		Int64("id", inst.GetID()).
		Msg("installed github app")

	return c.NoContent(http.StatusOK)
}
