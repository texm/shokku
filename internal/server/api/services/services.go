package services

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"github.com/texm/dokku-go"
	"gitlab.com/texm/shokku/internal/env"
	"gitlab.com/texm/shokku/internal/models"
	"gitlab.com/texm/shokku/internal/server/commands"
	"gitlab.com/texm/shokku/internal/server/dto"
	"net/http"
	"strings"
)

var (
	dokkuErrPrefix = "!    "
	serviceTypes   = []string{"redis", "postgres", "mysql", "mongo"}
)

func lookupDBServiceByName(e *env.Env, name string) (*models.Service, error) {
	dbSvc := models.Service{
		Name: name,
	}
	res := e.DB.Where("name = ?", name).Find(&dbSvc)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("no service found for %s", name)
	}
	return &dbSvc, nil
}

func splitDokkuListOutput(output string) ([]string, error) {
	if strings.HasPrefix(output, dokkuErrPrefix) {
		return nil, nil
	}
	if output == "" {
		return []string{}, nil
	}
	return strings.Split(output, "\n"), nil
}

func getServiceAppLinks(e *env.Env, serviceName string, serviceType string) ([]string, error) {
	linksCmd := fmt.Sprintf("%s:links %s --quiet", serviceType, serviceName)
	out, err := e.Dokku.Exec(linksCmd)
	if err != nil {
		return nil, err
	}
	return splitDokkuListOutput(out)
}

func getServiceList(e *env.Env, serviceType string) ([]string, error) {
	listCmd := fmt.Sprintf("%s:list --quiet", serviceType)
	out, err := e.Dokku.Exec(listCmd)
	if err != nil {
		return nil, err
	}
	if strings.Contains(out, "There are no") {
		return []string{}, nil
	}
	return splitDokkuListOutput(out)
}

func ListServices(e *env.Env, c echo.Context) error {
	serviceList := []dto.ServiceInfo{}

	for _, serviceType := range serviceTypes {
		services, err := getServiceList(e, serviceType)
		if err != nil {
			return fmt.Errorf("getting list for %s services: %w", serviceType, err)
		}
		for _, name := range services {
			serviceList = append(serviceList, dto.ServiceInfo{
				Name: name,
				Type: serviceType,
			})
		}
	}

	return c.JSON(http.StatusOK, dto.ListServicesResponse{
		Services: serviceList,
	})
}

func GetServiceType(e *env.Env, c echo.Context) error {
	var req dto.GetServiceTypeRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	dbSvc, err := lookupDBServiceByName(e, req.Name)
	if err != nil {
		return echo.ErrNotFound
	}

	return c.JSON(http.StatusOK, dto.GetServiceTypeResponse{
		Type: dbSvc.Type,
	})
}

func GetServiceInfo(e *env.Env, c echo.Context) error {
	var req dto.GenericServiceRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	info := map[string]string{
		"version": "", "internal-ip": "", "status": "", // "dsn": "",
	}
	for key := range info {
		cmd := fmt.Sprintf("%s:info %s --%s", req.Type, req.Name, key)
		out, err := e.Dokku.Exec(cmd)
		if err != nil {
			return fmt.Errorf("getting service info: %w", err)
		}
		info[key] = out
	}

	return c.JSON(http.StatusOK, dto.GetServiceInfoResponse{
		Info: info,
	})
}

func getGenericLinkFlags(req dto.LinkGenericServiceToAppRequest) string {
	var flags []string
	if req.Alias != "" {
		flag := fmt.Sprintf("--alias %s", req.Alias)
		flags = append(flags, flag)
	}
	if req.QueryString != "" {
		flag := fmt.Sprintf("--querystring %s", req.QueryString)
		flags = append(flags, flag)
	}
	return strings.Join(flags, " ")
}

func LinkGenericServiceToApp(e *env.Env, c echo.Context) error {
	var req dto.LinkGenericServiceToAppRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	dbSvc, err := lookupDBServiceByName(e, req.ServiceName)
	if err != nil {
		return echo.ErrNotFound
	}

	flags := getGenericLinkFlags(req)
	linkCmd := fmt.Sprintf("%s:link %s %s %s",
		dbSvc.Type, dbSvc.Name, req.AppName, flags)

	cmd := func() (*dokku.CommandOutputStream, error) {
		return e.Dokku.ExecStreaming(linkCmd)
	}

	return c.JSON(http.StatusOK, dto.CommandExecutionResponse{
		ExecutionID: commands.RequestExecution(cmd, nil),
	})
}

func UnlinkGenericServiceFromApp(e *env.Env, c echo.Context) error {
	var req dto.LinkGenericServiceToAppRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	dbSvc, err := lookupDBServiceByName(e, req.ServiceName)
	if err != nil {
		return echo.ErrNotFound
	}

	unlinkCmd := fmt.Sprintf("%s:unlink %s %s", dbSvc.Type,
		dbSvc.Name, req.AppName)

	cmd := func() (*dokku.CommandOutputStream, error) {
		return e.Dokku.ExecStreaming(unlinkCmd)
	}

	return c.JSON(http.StatusOK, dto.CommandExecutionResponse{
		ExecutionID: commands.RequestExecution(cmd, nil),
	})
}

func GetServiceLinkedApps(e *env.Env, c echo.Context) error {
	var req dto.GenericServiceRequest
	if err := dto.BindRequest(c, &req); err != nil {
		log.Error().Err(err.ToHTTP()).Msg("error")
		return err.ToHTTP()
	}

	apps, err := getServiceAppLinks(e, req.Name, req.Type)
	if err != nil {
		return fmt.Errorf("getting linked apps: %w", err)
	}

	return c.JSON(http.StatusOK, dto.GetServiceLinkedAppsResponse{
		Apps: apps,
	})
}

func GetServiceLogs(e *env.Env, c echo.Context) error {
	var req dto.GenericServiceRequest
	if err := dto.BindRequest(c, &req); err != nil {
		log.Error().Err(err.ToHTTP()).Msg("error")
		return err.ToHTTP()
	}

	cmd := fmt.Sprintf("%s:logs %s", req.Type, req.Name)
	out, err := e.Dokku.Exec(cmd)
	if err != nil {
		return fmt.Errorf("getting linked apps: %w", err)
	}

	logs := strings.Split(out, "\n")
	return c.JSON(http.StatusOK, dto.GetServiceLogsResponse{
		Logs: logs,
	})
}
