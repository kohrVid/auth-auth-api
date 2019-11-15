# Auth API

This is an authentication service that I'm writing with gRPC.


<!-- vim-markdown-toc GFM -->

* [Protobuf](#protobuf)
* [gRPC server](#grpc-server)

<!-- vim-markdown-toc -->

## Protobuf

Protobuf files can be found in `./pb`. If changes are made to them, the `protoc`
command will need to be re-run.

Firstly, ensure that you have protoc installed. The best way to install this on
my machine was with the following:

    PROTOC_ZIP=protoc-3.7.1-linux-x86_64.zip
    curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.7.1/$PROTOC_ZIP
    sudo unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
    sudo unzip -o $PROTOC_ZIP -d /usr/local 'include/*'
    rm -f $PROTOC_ZIP

To regenerate the go code associated with the protobuf files, run:

    protoc -I . --go_out=plugins=grpc:. ./*.proto

## gRPC server

In the dev environment, it should be possible to run the server with the following command:

    go run main.go

The server listens to TCP port 9999 by default.
