package main

import (
	"context"
	"fmt"
	"net"
	"path/filepath"

	certs "github.com/kohrVid/auth/certs/helpers"
	"github.com/kohrVid/auth/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	certPath := filepath.Join("..", "certs")
	certFile, keyFile := certs.TlsCerts(certPath, "server.crt", "server.key")
	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		log.Fatalf("unable to construct TLS credentials from certificate: %v", err)
	}

	s := grpc.NewServer(grpc.Creds(creds))
	lis, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Println("Listening on localhost:9999...")

	proto.RegisterAuthenticationServiceServer(s, &server{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) CredentialCheck(ctx context.Context, r *proto.AuthenticationRequest) (*proto.AuthenticationResponse, error) {
	// TODO - Actually authenticate against a database
	return &proto.AuthenticationResponse{Result: "OK"}, nil
}
