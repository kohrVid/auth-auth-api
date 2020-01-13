FROM golang:1.12.7

MAINTAINER kohrVid@zoho.com
WORKDIR $GOPATH/src/github.com/kohrVid/auth-auth-api

RUN go get -v ./...
RUN go install -v ./...

EXPOSE 9999

CMD [ "auth-api" ]
