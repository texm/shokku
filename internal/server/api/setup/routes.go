package setup

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/texm/shokku/internal/env"
)

func RegisterRoutes(e *env.Env, g *echo.Group) {
	g.GET("/status", e.H(GetStatus))
	g.GET("/verify-key", e.H(GetSetupKeyValid))

	g.POST("/github/create-app", e.H(CreateGithubApp))
	g.GET("/github/install-info", e.H(GetGithubAppInstallInfo))
	g.GET("/github/status", e.H(GetGithubSetupStatus))
	g.POST("/github/completed", e.H(CompleteGithubSetup))

	g.POST("/password", e.H(CompletePasswordSetup))
	g.POST("/totp/new", e.H(GenerateTotp))
	g.POST("/totp/confirm", e.H(ConfirmTotp))
}
