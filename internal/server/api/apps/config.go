package apps

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gitlab.com/texm/shokku/internal/env"
	"gitlab.com/texm/shokku/internal/server/dto"
	"net/http"
)

func GetAppConfig(e *env.Env, c echo.Context) error {
	var req dto.GetAppConfigRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	config, err := e.Dokku.GetAppConfig(req.Name)
	if err != nil {
		return fmt.Errorf("getting app config: %w", err)
	}

	return c.JSON(http.StatusOK, dto.GetAppConfigResponse{
		Config: config,
	})
}

func SetAppConfig(e *env.Env, c echo.Context) error {
	var req dto.SetAppConfigRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	if err := e.Dokku.SetAppConfigValues(req.Name, req.Config, false); err != nil {
		return fmt.Errorf("setting app config: %w", err)
	}

	return c.NoContent(http.StatusOK)
}
