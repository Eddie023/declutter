FROM golang:1.21 AS build 

ARG VERSION
WORKDIR /go/src/github.com/eddie023/declutter

RUN apt-get update \
    && apt-get install -y -q --no-install-recommends 

ENTRYPOINT [ "/bin/sh", "-c", "go test $(go list --buildvcs=false ./...)" ]