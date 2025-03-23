nats-start:
	docker compose -f compose-nats.yaml up -d || podman compose -f compose-nats.yaml up -d
publish:
	go run . publish
