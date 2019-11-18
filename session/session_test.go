package session

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	ldap "github.com/go-ldap/ldap/v3"
)

func TestCheckCredentialsArePresent(t *testing.T) {
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", "localhost", 389))
	if err != nil {
		t.Errorf(
			"Test failed.\nUnable to connect to LDAP",
		)
	}

	err = CheckCredentials(l, "admin", "nevermore")
	if err != nil {
		t.Errorf(
			"Test failed.\nGot:\n\t%v\n",
			err,
		)

	}
}

func TestCheckCredentialsAreMissing(t *testing.T) {
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", "localhost", 389))
	if err != nil {
		t.Errorf(
			"Test failed.\nUnable to connect to LDAP",
		)
	}

	expectedErr := errors.New(`LDAP Result Code 49 "Invalid Credentials":`).Error()
	err = CheckCredentials(l, "admin", "wrong_password")
	if strings.Trim(err.Error(), " ") != expectedErr {
		t.Errorf(
			"Test failed.\nGot:\n\t%v\nExpected:\n\t%v",
			err,
			expectedErr,
		)

	}
}
