package main

import (
	"context"
	"embed"
	"io/fs"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gitlab.com/texm/shokku/internal/env"
	"gitlab.com/texm/shokku/internal/server"
)

//go:embed all:dist/*
var embeddedFiles embed.FS

func runServer() {
	cfg, err := server.LoadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config")
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if cfg.DebugMode {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	appFS, err := fs.Sub(embeddedFiles, ".")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create embedded fs")
	}

	e, err := server.New(cfg, appFS)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create server service")
	}

	address := net.JoinHostPort(cfg.Host, cfg.Port)

	s := &http.Server{
		Addr:    address,
		Handler: e.Router,
	}

	defer shutdown(e, s)

	go serve(s)
	log.Info().Msgf("Serving on %s", address)

	quit := make(chan os.Signal, 1)
	defer close(quit)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

func serve(s *http.Server) {
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("Fatal server error")
	}
}

func shutdown(e *env.Env, s *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := e.Shutdown(); err != nil {
		log.Error().Err(err).Msg("failed to shut down env")
	}

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("failed to shut down server")
	}

	log.Info().Msg("Shutdown server successfully")
}
