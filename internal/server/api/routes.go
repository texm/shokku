package api

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/texm/shokku/internal/env"
	"gitlab.com/texm/shokku/internal/server/api/apps"
	"gitlab.com/texm/shokku/internal/server/api/auth"
	"gitlab.com/texm/shokku/internal/server/api/services"
	"gitlab.com/texm/shokku/internal/server/api/settings"
	"gitlab.com/texm/shokku/internal/server/api/setup"
)

func RegisterRoutes(e *env.Env, authMiddleware []echo.MiddlewareFunc) {
	g := e.Router.Group("/api")

	setup.RegisterRoutes(e, g.Group("/setup"))
	auth.RegisterRoutes(e, g.Group("/auth"), g.Group("/auth", authMiddleware...))

	protectedApi := g.Group("", authMiddleware...)
	apps.RegisterRoutes(e, protectedApi.Group("/apps"))
	services.RegisterRoutes(e, protectedApi.Group("/services"))
	settings.RegisterRoutes(e, protectedApi.Group("/settings"))

	protectedApi.GET("/exec/status", e.H(GetCommandExecutionStatus))
	g.POST("/github/events", e.H(ReceiveGithubWebhookEvent))
}
