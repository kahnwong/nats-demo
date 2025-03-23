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

## Refs

- <https://docs.nats.io/running-a-nats-service/nats_docker/jetstream_docker>
