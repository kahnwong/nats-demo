nats-start:
	docker compose -f compose-nats.yaml up -d || podman compose -f compose-nats.yaml up -d
teardown:
	docker compose -f compose-nats.yaml down || podman compose -f compose-nats.yaml down
publish:
	go run . publish
subscribe:
	go run . subscribe
