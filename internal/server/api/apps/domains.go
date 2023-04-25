package apps

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gitlab.com/texm/shokku/internal/env"
	"gitlab.com/texm/shokku/internal/server/dto"
	"net/http"
)

func GetAppDomainsReport(e *env.Env, c echo.Context) error {
	var req dto.GetAppDomainsReportRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return echo.ErrBadRequest
	}

	report, err := e.Dokku.GetAppDomainsReport(req.Name)
	if err != nil {
		return fmt.Errorf("getting app domains report: %w", err)
	}

	if len(report.AppDomains) == 0 {
		report.AppDomains = make([]string, 0)
	}

	if len(report.GlobalDomains) == 0 {
		report.GlobalDomains = make([]string, 0)
	}

	return c.JSON(http.StatusOK, dto.GetAppDomainsReportResponse{
		Domains: report.AppDomains,
		Enabled: report.AppEnabled,
	})
}

func SetAppDomainsEnabled(e *env.Env, c echo.Context) error {
	var req dto.SetAppDomainsEnabledRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return echo.ErrBadRequest
	}

	var err error
	if req.Enabled {
		err = e.Dokku.EnableAppDomains(req.Name)
	} else {
		err = e.Dokku.DisableAppDomains(req.Name)
	}
	if err != nil {
		return fmt.Errorf("setting app domains enabled: %w", err)
	}

	return c.NoContent(http.StatusOK)
}

func GetAppLetsEncryptEnabled(e *env.Env, c echo.Context) error {
	var req dto.GetAppLetsEncryptEnabledRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return echo.ErrBadRequest
	}

	return c.NoContent(http.StatusNotImplemented)
}

func SetAppLetsEncryptEnabled(e *env.Env, c echo.Context) error {
	var req dto.SetAppLetsEncryptEnabledRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return echo.ErrBadRequest
	}

	return c.NoContent(http.StatusNotImplemented)
}

func AddAppDomain(e *env.Env, c echo.Context) error {
	var req dto.AlterAppDomainRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return echo.ErrBadRequest
	}

	if err := e.Dokku.AddAppDomain(req.Name, req.Domain); err != nil {
		return fmt.Errorf("adding app domain: %w", err)
	}

	return c.NoContent(http.StatusOK)
}

func RemoveAppDomain(e *env.Env, c echo.Context) error {
	var req dto.AlterAppDomainRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	if err := e.Dokku.RemoveAppDomain(req.Name, req.Domain); err != nil {
		return fmt.Errorf("removing app domain: %w", err)
	}

	return c.NoContent(http.StatusOK)
}
