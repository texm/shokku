package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"math/rand"
	"os"
	"time"
)

const (
	BootstrapCmd = "bootstrap"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	if len(os.Args) <= 1 {
		runServer()
		return
	} else {
		switch cmd := os.Args[1]; cmd {
		case BootstrapCmd:
			bootstrapServer()
		default:
			log.Fatal().Msgf("invalid command '%s' provided", cmd)
		}
	}
}
