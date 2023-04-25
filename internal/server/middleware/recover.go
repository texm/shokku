package middleware

import (
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"strings"
)

func logErrorFunc(logger zerolog.Logger, debug bool) echoMiddleware.LogErrorFunc {
	return func(c echo.Context, err error, stack []byte) error {
		event := logger.Error().Err(err)
		if debug {
			stacklines := strings.Split(string(stack), "\n")
			funcName := strings.TrimRight(strings.SplitAfter(stacklines[5], "(0x")[0], "(0x")
			callSite := strings.Trim(strings.SplitAfter(stacklines[6], " ")[0], "\t ")
			event.Str("func", funcName)
			event.Str("callsite", callSite)
		}
		event.Msg("Recovered from panic")
		return nil
	}
}

func Recover(debug bool) echo.MiddlewareFunc {
	logger := middlewareLogger("recover")
	cfg := echoMiddleware.RecoverConfig{
		LogErrorFunc: logErrorFunc(logger, debug),
	}
	return echoMiddleware.RecoverWithConfig(cfg)
}
