package routes

import (
	"context"
	"testing"

	"github.com/kohrVid/auth/proto"
)

var s Server = Server{}

func TestCredentialsCheck(t *testing.T) {
	expectedResp := "OK"
	username := "magda"
	password := "magypi123"

	req := &proto.AuthenticationRequest{
		Username: username,
		Password: password,
	}

	resp, err := s.CredentialCheck(context.Background(), req)
	if err != nil {
		t.Errorf(
			"Test failed.\nGot:\n\t%v\n",
			err,
		)
	}

	if resp.Result != expectedResp {
		t.Errorf(
			"Test failed.\nGot:\n\t%v\nExpected:\n\t%v",
			s,
			expectedResp,
		)
	}
}

func TestHealthCheck(t *testing.T) {
	var expectedResp proto.HealthCheckResponse_ServingStatus = 1
	req := &proto.HealthCheckRequest{Service: "auth-api"}

	resp, err := s.HealthCheck(context.Background(), req)
	if err != nil {
		t.Errorf(
			"Test failed.\nGot:\n\t%v\n",
			err,
		)
	}

	if resp.Status != expectedResp {
		t.Errorf(
			"Test failed.\nGot:\n\t%v\nExpected:\n\t%v",
			s,
			expectedResp,
		)
	}
}
