package github

import (
	"context"
	"net/http"

	"github.com/bradleyfalzon/ghinstallation/v2"
	gh "github.com/google/go-github/v48/github"
	"github.com/rs/zerolog/log"

	"gitlab.com/texm/shokku/internal/env"
	"gitlab.com/texm/shokku/internal/models"
)

type AppClient struct {
	*gh.Client
	appsTransport *ghinstallation.AppsTransport
}

func GetAppClient(e *env.Env) (*AppClient, error) {
	var dbApp models.GithubApp
	if err := e.DB.Find(&dbApp).Error; err != nil {
		return nil, err
	}
	transport, err := ghinstallation.NewAppsTransport(
		http.DefaultTransport, dbApp.AppId, []byte(dbApp.PEM))
	if err != nil {
		log.Debug().Err(err).Msg("failed to create transport")
		return nil, err
	}
	appClient := &AppClient{
		Client:        gh.NewClient(&http.Client{Transport: transport}),
		appsTransport: transport,
	}
	return appClient, nil
}

type AppInstallationClient struct {
	*gh.Client
}

func (c *AppClient) GetInstallationClient(id int64) *AppInstallationClient {
	transport := ghinstallation.NewFromAppsTransport(c.appsTransport, id)
	client := gh.NewClient(&http.Client{Transport: transport})
	return &AppInstallationClient{Client: client}
}

func CompleteAppManifest(e *env.Env, code string) (*gh.AppConfig, error) {
	ctx := context.Background()
	client := gh.NewClient(nil)
	cfg, _, ghErr := client.Apps.CompleteAppManifest(ctx, code)
	if ghErr != nil {
		return nil, ghErr
	}

	appId := cfg.GetID()
	dbApp := models.GithubApp{AppId: appId}
	if dbErr := e.DB.FirstOrCreate(&dbApp).Error; dbErr != nil {
		log.Error().Err(dbErr).Msg("failed db lookup")
		return nil, dbErr
	}

	dbApp.AppId = appId
	dbApp.ClientId = cfg.GetClientID()
	dbApp.NodeId = cfg.GetNodeID()
	dbApp.Slug = cfg.GetSlug()
	dbApp.PEM = cfg.GetPEM()
	dbApp.ClientSecret = cfg.GetClientSecret()
	dbApp.WebhookSecret = cfg.GetWebhookSecret()

	// saveRes := e.DB.Where(&models.GithubApp{AppId: appId}).Save(&dbApp)
	if err := e.DB.Save(&dbApp).Error; err != nil {
		log.Error().Err(err).Msg("failed to save db app")
		return nil, err
	}

	return cfg, nil
}

func (c *AppClient) GetApp(ctx context.Context) (*gh.App, error) {
	app, _, err := c.Apps.Get(ctx, "")
	if err != nil {
		return nil, err
	}
	return app, nil
}
