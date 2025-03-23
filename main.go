package main

import (
	"os"

	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// ----- init -----
	// logger
	if os.Getenv("MODE") == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	// nats
	url := os.Getenv("NATS_URL")
	if url == "" {
		url = nats.DefaultURL
	}

	nc, err := nats.Connect(url)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to NATS")
	} else {
		log.Info().Msg("Connected to NATS")
	}

	defer func(nc *nats.Conn) {
		err := nc.Drain()
		if err != nil {
			log.Error().Err(err).Msg("Failed to drain NATS")
		}
	}(nc)
}
