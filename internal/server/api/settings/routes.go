package settings

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/texm/shokku/internal/env"
)

func RegisterRoutes(e *env.Env, g *echo.Group) {
	g.GET("/versions", e.H(GetVersions))

	g.GET("/events", e.H(GetEventLogs))
	g.GET("/events/list", e.H(GetEventLogsList))
	g.POST("/events", e.H(SetEventLoggingEnabled))

	g.GET("/users", e.H(GetUsers))
	g.GET("/ssh-keys", e.H(GetSSHKeys))

	g.GET("/domains", e.H(GetGlobalDomains))
	g.POST("/domains", e.H(AddGlobalDomain))
	g.DELETE("/domains", e.H(RemoveGlobalDomain))

	g.GET("/networks", e.H(ListNetworks))
	g.POST("/networks", e.H(CreateNetwork))
	g.DELETE("/networks", e.H(DestroyNetwork))

	g.GET("/plugins", e.H(ListPlugins))

	g.GET("/registry", e.H(GetDockerRegistryReport))
	g.POST("/registry", e.H(SetDockerRegistry))

	g.POST("/git-auth", e.H(AddGitAuth))
	g.DELETE("/git-auth", e.H(RemoveGitAuth))
}
