# Go parameters

VERSION=1.12.4
BINARY_NAME=PugServer
GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLIEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
GOTEST=$(GOCMD) test

all: test build

build:
	$(GOBUILD) -o $(BINARY_NAME)  -v main.go

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLIEAN)
	rm -f $(BINARY_NAME)

run:
	$(GOBUILD) -o $(BINARY_NAME) -v main.go
	./$(BINARY_NAME)

deps:
	$(GOGET) github.com/cihub/seelog
	$(GOGET) github.com/Unknwon/goconfig
	$(GOGET) github.com/gin-gonic/gin


docker-build:

