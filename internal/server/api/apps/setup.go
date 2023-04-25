package apps

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"github.com/texm/dokku-go"
	"gitlab.com/texm/shokku/internal/env"
	"gitlab.com/texm/shokku/internal/models"
	"gitlab.com/texm/shokku/internal/server/auth"
	"gitlab.com/texm/shokku/internal/server/commands"
	"gitlab.com/texm/shokku/internal/server/dto"
	"gorm.io/gorm/clause"
	"net/http"
	"net/url"
)

type setupMethod string

const (
	methodGit      = setupMethod("git")
	methodSyncRepo = setupMethod("sync-repo")
	methodDocker   = setupMethod("docker")
	methodArchive  = setupMethod("archive")
)

var (
	updateOnConflict = clause.OnConflict{
		DoUpdates: clause.AssignmentColumns([]string{"updated_at"}),
	}
)

func saveAppSetupMethod(e *env.Env, app *models.App, method setupMethod, cfg *models.AppSetupConfig) error {
	app.IsSetup = true
	app.SetupMethod = string(method)
	if appErr := e.DB.Save(app).Error; appErr != nil {
		log.Error().Err(appErr).
			Str("app", app.Name).
			Str("method", app.SetupMethod).
			Msg("failed to save app table")
		return fmt.Errorf("failed to save app table: %w", appErr)
	}

	cfg.AppID = app.ID

	updateErr := e.DB.Clauses(updateOnConflict).Create(&cfg).Error
	if updateErr != nil {
		log.Error().Err(updateErr).
			Str("app", app.Name).
			Interface("cfg", cfg).
			Msg("failed to update app setup config")
		return fmt.Errorf("failed to save app setup config: %w", updateErr)
	}

	return nil
}

func GetAppSetupStatus(e *env.Env, c echo.Context) error {
	var req dto.GetAppSetupStatusRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	dbApp, appErr := lookupDBAppByName(e, req.Name)
	if appErr != nil {
		return echo.ErrNotFound
	}

	return c.JSON(http.StatusOK, dto.GetAppSetupStatusResponse{
		IsSetup: dbApp.IsSetup,
		Method:  dbApp.SetupMethod,
	})
}

func GetAppSetupConfig(e *env.Env, c echo.Context) error {
	var req dto.GetAppSetupConfigRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	dbApp, appErr := lookupDBAppByName(e, req.Name)
	if appErr != nil {
		return echo.ErrNotFound
	}

	var cfg models.AppSetupConfig
	e.DB.Where("app_id = ?", dbApp.ID).FirstOrInit(&cfg)

	return c.JSON(http.StatusOK, dto.GetAppSetupConfigResponse{
		IsSetup:          dbApp.IsSetup,
		Method:           dbApp.SetupMethod,
		DeploymentBranch: cfg.DeployBranch,
		RepoURL:          cfg.RepoURL,
		RepoGitRef:       cfg.RepoGitRef,
		Image:            cfg.Image,
	})
}

func SetupAppNewRepo(e *env.Env, c echo.Context) error {
	var req dto.SetupAppNewRepoRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	dbApp, appErr := lookupDBAppByName(e, req.Name)
	if appErr != nil {
		return echo.ErrNotFound
	}

	if initErr := e.Dokku.GitInitializeApp(req.Name); initErr != nil {
		return fmt.Errorf("initialising git repo: %w", initErr)
	}

	branch := req.DeploymentBranch
	if branch == "" {
		branch = "master"
	}

	if branch != "master" {
		branchErr := e.Dokku.GitSetAppProperty(req.Name, dokku.GitPropertyDeployBranch, req.DeploymentBranch)
		if branchErr != nil {
			return fmt.Errorf("setting git deploy branch: %w", branchErr)
		}
	}

	cfg := &models.AppSetupConfig{
		DeployBranch: branch,
	}
	if saveErr := saveAppSetupMethod(e, dbApp, methodGit, cfg); saveErr != nil {
		return saveErr
	}

	return c.NoContent(http.StatusOK)
}

func SetupAppSyncRepo(e *env.Env, c echo.Context) error {
	var req dto.SetupAppSyncRepoRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	dbApp, appErr := lookupDBAppByName(e, req.Name)
	if appErr != nil {
		return echo.ErrNotFound
	}

	parsedURL, urlErr := url.Parse(req.RepositoryURL)
	if urlErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid repo url")
	}

	if e.Auth.GetMethod() == auth.MethodGithub {
		if parsedURL.Hostname() != "github.com" {
			//	return echo.ErrBadRequest
		}
	}

	syncOpts := &dokku.GitSyncOptions{
		Build:  true,
		GitRef: req.GitRef,
	}
	execFn := func() (*dokku.CommandOutputStream, error) {
		return e.Dokku.GitSyncAppRepo(req.Name, req.RepositoryURL, syncOpts)
	}

	cfg := &models.AppSetupConfig{
		AppID:      dbApp.ID,
		RepoURL:    req.RepositoryURL,
		RepoGitRef: req.GitRef,
	}
	cb := func() error {
		return saveAppSetupMethod(e, dbApp, methodSyncRepo, cfg)
	}

	return c.JSON(http.StatusOK, dto.CommandExecutionResponse{
		ExecutionID: commands.RequestExecution(execFn, cb),
	})
}

func SetupAppPullImage(e *env.Env, c echo.Context) error {
	var req dto.SetupAppPullImageRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	dbApp, appErr := lookupDBAppByName(e, req.Name)
	if appErr != nil {
		return echo.ErrNotFound
	}

	execFn := func() (*dokku.CommandOutputStream, error) {
		return e.Dokku.GitCreateFromImage(req.Name, req.Image, nil)
	}

	cfg := &models.AppSetupConfig{
		AppID: dbApp.ID,
		Image: req.Image,
	}
	cb := func() error {
		return saveAppSetupMethod(e, dbApp, methodDocker, cfg)
	}
	return c.JSON(http.StatusOK, dto.CommandExecutionResponse{
		ExecutionID: commands.RequestExecution(execFn, cb),
	})
}

func SetupAppUploadArchive(e *env.Env, c echo.Context) error {
	var req dto.SetupAppUploadArchiveRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	// disable for now
	if true {
		return echo.ErrForbidden
	}

	file, err := c.FormFile("archive")
	if err != nil {
		log.Error().Err(err).Msg("invalid form file")
		return echo.ErrInternalServerError
	}

	log.Info().
		Str("app", req.Name).
		Msgf("got file: %+v", file.Header)

	return c.NoContent(http.StatusNotImplemented)
}
