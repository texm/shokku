package api

import (
	"github.com/google/go-github/v48/github"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"gitlab.com/texm/shokku/internal/env"
	"gitlab.com/texm/shokku/internal/models"
	"net/http"
)

func ReceiveGithubWebhookEvent(e *env.Env, c echo.Context) error {
	var ghApp models.GithubApp
	if ghErr := e.DB.First(&ghApp).Error; ghErr != nil {
		log.Error().Err(ghErr).Msg("failed to retrieve github app info")
		return echo.ErrInternalServerError
	}

	req := c.Request()
	secret := []byte(ghApp.WebhookSecret)
	payload, validationErr := github.ValidatePayload(req, secret)
	if validationErr != nil {
		return echo.ErrBadRequest
	}
	event, parseErr := github.ParseWebHook(github.WebHookType(req), payload)
	if parseErr != nil {
		log.Error().Err(parseErr).Msg("failed to parse webhook")
		return echo.ErrInternalServerError
	}
	var err error
	switch event := event.(type) {
	case *github.MetaEvent:
		err = processMetaEvent(e, event)
	case *github.PushEvent:
		err = processPushEvent(e, event)
	default:
		log.Error().
			Interface("type", event).
			Msg("received unsupported webhook event")
		return echo.ErrInternalServerError
	}

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusOK)
}

func processMetaEvent(e *env.Env, event *github.MetaEvent) error {
	log.Debug().Interface("event", event).Msg("got meta event")

	return nil
}

func processPushEvent(e *env.Env, event *github.PushEvent) error {
	return nil
}
