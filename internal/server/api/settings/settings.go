package settings

import (
	"fmt"
	"gitlab.com/texm/shokku/internal/env"
	"gitlab.com/texm/shokku/internal/models"
	"gitlab.com/texm/shokku/internal/server/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetVersions(e *env.Env, c echo.Context) error {
	dokkuVersion, err := e.Dokku.GetDokkuVersion()
	if err != nil {
		return fmt.Errorf("getting dokku version: %w", err)
	}
	return c.JSON(http.StatusOK, &dto.GetVersionsResponse{
		Dokku:  dokkuVersion,
		Shokku: e.Version,
	})
}

func GetUsers(e *env.Env, c echo.Context) error {
	var dbUsers []models.User
	res := e.DB.Model(models.User{}).Preload("SSHKeys").Find(&dbUsers)
	if res.Error != nil {
		return fmt.Errorf("querying db users: %w", res.Error)
	}

	users := make([]dto.User, len(dbUsers))
	for i, dbUser := range dbUsers {
		keys := make([]string, len(dbUser.SSHKeys))
		for j, key := range dbUser.SSHKeys {
			keys[j] = key.Key
		}
		users[i] = dto.User{
			Name:    dbUser.Name,
			Source:  dbUser.Source,
			SSHKeys: keys,
		}
	}

	return c.JSON(http.StatusOK, dto.GetUsersResponse{
		Users: users,
	})
}

func GetSSHKeys(e *env.Env, c echo.Context) error {
	keys, err := e.Dokku.ListSSHKeys()
	if err != nil {
		return fmt.Errorf("listing ssh keys: %w", err)
	}
	return c.JSON(http.StatusOK, &dto.GetSSHKeysResponse{
		Keys: keys,
	})
}

func GetGlobalDomains(e *env.Env, c echo.Context) error {
	report, err := e.Dokku.GetGlobalDomainsReport()
	if err != nil {
		return fmt.Errorf("getting global domains report: %w", err)
	}

	if len(report.Domains) == 0 {
		report.Domains = make([]string, 0)
	}

	return c.JSON(http.StatusOK, &dto.GetGlobalDomainsResponse{
		Domains: report.Domains,
		Enabled: report.Enabled,
	})
}

func AddGlobalDomain(e *env.Env, c echo.Context) error {
	var req dto.AlterGlobalDomainRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	// TODO: domain verification etc
	if err := e.Dokku.AddGlobalDomain(req.Domain); err != nil {
		return fmt.Errorf("adding global domain: %w", err)
	}

	return c.NoContent(http.StatusOK)
}

func RemoveGlobalDomain(e *env.Env, c echo.Context) error {
	var req dto.DeleteGlobalDomainRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	// TODO: domain verification etc
	if err := e.Dokku.RemoveGlobalDomain(req.Domain); err != nil {
		return fmt.Errorf("removing global domain: %w", err)
	}

	return c.NoContent(http.StatusOK)
}

func SetEventLoggingEnabled(e *env.Env, c echo.Context) error {
	var req dto.SetEventLoggingEnabledRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	if err := e.Dokku.SetEventLoggingEnabled(req.Enabled); err != nil {
		return fmt.Errorf("setting event logging: %w", err)
	}

	return c.NoContent(http.StatusOK)
}

func GetEventLogsList(e *env.Env, c echo.Context) error {
	events, err := e.Dokku.ListLoggedEvents()
	if err != nil {
		return fmt.Errorf("removing global domain: %w", err)
	}

	return c.JSON(http.StatusOK, dto.GetEventLogsListResponse{
		Events: events,
	})
}

func GetEventLogs(e *env.Env, c echo.Context) error {
	logs, err := e.Dokku.GetEventLogs()
	if err != nil {
		return fmt.Errorf("getting event logs: %w", err)
	}

	return c.JSON(http.StatusOK, dto.GetEventLogsResponse{
		Logs: logs,
	})
}

func ListPlugins(e *env.Env, c echo.Context) error {
	plugins, err := e.Dokku.ListPlugins()
	if err != nil {
		return fmt.Errorf("listing plugins: %w", err)
	}
	info := make([]dto.PluginInfo, len(plugins))
	for i := 0; i < len(plugins); i++ {
		p := plugins[i]
		info[i] = dto.PluginInfo{
			Name:        p.Name,
			Version:     p.Version,
			Enabled:     p.Enabled,
			Description: p.Description,
		}
	}
	return c.JSON(http.StatusOK, &dto.ListPluginsResponse{
		Plugins: info,
	})
}
