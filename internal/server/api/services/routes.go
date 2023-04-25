package services

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/texm/shokku/internal/env"
)

func RegisterRoutes(e *env.Env, g *echo.Group) {
	g.GET("/list", e.H(ListServices))
	g.GET("/info", e.H(GetServiceInfo))
	g.GET("/type", e.H(GetServiceType))
	g.GET("/logs", e.H(GetServiceLogs))

	g.POST("/create", e.H(CreateNewGenericService))
	g.POST("/clone", e.H(CloneService))
	g.POST("/start", e.H(StartService))
	g.POST("/stop", e.H(StopService))
	g.POST("/restart", e.H(RestartService))
	g.POST("/destroy", e.H(DestroyService))

	g.POST("/link", e.H(LinkGenericServiceToApp))
	g.POST("/unlink", e.H(UnlinkGenericServiceFromApp))
	g.GET("/linked-apps", e.H(GetServiceLinkedApps))

	backups := g.Group("/backups")
	backups.GET("/report", e.H(GetServiceBackupReport))
	backups.POST("/auth", e.H(SetServiceBackupAuth))
	backups.POST("/bucket", e.H(SetServiceBackupBucket))
	backups.POST("/run", e.H(RunServiceBackup))
	backups.POST("/schedule", e.H(SetServiceBackupSchedule))
	backups.DELETE("/schedule", e.H(RemoveServiceBackupSchedule))
	backups.POST("/encryption", e.H(SetServiceBackupEncryption))
	backups.DELETE("/encryption", e.H(RemoveServiceBackupEncryption))
}
