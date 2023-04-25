package apps

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/texm/dokku-go"
	"gitlab.com/texm/shokku/internal/env"
	"gitlab.com/texm/shokku/internal/server/dto"
	"net/http"
)

func GetAppBuilder(e *env.Env, c echo.Context) error {
	var req dto.GetAppBuilderRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	report, err := e.Dokku.GetAppBuilderReport(req.Name)
	if err != nil {
		return fmt.Errorf("getting app builder: %w", err)
	}

	selectedBuilder := report.ComputedSelectedBuilder
	if selectedBuilder == "" {
		selectedBuilder = "auto"
	}

	return c.JSON(http.StatusOK, dto.GetAppBuilderResponse{
		Selected: selectedBuilder,
	})
}

func SetAppBuilder(e *env.Env, c echo.Context) error {
	var req dto.SetAppBuilderRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	builders := map[string]dokku.AppBuilder{
		"auto":       "",
		"dockerfile": dokku.AppBuilderDockerfile,
		"herokuish":  dokku.AppBuilderHerokuish,
		"null":       dokku.AppBuilderNull,
		"buildpack":  dokku.AppBuilderPack,
		"lambda":     dokku.AppBuilderLambda,
	}
	chosenBuilder, supported := builders[req.Builder]
	if !supported {
		return echo.NewHTTPError(http.StatusBadRequest,
			fmt.Sprintf("unsupported builder '%s'", req.Builder))
	}

	err := e.Dokku.SetAppSelectedBuilder(req.Name, chosenBuilder)
	if err != nil {
		return fmt.Errorf("setting app builder: %w", err)
	}

	return c.NoContent(http.StatusOK)
}

func GetAppBuildDirectory(e *env.Env, c echo.Context) error {
	var req dto.GetAppBuildDirectoryRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	report, err := e.Dokku.GetAppBuilderReport(req.Name)
	if err != nil {
		return fmt.Errorf("getting app build dir: %w", err)
	}

	return c.JSON(http.StatusOK, dto.GetAppBuildDirectoryResponse{
		Directory: report.ComputedBuildDir,
	})
}

func SetAppBuildDirectory(e *env.Env, c echo.Context) error {
	var req dto.SetAppBuildDirectoryRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	err := e.Dokku.SetAppBuilderProperty(req.Name, dokku.BuilderPropertyBuildDir, req.Directory)
	if err != nil {
		return fmt.Errorf("setting app build dir: %w", err)
	}

	return c.NoContent(http.StatusOK)
}

func ClearAppBuildDirectory(e *env.Env, c echo.Context) error {
	var req dto.ClearAppBuildDirectoryRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	err := e.Dokku.SetAppBuilderProperty(req.Name, dokku.BuilderPropertyBuildDir, "")
	if err != nil {
		return fmt.Errorf("clearing app build dir: %w", err)
	}

	return c.NoContent(http.StatusOK)
}
