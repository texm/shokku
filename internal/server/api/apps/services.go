package apps

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gitlab.com/texm/shokku/internal/env"
	"gitlab.com/texm/shokku/internal/server/dto"
	"net/http"
	"strings"
)

var (
	dokkuErrPrefix = "!    "
	serviceTypes   = []string{"redis", "postgres", "mysql", "mongo"}
)

func splitDokkuListOutput(output string) ([]string, error) {
	if strings.HasPrefix(output, dokkuErrPrefix) {
		return nil, nil
	}
	if output == "" {
		return []string{}, nil
	}
	return strings.Split(output, "\n"), nil
}

func getAppServiceLinks(e *env.Env, appName string, serviceType string) ([]string, error) {
	linksCmd := fmt.Sprintf("%s:app-links %s --quiet", serviceType, appName)
	out, err := e.Dokku.Exec(linksCmd)
	if err != nil {
		return nil, err
	}
	return splitDokkuListOutput(out)
}

func GetAppServices(e *env.Env, c echo.Context) error {
	var req dto.GetAppServicesRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	serviceList := []dto.ServiceInfo{}

	for _, serviceType := range serviceTypes {
		links, err := getAppServiceLinks(e, req.Name, serviceType)
		if err != nil {
			return fmt.Errorf("getting links for " + serviceType)
		}
		for _, name := range links {
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
