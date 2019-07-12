# GOFLAGS is the flags for the go compiler. Currently, only the commit is passed.
GOFLAGS = -ldflags "-X main.Version=$(shell git rev-parse --short HEAD)"

# force go modules
GO111MODULE = on
export GO111MODULE

# all is the default target
all: test

# help prints a help screen
help:
	@echo "build     - go build"
	@echo "install   - go install"
	@echo "test      - gotestsum"
	@echo "tools     - Install dependencies"
	@echo "gofmt     - go fmt"
	@echo "linux     - go build linux/amd64"
	@echo "clean     - remove temp files"

# build compiles fabio and the test dependencies
build: gofmt
	go build -mod=vendor $(GOFLAGS)

# test runs the tests
test: build
	gotestsum -- -mod=vendor -test.timeout 15s ./...

# gofmt runs gofmt on the code
gofmt:
	gofmt -s -w `find . -type f -name '*.go' | grep -v vendor`

# linux builds a linux binary
linux:
	GOOS=linux GOARCH=amd64 go build -tags netgo $(GOFLAGS)

# install runs go install
install:
	go install -mod=vendor $(GOFLAGS)

# clean removes intermediate files
clean:
	go clean -mod=vendor
	rm -rf dist fuschia
	find . -name '*.test' -delete

.PHONY: all build clean gofmt help install linux test

