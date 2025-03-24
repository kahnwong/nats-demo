# NATs Demo

## Setup

- NATs with jetstream support
- go app with two entrypoints: publish or subscribe (will keep listening to streams)

## Usage

```bash
make nats-start
make publish
make subscribe
```

Access ui via <http://localhost:31311>

## Throughput

On linux, in which a subscriber is already active, can publish 40k messages in 283.36 milliseconds. Which translates to roughly 141k in 1 second.

## Refs

- <https://docs.nats.io/running-a-nats-service/nats_docker/jetstream_docker>
