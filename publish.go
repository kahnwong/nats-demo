package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/nats-io/nats.go/jetstream"
	"github.com/rs/zerolog/log"
)

func publish(js jetstream.JetStream) {
	iterations, _ := stringToFloat(os.Getenv("PUBLISH_BATCH_SIZE"))

	var wg sync.WaitGroup
	wg.Add(iterations)
	for range iterations {
		go func() {
			currentTime := time.Now()
			payload := fmt.Sprintf("%s - %v", currentTime.Format("2006-01-02 15:04:05"), rand.Int())

			_, err := js.PublishAsync("events.sample_input", []byte(payload))
			if err != nil {
				log.Error().Err(err).Msg("Failed to publish async")
			}

			wg.Done()
		}()
	}
	wg.Wait()
	//time.Sleep(300 * time.Millisecond)

	select {
	case <-js.PublishAsyncComplete():
		log.Info().Msg("Published messages")
	case <-time.After(time.Second):
		log.Fatal().Msg("Publish took too long")
	}
}
