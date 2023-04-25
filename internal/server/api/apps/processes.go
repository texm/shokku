package apps

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"github.com/texm/dokku-go"
	"gitlab.com/texm/shokku/internal/env"
	"gitlab.com/texm/shokku/internal/server/commands"
	"gitlab.com/texm/shokku/internal/server/dto"
	"net/http"
)

func GetAppProcesses(e *env.Env, c echo.Context) error {
	var req dto.GetAppProcessesRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	scales, err := e.Dokku.GetAppProcessScale(req.Name)
	if err != nil {
		return fmt.Errorf("getting app process scale: %w", err)
	}

	processes := make([]string, len(scales))
	i := 0
	for processName := range scales {
		processes[i] = processName
		i++
	}

	return c.JSON(http.StatusOK, dto.GetAppProcessesResponse{
		Processes: processes,
	})
}

func GetAppProcessReport(e *env.Env, c echo.Context) error {
	var req dto.GetAppProcessReportRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	resourceReport, err := e.Dokku.GetAppResourceReport(req.Name)
	if err != nil {
		return fmt.Errorf("getting app resource report: %w", err)
	}

	processScale, err := e.Dokku.GetAppProcessScale(req.Name)
	if err != nil {
		return fmt.Errorf("getting app process scale: %w", err)
	}

	processMap := map[string]dto.AppProcessInfo{}
	for processName, scale := range processScale {
		appResources := dto.AppProcessInfo{
			Scale: scale,
		}
		if psSettings, ok := resourceReport.Processes[processName]; ok {
			appResources.Resources = psSettings
		}
		processMap[processName] = appResources
	}

	return c.JSON(http.StatusOK, dto.GetAppProcessReportResponse{
		ResourceDefaults: resourceReport.Defaults,
		Processes:        processMap,
	})
}

func SetAppProcessResources(e *env.Env, c echo.Context) error {
	var req dto.SetAppProcessResourcesRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	limits := req.ResourceLimits
	reservations := req.ResourceReservations

	if limits.CPU != nil {
		err := e.Dokku.SetAppProcessResourceLimit(req.Name, req.Process,
			dokku.ResourceCPU, *limits.CPU)
		if err != nil {
			return fmt.Errorf("setting app cpu limit: %w", err)
		}
	} else {
		err := e.Dokku.ClearAppProcessResourceLimit(req.Name, req.Process, dokku.ResourceCPU)
		if err != nil {
			return fmt.Errorf("clearing app cpu limit: %w", err)
		}
	}

	if limits.Memory != nil && limits.MemoryUnit != nil {
		resSpec := dokku.ResourceSpec{
			Name:   "memory",
			Suffix: *limits.MemoryUnit,
		}
		err := e.Dokku.SetAppProcessResourceLimit(req.Name, req.Process, resSpec, *limits.Memory)
		if err != nil {
			return fmt.Errorf("setting app mem limit: %w", err)
		}
	} else {
		err := e.Dokku.ClearAppProcessResourceLimit(req.Name, req.Process, dokku.ResourceMemoryBytes)
		if err != nil {
			return fmt.Errorf("clearing app mem limit: %w", err)
		}
	}

	if reservations.Memory != nil && reservations.MemoryUnit != nil {
		resSpec := dokku.ResourceSpec{
			Name:   "memory",
			Suffix: *reservations.MemoryUnit,
		}
		err := e.Dokku.SetAppProcessResourceReservation(req.Name, req.Process, resSpec,
			*reservations.Memory)
		if err != nil {
			return fmt.Errorf("setting app mem reservation: %w", err)
		}
	} else {
		err := e.Dokku.ClearAppProcessResourceReservation(req.Name, req.Process,
			dokku.ResourceMemoryBytes)
		if err != nil {
			return fmt.Errorf("clearing app mem reservation: %w", err)
		}
	}

	return c.NoContent(http.StatusOK)
}

func GetAppProcessScale(e *env.Env, c echo.Context) error {
	var req dto.GetAppProcessScaleRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	scale, err := e.Dokku.GetAppProcessScale(req.Name)
	if err != nil {
		return fmt.Errorf("getting app process scale: %w", err)
	}

	return c.JSON(http.StatusOK, dto.GetAppProcessScaleResponse{
		ProcessScale: scale,
	})
}

func SetAppProcessDeployChecks(e *env.Env, c echo.Context) error {
	var req dto.SetAppProcessDeployChecksRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	processes := []string{req.Process}
	var err error
	switch req.State {
	case "enabled":
		err = e.Dokku.EnableAppProcessesDeployChecks(req.Name, processes)
	case "disabled":
		err = e.Dokku.DisableAppProcessesDeployChecks(req.Name, processes)
	case "skipped":
		err = e.Dokku.SetAppProcessesDeployChecksSkipped(req.Name, processes)
	default:
		return echo.NewHTTPError(http.StatusBadRequest, "unknown state")
	}

	if err != nil {
		return fmt.Errorf("setting app deploy checks to " + req.State)
	}

	return c.NoContent(http.StatusOK)
}

func SetAppProcessScale(e *env.Env, c echo.Context) error {
	var req dto.SetAppProcessScaleRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	_, err := e.Auth.GetUserFromContext(c)
	if err != nil {
		log.Error().Err(err).Msg("failed to retrieve user from context")
		return echo.ErrInternalServerError
	}

	cmd := func() (*dokku.CommandOutputStream, error) {
		return e.Dokku.SetAppProcessScale(req.Name, req.Process, req.Scale, req.SkipDeploy)
	}

	return c.JSON(http.StatusOK, dto.CommandExecutionResponse{
		ExecutionID: commands.RequestExecution(cmd, nil),
	})
}
