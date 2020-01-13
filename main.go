package main

import (
	"fmt"
	"net"
	"path/filepath"

	"github.com/kohrVid/auth-auth-api/routes"
	certs "github.com/kohrVid/auth-certs/helpers"
	proto "github.com/kohrVid/auth-proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

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

	proto.RegisterAuthenticationServiceServer(s, &routes.Server{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
