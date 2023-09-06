mocks:
	go generate ./...
run:
	go run .
test-unit:
	go test ./service
docker-up:
	docker-compose up
docker-down:
	docker-compose down
