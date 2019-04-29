GOCMD=go
GOBUILD=$(GOCMD) build 
BUILDPATH=src/main.go
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=podpitcher

build:
	$(GOBUILD) -o $(BINARY_NAME) $(BUILDPATH)
install:
	$(GOBUILD) -o $(GOPATH)/bin/$(BINARY_NAME) $(BUILDPATH)
