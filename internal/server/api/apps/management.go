package apps

import (
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"gitlab.com/texm/shokku/internal/env"
	"gitlab.com/texm/shokku/internal/models"
	"gitlab.com/texm/shokku/internal/server/commands"
	"gitlab.com/texm/shokku/internal/server/dto"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/texm/dokku-go"
)

const FilteredApp = "shokku"

func lookupDBAppByName(e *env.Env, name string) (*models.App, error) {
	dbApp := models.App{Name: name}
	res := e.DB.Where("name = ?", name).Find(&dbApp)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("no app found for %s", name)
	}
	return &dbApp, nil
}

func GetAppsList(e *env.Env, c echo.Context) error {
	appsList, err := e.Dokku.ListApps()

	apps := make([]dto.GetAppsListItem, 0)
	for _, name := range appsList {
		if name != FilteredApp {
			apps = append(apps, dto.GetAppsListItem{
				Name: name,
				// TODO: get type
				Type: "",
			})
		}
	}

	if err != nil && !errors.Is(err, dokku.NoDeployedAppsError) {
		return fmt.Errorf("getting apps overview: %w", err)
	}

	return c.JSON(http.StatusOK, dto.GetAppsListResponse{
		Apps: apps,
	})
}

func GetAppsProcessReport(e *env.Env, c echo.Context) error {
	allReports, err := e.Dokku.GetAllProcessReport()
	if err != nil {
		return fmt.Errorf("failed to get apps report: %w", err)
	}

	apps := make([]dto.GetAppOverviewResponse, 0)
	for name, psReport := range allReports {
		if name == FilteredApp {
			continue
		}

		app, lookupErr := lookupDBAppByName(e, name)
		if lookupErr != nil {
			return fmt.Errorf("failed to lookup app %s: %w", name, lookupErr)
		}

		apps = append(apps, dto.GetAppOverviewResponse{
			Name:         name,
			IsSetup:      app.IsSetup,
			SetupMethod:  app.SetupMethod,
			IsDeployed:   psReport.Deployed,
			IsRunning:    psReport.Running,
			NumProcesses: psReport.Processes,
			CanScale:     psReport.CanScale,
			Restore:      psReport.Restore,
		})
	}

	return c.JSON(http.StatusOK, dto.GetAllAppsOverviewResponse{
		Apps: apps,
	})
}

func GetAppOverview(e *env.Env, c echo.Context) error {
	var req dto.GetAppOverviewRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	app, lookupErr := lookupDBAppByName(e, req.Name)
	if lookupErr != nil {
		return fmt.Errorf("failed to lookup app %s: %w", req.Name, lookupErr)
	}

	psReport, psErr := e.Dokku.GetAppProcessReport(req.Name)
	if psErr != nil {
		return fmt.Errorf("getting apps process report: %w", psErr)
	}

	gitReport, gitErr := e.Dokku.GitGetAppReport(req.Name)
	if gitErr != nil {
		return fmt.Errorf("getting app git report: %w", gitErr)
	}

	return c.JSON(http.StatusOK, dto.GetAppOverviewResponse{
		IsSetup:         app.IsSetup,
		SetupMethod:     app.SetupMethod,
		IsDeployed:      psReport.Deployed,
		GitDeployBranch: gitReport.DeployBranch,
		GitLastUpdated:  gitReport.LastUpdatedAt,
		IsRunning:       psReport.Running,
	})
}

func GetAppInfo(e *env.Env, c echo.Context) error {
	var req dto.GetAppInfoRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	info, err := e.Dokku.GetAppReport(req.Name)
	if err != nil {
		return fmt.Errorf("getting app info: %w", err)
	}

	res := &dto.GetAppInfoResponse{
		Info: dto.AppInfo{
			Name:                 req.Name,
			Directory:            info.Directory,
			DeploySource:         info.DeploySource,
			DeploySourceMetadata: info.DeploySourceMetadata,
			CreatedAt:            time.UnixMilli(info.CreatedAtTimestamp),
			IsLocked:             info.IsLocked,
		},
	}

	return c.JSON(http.StatusOK, res)
}

