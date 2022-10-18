APP_NAME=fiber-go-exercise
VERSION_TAG?=0.0.1

build:
	CGO_ENABLED=0 go build -ldflags='-extldflags=-static'  -o service cmd/main.go

build-docker:
	docker build --tag="$(APP_NAME):$(VERSION_TAG)" --tag="$(APP_NAME):latest" .

run: build-docker
	docker-compose up -d

stop:
	docker-compose down

clean:
	docker rmi -f $(APP_NAME):$(VERSION_TAG) $(APP_NAME):latest