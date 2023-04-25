package services

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

func GetServiceBackupReport(e *env.Env, c echo.Context) error {
	var req dto.GetServiceBackupReportRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	dbSvc, err := lookupDBServiceByName(e, req.Name)
	if err != nil {
		return echo.ErrNotFound
	}

	cmd := fmt.Sprintf("%s:backup-schedule-cat %s", dbSvc.Type, req.Name)
	backupSchedule, err := e.Dokku.Exec(cmd)
	if err != nil {
		backupSchedule = ""
	}

	report := dto.ServiceBackupReport{
		AuthSet:       dbSvc.BackupAuthSet,
		EncryptionSet: dbSvc.BackupEncryptionSet,
		Bucket:        dbSvc.BackupBucket,
		Schedule:      backupSchedule,
	}

	return c.JSON(http.StatusOK, dto.GetServiceBackupReportResponse{
		Report: report,
	})
}

func SetServiceBackupAuth(e *env.Env, c echo.Context) error {
	var req dto.SetServiceBackupsAuthRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	dbSvc, err := lookupDBServiceByName(e, req.Name)
	if err != nil {
		return echo.ErrNotFound
	}

	cfg := req.Config
	args := fmt.Sprintf("%s %s %s %s %s", cfg.AccessKeyId, cfg.SecretKey,
		cfg.Region, cfg.SignatureVersion, cfg.EndpointUrl)
	cmd := fmt.Sprintf("%s:backup-auth %s %s", dbSvc.Type, req.Name, args)
	if _, execErr := e.Dokku.Exec(cmd); execErr != nil {
		return fmt.Errorf("setting backup auth: %w", execErr)
	}

	dbSvc.BackupAuthSet = true
	if err := e.DB.Save(&dbSvc).Error; err != nil {
		log.Error().Err(err).Msg("error updating service backup auth")
		return echo.ErrInternalServerError
	}

	return c.NoContent(http.StatusOK)
}

func SetServiceBackupBucket(e *env.Env, c echo.Context) error {
	var req dto.SetServiceBackupsBucketRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	dbSvc, err := lookupDBServiceByName(e, req.Name)
	if err != nil {
		return echo.ErrNotFound
	}

	dbSvc.BackupBucket = req.Bucket
	if dbErr := e.DB.Save(&dbSvc).Error; dbErr != nil {
		log.Error().Err(dbErr).Msg("error updating service backup bucket")
		return echo.ErrInternalServerError
	}

	return c.NoContent(http.StatusOK)
}

func RunServiceBackup(e *env.Env, c echo.Context) error {
	var req dto.RunServiceBackupRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	dbSvc, err := lookupDBServiceByName(e, req.Name)
	if err != nil {
		return echo.ErrNotFound
	}
	if dbSvc.BackupBucket == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "backup bucket not set")
	} else if !dbSvc.BackupAuthSet {
		return echo.NewHTTPError(http.StatusBadRequest, "backup auth not set")
	}

	dokkuCmd := fmt.Sprintf("%s:backup %s %s", dbSvc.Type, dbSvc.Name, dbSvc.BackupBucket)

	cmd := func() (*dokku.CommandOutputStream, error) {
		return e.Dokku.ExecStreaming(dokkuCmd)
	}

	return c.JSON(http.StatusOK, dto.CommandExecutionResponse{
		ExecutionID: commands.RequestExecution(cmd, nil),
	})
}

func SetServiceBackupSchedule(e *env.Env, c echo.Context) error {
	var req dto.SetServiceBackupsScheduleRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	dbSvc, err := lookupDBServiceByName(e, req.Name)
	if err != nil {
		return echo.ErrNotFound
	}
	if dbSvc.BackupBucket == "" {
		return echo.NewHTTPError(http.StatusBadRequest,
			"service backup bucket not set")
	}

	cmd := fmt.Sprintf(`%s:backup-schedule %s "%s" %s`, dbSvc.Type,
		req.Name, req.Schedule, dbSvc.BackupBucket)
	if out, err := e.Dokku.Exec(cmd); err != nil {
		log.Debug().Str("output", out).Msg("backup schedule output")
		return fmt.Errorf("setting backup schedule: %w", err)
	}

	return c.NoContent(http.StatusOK)
}

func RemoveServiceBackupSchedule(e *env.Env, c echo.Context) error {
	var req dto.ManageServiceRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	cmd := fmt.Sprintf(`%s:backup-unschedule %s`, req.Type, req.Name)
	if out, err := e.Dokku.Exec(cmd); err != nil {
		log.Debug().Str("output", out).Msg("backup schedule output")
		return fmt.Errorf("removing backup schedule: %w", err)
	}

	return c.NoContent(http.StatusOK)
}

func SetServiceBackupEncryption(e *env.Env, c echo.Context) error {
	var req dto.SetServiceBackupsEncryptionRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	dbSvc, err := lookupDBServiceByName(e, req.Name)
	if err != nil {
		return echo.ErrNotFound
	}

	cmd := fmt.Sprintf(`%s:backup-set-encryption %s %s`, dbSvc.Type,
		req.Name, req.Passphrase)
	if out, err := e.Dokku.Exec(cmd); err != nil {
		log.Debug().Str("output", out).Msg("set backup encryption output")
		return fmt.Errorf("setting backup encryption: %w", err)
	}

	dbSvc.BackupEncryptionSet = true
	if saveErr := e.DB.Save(&dbSvc).Error; saveErr != nil {
		log.Error().Err(saveErr).Msg("error updating service backup encryption")
		return echo.ErrInternalServerError
	}

	return c.NoContent(http.StatusOK)
}

func RemoveServiceBackupEncryption(e *env.Env, c echo.Context) error {
	var req dto.ManageServiceRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	dbSvc, err := lookupDBServiceByName(e, req.Name)
	if err != nil {
		return echo.ErrNotFound
	}

	cmd := fmt.Sprintf(`%s:backup-unset-encryption %s`, req.Type, req.Name)
	if out, err := e.Dokku.Exec(cmd); err != nil {
		log.Debug().Str("output", out).Msg("unset backup encryption output")
		return fmt.Errorf("removing backup encryption: %w", err)
	}

	dbSvc.BackupEncryptionSet = false
	if saveErr := e.DB.Save(&dbSvc).Error; saveErr != nil {
		log.Error().Err(saveErr).Msg("error updating service backup encryption")
		return echo.ErrInternalServerError
	}

	return c.NoContent(http.StatusOK)
}
