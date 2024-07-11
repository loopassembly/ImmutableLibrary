
PROJECT_NAME = immutablelibrary


GO = go


DOCKER = docker


build:
	$(GO) build -o hack4bengal


run:
	./server.go

# Run tests
test:
	$(GO) testing ./...


clean:
	rm -f $(PROJECT_NAME)


docker-build:
	$(DOCKER) build -t $(PROJECT_NAME) .


docker-run:
	$(DOCKER) run -p 8080:8080 $(PROJECT_NAME)


docker-push:
	$(DOCKER) push $(PROJECT_NAME)


all: build run
