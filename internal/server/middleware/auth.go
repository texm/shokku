package middleware

import (
	"fmt"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"gitlab.com/texm/shokku/internal/env"
	"gitlab.com/texm/shokku/internal/server/auth"
	"strings"
)

const (
	tokenContextKey          = "user-token"
	usedCookieAuthContextKey = "cookie-auth"
)

func skipDuringSetup(e *env.Env, c echo.Context) bool {
	reqPath := c.Request().URL.Path
	return !e.SetupCompleted && strings.HasPrefix(reqPath, "/api/setup")
}

func ProvideUserContext(e *env.Env) echo.MiddlewareFunc {
	logger := middlewareLogger("userContext")
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if skipDuringSetup(e, c) {
				return next(c)
			}

			if err := e.Auth.SetUserContext(c, tokenContextKey); err != nil {
				logger.Error().Err(err).Msg("failed to set context")
				return echo.ErrInternalServerError
			}
			return next(c)
		}
	}
}

func tokenAuthSkipper(e *env.Env) echoMiddleware.Skipper {
	return func(c echo.Context) bool {
		return skipDuringSetup(e, c)
	}
}

func TokenAuth(e *env.Env) echo.MiddlewareFunc {
	config := echoMiddleware.JWTConfig{
		Claims:           &auth.User{},
		SigningKey:       e.Auth.GetSigningKey(),
		TokenLookupFuncs: []echoMiddleware.ValuesExtractor{SplitTokenLookup},
		TokenLookup:      "header:Authorization",
		ContextKey:       tokenContextKey,
		Skipper:          tokenAuthSkipper(e),
	}
	return echoMiddleware.JWTWithConfig(config)
}

func SplitTokenLookup(c echo.Context) ([]string, error) {
	dataCookie, err := c.Request().Cookie(auth.DataCookieName)
	if err != nil {
		return nil, fmt.Errorf("no data cookie: %w", err)
	}
	signatureCookie, err := c.Request().Cookie(auth.SignatureCookieName)
	if err != nil {
		return nil, fmt.Errorf("no signature cookie: %w", err)
	}

	c.Set(usedCookieAuthContextKey, true)
	authToken := dataCookie.Value + "." + signatureCookie.Value
	return []string{authToken}, nil
}

func CheckCookieAuthUsed(c echo.Context) bool {
	v, ok := c.Get(usedCookieAuthContextKey).(bool)
	return ok && v
}
