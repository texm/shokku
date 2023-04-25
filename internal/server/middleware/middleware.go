package middleware

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func middlewareLogger(ware string) zerolog.Logger {
	return log.With().Str("middleware", ware).Logger()
}
