package main

import (
	"context"
	"fmt"
	"net"

	"github.com/kohrVid/auth/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	fmt.Println("Listening on localhost:9999...")

	s := grpc.NewServer()
	proto.RegisterAuthenticationServiceServer(s, &server{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) CredentialCheck(ctx context.Context, r *proto.AuthenticationRequest) (*proto.AuthenticationResponse, error) {
	// TODO - Actually authenticate against a database
	return &proto.AuthenticationResponse{Result: "OK"}, nil
}
