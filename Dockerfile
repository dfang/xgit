FROM golang:1.21 AS builder
MAINTAINER dfang <df1228@gmail.com>

ENV GOPATH /go
ENV GO111MODULE on

COPY . /go/src/github.com/dfang/xgit
WORKDIR /go/src/github.com/dfang/xgit

RUN --mount=type=cache,target=/go/pkg/mod go mod download
RUN --mount=type=cache,target=/go/pkg/mod --mount=type=cache,target=/root/.cache/go-build go install


# FROM golang:1.21
FROM alpine
RUN apk update && add git ca-certificates --no-cache
COPY --from=builder /go/bin/xgit  /go/bin/xgit
CMD ["/go/bin/xgit"]