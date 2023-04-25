package middleware

import (
	"fmt"
	"gitlab.com/texm/shokku/internal/server/dto"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func parseDokkuError(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err == nil {
			return nil
		}

		if dokkuHttpErr := dto.MaybeConvertDokkuError(err); dokkuHttpErr != nil {
			log.Debug().Err(err).Msgf("converted dokku error to %s", dokkuHttpErr.Error())
			return dokkuHttpErr
		}

		log.Error().Err(err).Str("path", c.Path()).Msg("got error")
		return echo.ErrInternalServerError
	}
}

func RequestLogger(debug bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			chainErr := next(c)
			if chainErr != nil {
				c.Error(chainErr)
			}

			req := c.Request()
			res := c.Response()

			if isStaticFile, ok := c.Get("static").(bool); ok && isStaticFile {
				contentType := res.Header().Get("Content-Type")
				isHTML := strings.HasPrefix(contentType, echo.MIMETextHTML)
				if !isHTML {
					return nil
				}
			}

			n := res.Status
			l := log.Info()
			msg := "Success"
			switch {
			case n >= 500:
				l = log.Error().Err(chainErr)
				// logger.With(zap.Error(chainErr)).Error("Server error", fields...)
				msg = "Server error"
			case n >= 400:
				msg = "Client error"
			case n >= 300:
				msg = "Redirect"
			}

			if debug {
				l.Str("request", fmt.Sprintf("%s %s", req.Method, req.RequestURI))
				l.Int("status", res.Status)
			} else {
				l.Str("remote_ip", c.RealIP())
				l.Str("latency", time.Since(start).String())
				l.Str("host", req.Host)
				l.Str("request", fmt.Sprintf("%s %s", req.Method, req.RequestURI))
				l.Int("status", res.Status)
				l.Int64("size", res.Size)
			}

			id := req.Header.Get(echo.HeaderXRequestID)
			if id != "" {
				l.Str("request_id", id)
			}

			l.Msg(msg)

			return nil
		}
	}
}
