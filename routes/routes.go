package routes

import (
	"context"

	"github.com/kohrVid/auth/proto"
)

type Server struct{}

func (s *Server) CredentialCheck(ctx context.Context, r *proto.AuthenticationRequest) (*proto.AuthenticationResponse, error) {
	// TODO - Actually authenticate against a database
	return &proto.AuthenticationResponse{Result: "OK"}, nil
}