func CreateApp(e *env.Env, c echo.Context) error {
	var req dto.ManageAppRequest
	if reqErr := dto.BindRequest(c, &req); reqErr != nil {
		log.Debug().
			Err(reqErr.ToHTTP()).
			Str("appName", req.Name).
			Msg("bind err")
		return reqErr.ToHTTP()
	}

	_, lookupErr := lookupDBAppByName(e, req.Name)
	if lookupErr == nil {
		return echo.ErrBadRequest
	}

	if createErr := e.Dokku.CreateApp(req.Name); createErr != nil {
		return fmt.Errorf("creating app: %w", createErr)
	}

	if dbErr := e.DB.Create(&models.App{Name: req.Name}).Error; dbErr != nil {
		log.Error().Err(dbErr).Str("name", req.Name).Msg("failed to create db app")
		return echo.ErrInternalServerError
	}

	return c.NoContent(http.StatusOK)
}

func DestroyApp(e *env.Env, c echo.Context) error {
	var req dto.DestroyAppRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	dbApp, dbErr := lookupDBAppByName(e, req.Name)
	if dbErr != nil {
		log.Error().Err(dbErr).Str("name", req.Name).Msg("failed to lookup app")
		return echo.ErrNotFound
	}

	if err := e.Dokku.DestroyApp(req.Name); err != nil {
		return fmt.Errorf("destroying app: %w", err)
	}

	// TODO: hard delete app
	if err := e.DB.Delete(&dbApp).Error; err != nil {
		log.Error().Err(err).Str("name", req.Name).Msg("failed to delete app")
		return echo.ErrInternalServerError
	}

	return c.NoContent(http.StatusOK)
}

func RenameApp(e *env.Env, c echo.Context) error {
	var req dto.RenameAppRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	dbApp, dbErr := lookupDBAppByName(e, req.CurrentName)
	if dbErr != nil {
		log.Error().Err(dbErr).Str("name", req.CurrentName).
			Msg("failed to lookup app")
		return echo.ErrNotFound
	}

	if _, newDbErr := lookupDBAppByName(e, req.NewName); newDbErr == nil {
		return echo.ErrBadRequest
	}

	dbApp.Name = req.NewName
	if saveErr := e.DB.Save(&dbApp).Error; saveErr != nil {
		log.Error().Err(saveErr).
			Str("name", req.NewName).
			Msg("failed to save db app")
	}

	options := &dokku.AppManagementOptions{SkipDeploy: true}
	if renameErr := e.Dokku.RenameApp(req.CurrentName, req.NewName, options); renameErr != nil {
		return fmt.Errorf("renaming app: %w", renameErr)
	}

	return c.NoContent(http.StatusOK)
}

func StartApp(e *env.Env, c echo.Context) error {
	var req dto.ManageAppRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	cmd := func() (*dokku.CommandOutputStream, error) {
		return e.Dokku.StartApp(req.Name, nil)
	}

	return c.JSON(http.StatusOK, dto.CommandExecutionResponse{
		ExecutionID: commands.RequestExecution(cmd, nil),
	})
}

func StopApp(e *env.Env, c echo.Context) error {
	var req dto.ManageAppRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	cmd := func() (*dokku.CommandOutputStream, error) {
		return e.Dokku.StopApp(req.Name, nil)
	}

	return c.JSON(http.StatusOK, dto.CommandExecutionResponse{
		ExecutionID: commands.RequestExecution(cmd, nil),
	})
}

func RestartApp(e *env.Env, c echo.Context) error {
	var req dto.ManageAppRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	cmd := func() (*dokku.CommandOutputStream, error) {
		return e.Dokku.RestartApp(req.Name, nil)
	}

	return c.JSON(http.StatusOK, dto.CommandExecutionResponse{
		ExecutionID: commands.RequestExecution(cmd, nil),
	})
}

