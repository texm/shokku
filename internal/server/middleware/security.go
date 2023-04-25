package middleware

import (
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"net/http"
)

func Secure() echo.MiddlewareFunc {
	// logger := e.Logger.Desugar()
	// debug := e.DebugMode

	// cfg := echomiddleware.SecureConfig{}
	cfg := echoMiddleware.DefaultSecureConfig
	return echoMiddleware.SecureWithConfig(cfg)
}

func CSRF() echo.MiddlewareFunc {
	// we skip requests where cookie authentication was not used,
	// as these are api requests - not from the browser
	cfg := echoMiddleware.CSRFConfig{
		CookieName:     "_csrf",
		CookiePath:     "/",
		CookieSameSite: http.SameSiteStrictMode,
		Skipper: func(c echo.Context) bool {
			return !CheckCookieAuthUsed(c)
		},
	}
	return echoMiddleware.CSRFWithConfig(cfg)
}
