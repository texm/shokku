package dokku

import (
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/texm/dokku-go"
	"gitlab.com/texm/shokku/internal/env"
	"gitlab.com/texm/shokku/internal/models"
	"strings"
)

// todo: deduplicate from api/services
var (
	dokkuErrPrefix = "!    "
	serviceTypes   = []string{"redis", "postgres", "mysql", "mongo"}
)

const FilteredApp = "shokku"

func SyncState(e *env.Env) {
	syncApps(e)
	syncServices(e)
}

func getServiceList(e *env.Env, serviceType string) ([]string, error) {
	listCmd := fmt.Sprintf("%s:list --quiet", serviceType)
	out, err := e.Dokku.Exec(listCmd)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(out, dokkuErrPrefix) {
		return []string{}, nil
	}
	return strings.Split(out, "\n"), nil
}

func syncApps(e *env.Env) {
	logger := log.With().Str("dokku_sync", "apps").Logger()

	apps, listErr := e.Dokku.ListApps()
	if listErr != nil {
		if !errors.Is(listErr, dokku.NoDeployedAppsError) {
			logger.Error().
				Err(listErr).
				Msg("Failed to get dokku apps")
			return
		}
	}

	var filtered []string
	for _, name := range apps {
		if name != FilteredApp {
			filtered = append(filtered, name)
		}
	}

	var dbApps []models.App
	if dbErr := e.DB.Find(&dbApps).Error; dbErr != nil {
		logger.Error().
			Err(dbErr).
			Msg("Failed to query db apps")
		return
	}

	appMap := map[string]bool{}
	for _, name := range filtered {
		var dbApp models.App
		res := e.DB.Limit(1).
			Where(&models.App{Name: name}).
			FirstOrCreate(&dbApp)
		if res.Error != nil {
			logger.Error().
				Err(res.Error).
				Str("app_name", name).
				Msg("failed to create db app")
		}

		appMap[name] = true
	}

	for _, dbApp := range dbApps {
		status, found := appMap[dbApp.Name]
		if !found || !status {
			if err := e.DB.Delete(&dbApp).Error; err != nil {
				logger.Error().
					Err(err).
					Uint("id", dbApp.ID).
					Str("name", dbApp.Name).
					Msg("failed to clean old app")
			}
			continue
		}

		if syncErr := syncApp(e, dbApp); syncErr != nil {
			logger.Error().
				Err(syncErr).
				Str("name", dbApp.Name).
				Msg("failed to sync dokku app")
		}
	}
}

func syncApp(e *env.Env, dbApp models.App) error {
	// TODO: sync setup info
	return nil
}

func serviceKey(name, svcType string) string {
	return fmt.Sprintf("%s_%s", svcType, name)
}

func syncServices(e *env.Env) {
	logger := log.With().Str("dokku_sync", "services").Logger()

	var dbServices []models.Service
	if dbErr := e.DB.Find(&dbServices).Error; dbErr != nil {
		logger.Error().
			Err(dbErr).
			Msg("Failed to query db services")
		return
	}

	svcMap := map[string]bool{}
	for _, serviceType := range serviceTypes {
		services, err := getServiceList(e, serviceType)
		if err != nil {
			logger.Error().
				Err(err).
				Msgf("failed getting %s services", serviceType)
			continue
		}
		for _, name := range services {
			svcMap[serviceKey(name, serviceType)] = true
			var dbSvc models.Service
			e.DB.Where(&models.Service{Name: name, Type: serviceType}).
				Limit(1).
				FirstOrCreate(&dbSvc)

			e.DB.Save(&dbSvc)
		}
	}

	for _, svc := range dbServices {
		key := serviceKey(svc.Name, svc.Type)
		status, found := svcMap[key]
		if !found || !status {
			if err := e.DB.Delete(&svc).Error; err != nil {
				logger.Error().
					Err(err).
					Uint("id", svc.ID).
					Str("type", svc.Type).
					Str("name", svc.Name).
					Msg("failed to delete old service")
			}
		}
	}
}
