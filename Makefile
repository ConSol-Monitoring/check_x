#!/usr/bin/make -f

MAKE:=make
SHELL:=bash

all:

updatedeps:
	$(MAKE) clean
	go mod download
	go get -u ./...
	go get -t -u ./...
	go mod tidy

vendor:
	go mod download
	go mod tidy
	go mod vendor

test: fmt vendor
	go test -v -timeout=1m ./ ./Units
	if grep -rn TODO: *.go ./Units; then exit 1; fi

clean:

fmt:
	go vet -all -assign -atomic -bool -composites -copylocks -nilfunc -rangeloops -unsafeptr -unreachable . ./Units
	gofmt -w -s *.go ./Units/*.go

