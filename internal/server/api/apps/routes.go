package apps

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/texm/shokku/internal/env"
)

func RegisterRoutes(e *env.Env, g *echo.Group) {
	g.GET("/list", e.H(GetAppsList))
	g.GET("/report", e.H(GetAppsProcessReport))
	g.GET("/overview", e.H(GetAppOverview))

	g.POST("/create", e.H(CreateApp))
	g.POST("/start", e.H(StartApp))
	g.POST("/stop", e.H(StopApp))
	g.POST("/restart", e.H(RestartApp))
	g.POST("/rebuild", e.H(RebuildApp))

	setupGroup := g.Group("/setup")
	setupGroup.GET("/status", e.H(GetAppSetupStatus))
	setupGroup.GET("/config", e.H(GetAppSetupConfig))
	setupGroup.POST("/new-repo", e.H(SetupAppNewRepo))
	setupGroup.POST("/sync-repo", e.H(SetupAppSyncRepo))
	setupGroup.POST("/pull-image", e.H(SetupAppPullImage))
	setupGroup.POST("/upload-archive", e.H(SetupAppUploadArchive))

	g.POST("/destroy", e.H(DestroyApp))
	g.GET("/info", e.H(GetAppInfo))
	g.POST("/rename", e.H(RenameApp))
	g.GET("/services", e.H(GetAppServices))

	g.GET("/deploy-checks", e.H(GetAppDeployChecks))
	g.POST("/deploy-checks", e.H(SetAppDeployChecks))

	process := g.Group("/process")
	process.GET("/list", e.H(GetAppProcesses))
	process.GET("/report", e.H(GetAppProcessReport))
	process.POST("/deploy-checks", e.H(SetAppProcessDeployChecks))
	process.POST("/resources", e.H(SetAppProcessResources))
	process.GET("/scale", e.H(GetAppProcessScale))
	process.POST("/scale", e.H(SetAppProcessScale))

	g.GET("/letsencrypt", e.H(GetAppLetsEncryptEnabled))
	g.POST("/letsencrypt", e.H(SetAppLetsEncryptEnabled))

	g.GET("/domains", e.H(GetAppDomainsReport))
	g.POST("/domains/state", e.H(SetAppDomainsEnabled))

	g.POST("/domain", e.H(AddAppDomain))
	g.DELETE("/domain", e.H(RemoveAppDomain))

	g.GET("/networks", e.H(GetAppNetworksReport))
	g.POST("/networks", e.H(SetAppNetworks))

	g.GET("/logs", e.H(GetAppLogs))

	g.GET("/config", e.H(GetAppConfig))
	g.POST("/config", e.H(SetAppConfig))

	g.GET("/storage", e.H(GetAppStorage))
	g.POST("/storage/mount", e.H(MountAppStorage))
	g.POST("/storage/unmount", e.H(UnmountAppStorage))

	g.GET("/builder", e.H(GetAppBuilder))
	g.POST("/builder", e.H(SetAppBuilder))

	g.GET("/build-directory", e.H(GetAppBuildDirectory))
	g.POST("/build-directory", e.H(SetAppBuildDirectory))
	g.DELETE("/build-directory", e.H(ClearAppBuildDirectory))
}
