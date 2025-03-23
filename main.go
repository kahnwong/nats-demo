package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/nats-io/nats.go/jetstream"

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

	// init stream
	js, err := jetstream.New(nc)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to JetStream")
	} else {
		log.Info().Msg("Connected to JetStream")
	}

	cfg := jetstream.StreamConfig{
		Name:     os.Getenv("NATS_STREAM_NAME"),
		Subjects: []string{"events.>"},
		MaxMsgs:  2000,
	}

	cfg.Storage = jetstream.FileStorage

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := js.CreateOrUpdateStream(ctx, cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create stream")
	} else {
		log.Info().Msg("Created stream")
	}

	// run modes
	if len(os.Args) > 1 {
		if os.Args[1] == "publish" {
			publish(js)
		} else if os.Args[1] == "subscribe" {
			cons, _ := stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
				Durable:   "CONS",
				AckPolicy: jetstream.AckExplicitPolicy,
			})

			// Receive messages continuously in a callback
			iter, _ := cons.Messages()
			for {
				msg, err := iter.Next()
				if err != nil {
					log.Error().Err(err).Msg("Failed to read next message")
				}

				// ack
				log.Info().Msgf("Received a message: %s", string(msg.Data()))
				err = msg.Ack()
				if err != nil {
					log.Error().Err(err).Msg("Failed to ack message")
				}
			}
			//iter.Stop()
		}

		// qa
		printStreamState(ctx, stream)
	}
}

func printStreamState(ctx context.Context, stream jetstream.Stream) {
	info, _ := stream.Info(ctx)
	b, _ := json.MarshalIndent(info.State, "", " ")
	log.Info().Msg("Inspecting stream info")
	fmt.Println(string(b))
}
