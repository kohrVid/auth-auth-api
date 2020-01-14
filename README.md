# Auth API

This is an authentication service that I'm writing with gRPC.

<!-- vim-markdown-toc GFM -->

* [gRPC server](#grpc-server)
* [Docker](#docker)
  * [LDAP](#ldap)
  * [API](#api)

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


### LDAP

The OpenLDAP instance can be run with the following commands:

    docker build -f Dockerfile.ldap -t="kohrvid/auth-openldap:latest" .
    docker run --privileged -d -p 389:389 kohrvid/auth-openldap:latest


### API

To build the API, run:

    docker build -t="kohrvid/auth-api:latest" .
    docker run --net kohrvid-dev -p 9999:9999 kohrvid/auth-api:latest
