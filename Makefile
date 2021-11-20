run:
	go run ./cmd/rest-server/main.go -env=".env"
docker_up:
	docker-compose up -d
docker_down:
	docker-compose down