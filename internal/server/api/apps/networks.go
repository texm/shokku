package apps

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/texm/dokku-go"
	"gitlab.com/texm/shokku/internal/env"
	"gitlab.com/texm/shokku/internal/server/dto"
	"net/http"
)

func GetAppNetworksReport(e *env.Env, c echo.Context) error {
	var req dto.GetAppNetworksReportRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	report, err := e.Dokku.GetAppNetworkReport(req.Name)
	if err != nil {
		return fmt.Errorf("getting app network report: %w", err)
	}

	return c.JSON(http.StatusOK, dto.GetAppNetworksReportResponse{
		AttachInitial:     report.ComputedInitialNetwork,
		AttachPostCreate:  report.ComputedAttachPostCreate,
		AttachPostDeploy:  report.ComputedAttachPostDeploy,
		BindAllInterfaces: report.ComputedBindAllInterfaces,
		TLD:               report.ComputedTLD,
		WebListeners:      report.WebListeners,
	})
}

func SetAppNetworks(e *env.Env, c echo.Context) error {
	var req dto.SetAppNetworksRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	if req.Initial != nil {
		err := e.Dokku.SetAppNetworkProperty(req.Name,
			dokku.NetworkPropertyInitialNetwork, *req.Initial)
		if err != nil {
			return fmt.Errorf("setting initial network: %w", err)
		}
	}

	if req.PostCreate != nil {
		err := e.Dokku.SetAppNetworkProperty(req.Name,
			dokku.NetworkPropertyAttachPostCreate, *req.PostCreate)
		if err != nil {
			return fmt.Errorf("setting postcreate network: %w", err)
		}
	}

	if req.PostDeploy != nil {
		err := e.Dokku.SetAppNetworkProperty(req.Name,
			dokku.NetworkPropertyAttachPostDeploy, *req.PostDeploy)
		if err != nil {
			return fmt.Errorf("setting postdeploy network: %w", err)
		}
	}

	if req.BindAllInterfaces != nil {
		bindAll := "false"
		if *req.BindAllInterfaces {
			bindAll = "true"
		}
		err := e.Dokku.SetAppNetworkProperty(req.Name,
			dokku.NetworkPropertyBindAllInterfaces, bindAll)
		if err != nil {
			return fmt.Errorf("setting bind-all-interfaces: %w", err)
		}
	}

	if req.TLD != nil {
		err := e.Dokku.SetAppNetworkProperty(req.Name,
			dokku.NetworkPropertyTLD, *req.TLD)
		if err != nil {
			return fmt.Errorf("setting TLD: %w", err)
		}
	}

	return c.NoContent(http.StatusOK)
}
