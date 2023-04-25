package settings

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/texm/dokku-go"
	"gitlab.com/texm/shokku/internal/env"
	"gitlab.com/texm/shokku/internal/server/dto"
	"net/http"
)

func SetDockerRegistry(e *env.Env, c echo.Context) error {
	var req dto.SetDockerRegistryRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	if err := e.Dokku.LoginDockerRegistry(req.Server, req.Username, req.Password); err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err.Error())
	}

	propErr := e.Dokku.SetAppDockerRegistryProperty("--global", dokku.DockerRegistryPropertyServer, req.Server)
	if propErr != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, propErr.Error())
	}

	return c.NoContent(http.StatusOK)
}

func GetDockerRegistryReport(e *env.Env, c echo.Context) error {
	report, err := e.Dokku.GetDockerRegistryReport()
	if err != nil && !errors.Is(err, dokku.NoDeployedAppsError) {
		return fmt.Errorf("failed to get registry report: %w", err)
	}

	var response dto.GetDockerRegistryReportResponse
	for _, appReport := range report {
		response.Server = appReport.GlobalServer
		response.PushOnRelease = appReport.GlobalPushOnRelease
		break
	}

	return c.JSON(http.StatusOK, response)
}

func AddGitAuth(e *env.Env, c echo.Context) error {
	var req dto.AddGitAuthRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	if err := e.Dokku.GitSetAuth(req.Host, req.Username, req.Password); err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err.Error())
	}

	if err := e.Dokku.GitAllowHost(req.Host); err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func RemoveGitAuth(e *env.Env, c echo.Context) error {
	var req dto.RemoveGitAuthRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	if err := e.Dokku.GitRemoveAuth(req.Host); err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
