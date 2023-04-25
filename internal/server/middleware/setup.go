package middleware

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/texm/shokku/internal/env"
	"net/http"
	"strings"
)

const (
	notSetupErr    = "server not setup"
	invalidKeyErr  = "setup key invalid"
	setupKeyHeader = "X-Setup-Key"
)

func ServerSetupBlocker(e *env.Env, setupKey string) echo.MiddlewareFunc {
	logger := middlewareLogger("setup")
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			reqPath := c.Request().URL.Path

			if !strings.HasPrefix(reqPath, "/api") {
				return next(c)
			}

			if reqPath == "/api/github/events" {
				return next(c)
			}

			if reqPath == "/api/setup/status" {
				return next(c)
			}

			isSetupRoute := strings.HasPrefix(reqPath, "/api/setup")
			if e.SetupCompleted {
				if isSetupRoute {
					return echo.NewHTTPError(http.StatusBadRequest, "already set up")
				}
				return next(c)
			}

			providedKey := c.Request().Header.Get(setupKeyHeader)
			if providedKey != setupKey {
				logger.Debug().
					Str("path", reqPath).
					Str("provided", providedKey).
					Msg("invalid setup key")
				return echo.NewHTTPError(http.StatusForbidden, invalidKeyErr)
			}

			if isSetupRoute {
				return next(c)
			}

			logger.Debug().Str("path", reqPath).Msg("not setup path")
			return echo.NewHTTPError(http.StatusForbidden, notSetupErr)
		}
	}
}
