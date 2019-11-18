package session

import (
	"fmt"

	ldap "github.com/go-ldap/ldap/v3"
)

func CheckCredentials(l *ldap.Conn, username, password string) error {
	defer l.Close()
	err := l.Bind(fmt.Sprintf("cn=%v,dc=kohrvid,dc=com", username), password)
	return err
}
