FROM golang:1.8
MAINTAINER linyows <linyows@gmail.com>

RUN mkdir -p /go/src/github.com/linyows/go-plugin-sandbox
WORKDIR /go/src/github.com/linyows/go-plugin-sandbox
