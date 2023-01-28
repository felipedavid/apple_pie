up:
	docker-compose up -d

build: build_broker
	docker-compose down
	docker-compose up --build -d

build_broker:
	@echo "Building broker..."
	cd ./broker && env GOOS=linux CGO_ENABLED=0 go build -o broker_bin ./cmd/api

down:
	docker-compose down