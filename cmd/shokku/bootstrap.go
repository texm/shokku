package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gitlab.com/texm/shokku/internal/server"
	"os"
)

func bootstrapServer() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	if err := server.Bootstrap(); err != nil {
		log.Fatal().Err(err).Msg("failed to create key")
	}
}
