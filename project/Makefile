## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker compose up -d
	@echo "Docker images started!"

## up_build: stops docker compose (if running), builds all projects and starts docker compose
up_build: build_api_service build_nextjs_fe
	@echo "Stopping docker images (if running...)"
	docker compose down
	@echo "Building (when required) and starting docker images..."
	docker compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker compose down
	@echo "Done!"

clean:
	rm -rf ../movies-back-end/apiBinary

## build_api_service: builds the go API service binary as a linux executable
build_api_service:
	@echo "Building go API service binary..."
	cd ../movies-back-end && env GOOS=linux CGO_ENABLED=0 go build -o apiBinary ./cmd/api
	@echo "Done!"

## build_nextjs_fe: builds the nextjs website frontend jar as a linux executable
build_nextjs_fe:
	@echo "Building website frontend..."
	cd ../movies-front-end && npm i --force && npm run build
	@echo "Done!"