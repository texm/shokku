package settings

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gitlab.com/texm/shokku/internal/env"
	"gitlab.com/texm/shokku/internal/server/dto"
	"net/http"
)

func CreateNetwork(e *env.Env, c echo.Context) error {
	var req dto.AlterNetworkRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	if err := e.Dokku.CreateNetwork(req.Network); err != nil {
		return fmt.Errorf("error creating network: %w", err)
	}

	return c.NoContent(http.StatusOK)
}

func DestroyNetwork(e *env.Env, c echo.Context) error {
	var req dto.AlterNetworkRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	if err := e.Dokku.DestroyNetwork(req.Network); err != nil {
		return fmt.Errorf("error destroying network: %w", err)
	}

	return c.NoContent(http.StatusOK)
}

func ListNetworks(e *env.Env, c echo.Context) error {
	networks, err := e.Dokku.ListNetworks()
	if err != nil {
		return fmt.Errorf("error listing networks: %w", err)
	}

	return c.JSON(http.StatusOK, dto.ListNetworksResponse{
		Networks: networks,
	})
}
