package services

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"github.com/texm/dokku-go"
	"gitlab.com/texm/shokku/internal/env"
	"gitlab.com/texm/shokku/internal/models"
	"gitlab.com/texm/shokku/internal/server/dto"
	"net/http"
	"strings"
)

var (
	ErrDestroyingLinkedService = echo.NewHTTPError(http.StatusBadRequest,
		"cannot destroy a linked service")
	ErrServiceNameTaken = echo.NewHTTPError(http.StatusBadRequest,
		"service name exists")
)

func manageService(e *env.Env, c echo.Context, mgmt string, flags string) error {
	var req dto.GenericServiceRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	cmd := fmt.Sprintf("%s:%s %s %s", req.Type, mgmt, req.Name, flags)
	if _, err := e.Dokku.Exec(cmd); err != nil {
		return fmt.Errorf(mgmt+"ing service", err)
	}

	return c.NoContent(http.StatusOK)
}

func StartService(e *env.Env, c echo.Context) error {
	return manageService(e, c, "start", "")
}
func StopService(e *env.Env, c echo.Context) error {
	return manageService(e, c, "stop", "")
}
func RestartService(e *env.Env, c echo.Context) error {
	return manageService(e, c, "restart", "")
}

func DestroyService(e *env.Env, c echo.Context) error {
	var req dto.GenericServiceRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	dbSvc, lookupErr := lookupDBServiceByName(e, req.Name)
	if lookupErr != nil {
		return echo.ErrNotFound
	}

	cmd := fmt.Sprintf("%s:destroy %s -f", req.Type, req.Name)
	if _, err := e.Dokku.Exec(cmd); err != nil {
		var dokkuErr *dokku.ExitCodeError
		if errors.As(err, &dokkuErr) {
			if strings.HasSuffix(dokkuErr.Output(), "Cannot delete linked service") {
				return ErrDestroyingLinkedService
			}
		}
		return fmt.Errorf("destroying service: %w", err)
	}

	if err := e.DB.Delete(&dbSvc).Error; err != nil {
		log.Error().Err(err).Interface("svc", dbSvc).Msg("error deleting service")
		return echo.ErrInternalServerError
	}

	return c.NoContent(http.StatusOK)
}

// use short opt strings otherwise arg parse is funky
func maybeAppendStringOptionFlag(flags *[]string, opt string, cfgOption *string) {
	if cfgOption != nil {
		*flags = append(*flags, fmt.Sprintf("-%s '%s'", opt, *cfgOption))
	}
}

func getServiceCreateFlags(req dto.GenericServiceCreationConfig) string {
	flags := &[]string{}

	if req.CustomEnv != nil {
		envVars := make([]string, len(*req.CustomEnv))
		for i, e := range *req.CustomEnv {
			envVars[i] = fmt.Sprintf("%s=%s", e[0], e[1])
		}
		envStr := strings.Join(envVars, ";")
		*flags = append(*flags, fmt.Sprintf("-C \"%s\"", envStr))
	}

	maybeAppendStringOptionFlag(flags, "c", req.ConfigOptions)
	maybeAppendStringOptionFlag(flags, "i", req.Image)
	maybeAppendStringOptionFlag(flags, "m", req.MemoryLimit)
	maybeAppendStringOptionFlag(flags, "p", req.Password)
	maybeAppendStringOptionFlag(flags, "r", req.RootPassword)
	maybeAppendStringOptionFlag(flags, "s", req.SharedMemorySize)

	return strings.Join(*flags, " ")
}

func CreateNewGenericService(e *env.Env, c echo.Context) error {
	var req dto.CreateGenericServiceRequest
	if err := dto.BindRequest(c, &req); err != nil {
		log.Debug().Err(err.ToHTTP()).Interface("req", req).Msgf("req failed")
		return err.ToHTTP()
	}

	dbSvc := models.Service{
		Name: req.Name,
		Type: req.ServiceType,
	}
	if e.DB.Where("name = ?", req.Name).Find(&dbSvc).RowsAffected != 0 {
		return ErrServiceNameTaken
	}

	e.DB.Save(&dbSvc)

	flags := getServiceCreateFlags(req.Config)
	cmd := fmt.Sprintf("%s:create %s %s", req.ServiceType, req.Name, flags)
	_, err := e.Dokku.Exec(cmd)
	if err != nil {
		return fmt.Errorf("creating %s service: %w", req.ServiceType, err)
	}

	return c.NoContent(http.StatusOK)
}

func CloneService(e *env.Env, c echo.Context) error {
	var req dto.CloneServiceRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	dbSvc, notFoundErr := lookupDBServiceByName(e, req.Name)
	if notFoundErr != nil {
		return echo.NewHTTPError(http.StatusNotFound, "service not found")
	}

	_, notFoundErr = lookupDBServiceByName(e, req.NewName)
	if notFoundErr == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "new service name exists")
	}

	cmd := fmt.Sprintf("%s:clone %s %s", dbSvc.Type, req.Name, req.NewName)
	_, err := e.Dokku.Exec(cmd)
	if err != nil {
		return fmt.Errorf("cloning %s service: %w", dbSvc.Type, err)
	}

	newSvc := models.Service{
		Name: req.NewName,
		Type: dbSvc.Type,
	}
	if err := e.DB.Save(&newSvc).Error; err != nil {
		log.Error().Err(err).
			Str("name", req.NewName).Str("type", dbSvc.Type).
			Msg("failed to save cloned service to db")
		return echo.ErrInternalServerError
	}

	return c.NoContent(http.StatusOK)
}
