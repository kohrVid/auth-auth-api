package routes

import (
	"context"
	"testing"

	"github.com/kohrVid/auth/proto"
)

func TestCredentialsCheck(t *testing.T) {
	s := Server{}
	expectedResp := "OK"
	username := "magda"
	password := "magypi123"

	req := &proto.AuthenticationRequest{Username: username, Password: password}
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
