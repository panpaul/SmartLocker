PROJECT=SmartLocker

GO := go

LDFLAGS += -X "main.GitHash=$(shell git rev-parse HEAD)"
LDFLAGS += -X "main.Version=$(shell cat VERSION)"
LDFLAGS += -X "main.BuildTimeStamp=$(shell date -u '+%Y-%m-%d %I:%M:%S')"

.PHONY: all server clean pack run dep fmt

default: all

all: clean dep fmt server

server:
	$(GO) build -o server -ldflags '$(LDFLAGS)' ./cmd/server/

clean:
	rm -rf server
	rm -rf SmartLocker
	rm -rf SmartLocker.zip

pack: clean dep server
	mkdir SmartLocker
	mv server SmartLocker
	cp config_sqlite.yaml SmartLocker/config.yaml
	cp -r resources SmartLocker/
	zip -r9 -o SmartLocker.zip ./SmartLocker

run: dep server
	./server s

dep:
	go get -u ./...

fmt:
	go fmt ./...
