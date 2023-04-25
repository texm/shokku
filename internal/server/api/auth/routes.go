package auth

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/texm/shokku/internal/env"
)

func RegisterRoutes(e *env.Env, g *echo.Group, authG *echo.Group) {
	g.POST("/login", e.H(HandlePasswordLogin))
	g.GET("/github", e.H(GetGithubAuthInfo))
	g.POST("/github/auth", e.H(CompleteGithubAuth))
	g.POST("/logout", e.H(HandleLogout))

	authG.GET("/details", e.H(HandleGetDetails))
	authG.POST("/refresh", e.H(HandleRefreshAuth))
}
