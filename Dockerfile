FROM golang:1.6
MAINTAINER Octoblu, Inc. <docker@octoblu.com>

WORKDIR /go/src/github.com/octoblu/health-checker-upper
COPY . /go/src/github.com/octoblu/health-checker-upper

RUN env CGO_ENABLED=0 go build -o health-checker-upper -a -ldflags '-s' .

CMD ["./health-checker-upper"]
