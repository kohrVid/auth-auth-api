package routes

import (
	"context"
	"fmt"

	"github.com/go-ldap/ldap/v3"
	"github.com/kohrVid/auth/auth-api/session"
	"github.com/kohrVid/auth/proto"
)

type Server struct{}

func (s *Server) CredentialCheck(ctx context.Context, r *proto.AuthenticationRequest) (*proto.AuthenticationResponse, error) {
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", "localhost", 389))
	if err != nil {
		return &proto.AuthenticationResponse{Result: "SERVER ERROR"}, err
	}

	err = session.CheckCredentials(l, r.Username, r.Password)
	if err != nil {
		return &proto.AuthenticationResponse{Result: "UNAUTHORISED"}, err

	}

	return &proto.AuthenticationResponse{Result: "OK"}, nil
}

func (s *Server) HealthCheck(ctx context.Context, r *proto.HealthCheckRequest) (*proto.HealthCheckResponse, error) {
	return &proto.HealthCheckResponse{Status: 1}, nil
}
