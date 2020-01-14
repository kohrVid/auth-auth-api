FROM golang:1.12.7 AS builder

LABEL maintainer="kohrVid@zoho.com"

RUN go version

COPY . "/go/src/github.com/kohrVid/auth-auth-api"
WORKDIR "/go/src/github.com/kohrVid/auth-auth-api"
ADD . $GOPATH/src/github.com/kohrVid/auth-auth-api

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o /auth-auth-api


CMD [ "/auth-auth-api" ]

EXPOSE 9999
