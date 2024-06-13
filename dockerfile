FROM golang:1.21 AS build 

ARG VERSION
WORKDIR /go/src/github.com/eddie023/declutter

COPY go.mod go.sum ./
RUN go mod download 

ADD . . 
RUN go install -ldflags "-X 'internal/build.Version=$VERSION'" ./cmd/declutter

FROM debian:bookworm
RUN apt-get update \
    && apt-get install -y -q --no-install-recommends 

COPY --from=build /go/bin/declutter /bin
ENTRYPOINT [ "declutter"]