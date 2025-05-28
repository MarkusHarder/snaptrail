package main

import (
	"os"
	"snaptrail/internal/config"
	"snaptrail/internal/db"
	"snaptrail/internal/s3"
	"snaptrail/internal/server"
	"snaptrail/internal/setup"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
	log.Logger = logger
	zerolog.DefaultContextLogger = &logger

	if config.Get().IsDebug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	s3.NewS3ClientFromEnv()

	err := db.Connect(config.Get().DbUrl)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to connect to db")
		return
	}

	err = setup.CreateAdmin()
	if err != nil {
		log.Fatal().Err(err).Msg("unable to retrieve or create admin")
		return
	}
	s := server.New(config.Get().UiDir)
	s.Start()
}
