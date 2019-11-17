# Auth API

This is an authentication service that I'm writing with gRPC.

<!-- vim-markdown-toc GFM -->

* [gRPC server](#grpc-server)
* [Mockgen](#mockgen)

<!-- vim-markdown-toc -->

## gRPC server

In the dev environment, it should be possible to run the server with the following command:

    go run main.go

The server listens to TCP port 9999 by default.

## Mockgen

The Mockgen tool in [GoMock](https://github.com/golang/mock#running-mockgen) has
been used to generate a mock client for use in client-side tests. The code for
this can be found in `./mock_pb`

If any changes are made to protobuf, most probably this will also need to be
regenerated. The command used to generate the existing code was as follows:

    mockgen github.com/kohrVid/auth/auth-api/pb AuthenticationServiceClient > mock-auth/pbMock.go
