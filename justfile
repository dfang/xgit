# https://github.com/go-semantic-release/example-go-application

flags := ("-X 'main.buildTimestamp=" + `date -u '+%Y-%m-%d %H:%M:%S'` + "'"
          + " -X 'main.xgitVersion=" + `git describe --tags HEAD --abbrev=0` + "'"
          + " -X 'main.xgitBuild=" + `git rev-parse --short HEAD` + "'"
          + " -X 'main.goVersion=" + `go version` + "'")

alias b := build

default:
    just --choose

build:
    # GOOS=darwin GOARCH=amd64 go build -ldflags "{{ flags }} -s -w"  -o tmp/darwin/xgit .
    # GOOS=linux GOARCH=amd64 go build -ldflags "{{ flags }} -s -w"  -o tmp/linux/xgit .
    # GOOS=windows GOARCH=amd64 go build -ldflags "{{ flags }} -s -w"  -o tmp/windows/xgit.exe .
    # goreleaser release --snapshot --skip-publish --clean
    goreleaser build --single-target --snapshot --clean

release:
    # goreleaser release --skip-publish --clean
    semantic-release --provider github --ci-condition default --hooks goreleaser --provider-opt "slug=dfang/xgit"

install:
    go install -ldflags "{{ flags }} -s -w"

lint:
    go vet ./...
    golangci-lint run ./...

build-docker-image:
    docker build -t dfang/xgit .
