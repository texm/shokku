package server

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"gitlab.com/texm/shokku/internal/server/dto"
	"gitlab.com/texm/shokku/internal/server/middleware"
	"io/fs"
	"net/http"
)

func initRouter(debugMode bool, appFS fs.FS) *echo.Echo {
	r := echo.New()
	r.Debug = debugMode
	r.HideBanner = true
	r.Validator = dto.NewRequestValidator()
	r.HTTPErrorHandler = errorHandler(debugMode)

	r.Use(middleware.Recover(debugMode))
	r.Use(middleware.Secure())
	r.Use(middleware.RequestLogger(debugMode))
	r.Use(middleware.StaticFiles(appFS))
	r.Use(middleware.CSRF())

	return r
}

func errorHandler(debug bool) func(error, echo.Context) {
	return func(err error, c echo.Context) {
		if c.Response().Committed {
			return
		}

		httpErr, ok := err.(*echo.HTTPError)
		if ok && httpErr.Internal != nil {
			if herr, ok := httpErr.Internal.(*echo.HTTPError); ok {
				httpErr = herr
			}
		} else if !ok {
			httpErr = &echo.HTTPError{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			}
		}

		// Issue #1426
		code := httpErr.Code
		message := httpErr.Message
		if m, strOk := httpErr.Message.(string); strOk {
			if debug {
				message = echo.Map{"message": m, "error": httpErr.Error()}
			} else {
				message = echo.Map{"message": m}
			}
		}

		// Send response
		if c.Request().Method == http.MethodHead { // Issue #608
			err = c.NoContent(httpErr.Code)
		} else {
			err = c.JSON(code, message)
		}
		if err != nil {
			log.Error().Err(err).Msg("error handler response")
		}
	}
}
