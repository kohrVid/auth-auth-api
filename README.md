# Auth API

This is an authentication service that I'm writing with gRPC.

<!-- vim-markdown-toc GFM -->

* [gRPC server](#grpc-server)
* [Docker](#docker)

<!-- vim-markdown-toc -->

## gRPC server

In the dev environment, it should be possible to run the server with the
following command:

    go run main.go

The server listens to TCP port 9999 by default.

## Docker

To run this server locally with [Docker](https://www.docker.com/), first create
the network:

    docker network create --driver bridge kohrvid-dev

Then build and run the image:

    docker build -t="auth-api" .
    docker run --net kohrvid-dev -p 9999:9999 auth-api
