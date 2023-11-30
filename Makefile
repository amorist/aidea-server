Version := $(shell date "+%Y%m%d%H%M")
GitCommit := $(shell git rev-parse HEAD)
DIR := $(shell pwd)
LDFLAGS := -s -w -X main.Version=$(Version) -X main.GitCommit=$(GitCommit) -X main.DEBUG=true

build:
	go build -race -ldflags "$(LDFLAGS)" -o build/debug/aidea-server cmd/main.go

build-release:
	GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o build/release/aidea-server-linux cmd/main.go
	GOOS=linux GOARCH=arm go build -ldflags "$(LDFLAGS)" -o build/release/aidea-server-linux-arm cmd/main.go
	GOOS=darwin GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o build/release/aidea-server-darwin cmd/main.go
	GOOS=windows GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o build/release/aidea-server.exe cmd/main.go

build-linux:
	GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o build/release/aidea-server-linux cmd/main.go

orm:
	# https://github.com/mylxsw/eloquent
	eloquent gen --source 'pkg/repo/model/*.yaml'
	gofmt -s -w pkg/repo/model/*.go

scp:
	scp build/release/aidea-server-linux root@106.15.32.100:/var/www/aidea-server

.PHONY: build build-release orm build-linux