func RebuildApp(e *env.Env, c echo.Context) error {
	var req dto.ManageAppRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	dbApp, appErr := lookupDBAppByName(e, req.Name)
	if appErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "couldnt get app")
	}
	if !dbApp.IsSetup {
		return echo.NewHTTPError(http.StatusBadRequest, "not setup")
	}

	var cfg models.AppSetupConfig
	cfgErr := e.DB.Where("app_id = ?", dbApp.ID).First(&cfg).Error
	if cfgErr != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "no setup config")
	}
	method := setupMethod(dbApp.SetupMethod)
	var cmd commands.AsyncDokkuCommand
	if method == methodGit {
		cmd = func() (*dokku.CommandOutputStream, error) {
			return e.Dokku.RebuildApp(req.Name, nil)
		}
	} else if method == methodSyncRepo {
		opts := &dokku.GitSyncOptions{
			Build:  true,
			GitRef: cfg.RepoGitRef,
		}
		cmd = func() (*dokku.CommandOutputStream, error) {
			return e.Dokku.GitSyncAppRepo(req.Name, cfg.RepoURL, opts)
		}
	} else if method == methodDocker {
		image := cfg.Image
		cmd = func() (*dokku.CommandOutputStream, error) {
			return e.Dokku.GitCreateFromImage(req.Name, image, nil)
		}
	} else {
		log.Error().
			Str("method", dbApp.SetupMethod).
			Str("name", dbApp.Name).
			Msg("invalid app setup method")
		return echo.NewHTTPError(http.StatusBadRequest, "invalid setup method")
	}

	return c.JSON(http.StatusOK, dto.CommandExecutionResponse{
		ExecutionID: commands.RequestExecution(cmd, nil),
	})
}

func GetAppDeployChecks(e *env.Env, c echo.Context) error {
	var req dto.GetAppDeployChecksRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	report, err := e.Dokku.GetAppDeployChecksReport(req.Name)
	if err != nil {
		return fmt.Errorf("getting app deploy checks: %w", err)
	}

	return c.JSON(http.StatusOK, dto.GetAppDeployChecksResponse{
		AllDisabled:       report.AllDisabled,
		AllSkipped:        report.AllSkipped,
		DisabledProcesses: report.DisabledProcesses,
		SkippedProcesses:  report.SkippedProcesses,
	})
}

func SetAppDeployChecks(e *env.Env, c echo.Context) error {
	var req dto.SetAppDeployChecksRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	var err error
	switch req.State {
	case "enabled":
		err = e.Dokku.EnableAppDeployChecks(req.Name)
	case "disabled":
		err = e.Dokku.DisableAppDeployChecks(req.Name)
	case "skipped":
		err = e.Dokku.SetAppDeployChecksSkipped(req.Name)
	default:
		return echo.NewHTTPError(http.StatusBadRequest, "unknown state")
	}

	if err != nil {
		return fmt.Errorf("setting app deploy checks to %s: %w", req.State, err)
	}

	return c.NoContent(http.StatusOK)
}

func GetAppLogs(e *env.Env, c echo.Context) error {
	var req dto.GetAppLogsRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	logs, err := e.Dokku.GetAppLogs(req.Name)
	if err != nil {
		if errors.Is(err, dokku.AppNotDeployedError) {
			return c.JSON(http.StatusOK, dto.GetAppLogsResponse{})
		}
		return fmt.Errorf("getting app logs: %w", err)
	}

	return c.JSON(http.StatusOK, dto.GetAppLogsResponse{
		Logs: strings.Split(logs, "\n"),
	})
}

/*
func AppExecInProcess(e *env.Env, c echo.Context) error {
	var req dto.AppExecInProcessRequest
	if err := dto.BindRequest(c, &req); err != nil {
		log.Debug().Str("err", err.String()).Interface("req", req).Msg("bind")
		return err.ToHTTP()
	}

	cmd := fmt.Sprintf(`enter %s %s %s`, req.AppName, req.ProcessName, req.Command)

	output, execErr := e.Dokku.Exec(cmd)

	res := dto.AppExecInProcessResponse{
		Output: output,
	}
	if execErr != nil {
		res.Error = execErr.Error()

		var sshExitErr *dokku.ExitCodeError
		if errors.As(execErr, &sshExitErr) {
			res.Error = sshExitErr.Output()
		}
	}

	return c.JSON(http.StatusOK, res)
}
*/
