package middleware

import (
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"io/fs"
	"net/http"
	"strings"
)

func staticFileSkipperFunc(c echo.Context) bool {
	if strings.HasPrefix(c.Request().URL.Path, "/api") {
		c.Set("static", false)
		return true
	}
	c.Set("static", true)
	return false
}

func StaticFiles(staticFS fs.FS) echo.MiddlewareFunc {
	cfg := echomiddleware.StaticConfig{
		Root:       "dist",
		Index:      "app.html",
		HTML5:      true,
		Filesystem: http.FS(staticFS),
		Skipper:    staticFileSkipperFunc,
	}
	return echomiddleware.StaticWithConfig(cfg)
}
