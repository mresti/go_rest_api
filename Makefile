# Project specific variables
PROJECT=api
OS=$(shell uname)
GOARCH = amd64

# GO env
GOPATH=$(shell pwd)
GO=go
GOCMD=GOPATH=$(GOPATH) $(GO)
GOBUILD = $(GOCMD) build

# Build the project
.PHONY: all
all:	build

.PHONY: build
build: format test compile

.PHONY: compile
compile: darwin linux windows

.PHONY: format
format:
	@for gofile in $$(find ./src -name "*.go"); do \
		echo "formatting" $$gofile; \
		gofmt -w $$gofile; \
	done

.PHONY: test
test:
	$(GOCMD) test -v -race ./src/...

.PHONY: run
run:
	$(GOCMD) run ./src/main.go

multi: build darwin linux windows

darwin:
	GOOS=darwin GOARCH=${GOARCH} $(GOBUILD) -o bin/$(PROJECT)_darwin src/main.go
linux:
	GOOS=linux GOARCH=${GOARCH} $(GOBUILD) -o bin/$(PROJECT)_linux src/main.go
windows:
	GOOS=windows GOARCH=${GOARCH} $(GOBUILD) -o bin/$(PROJECT)_windows.exe src/main.go
