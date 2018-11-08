// Package ctfsendai2018 is a pseudo customer management system created for the task of CTF Sendai 2018.
// There are three vulnerabilities in this system, and we exploit the vulnerability from the outside to obtain a hidden flag.
package ctfsendai2018

import (
	"encoding/base64"
	"errors"
	"strings"
)

// Specification of user information
// User information has information of name, age, e-mail address, remarks.
// The name must be between 1 and 20 characters in a string.
// The age shall be a numerical value of 0 or more.
// The mail address is a character string and conforms to the mailbox of RFC 5321 and the addr-spec of RFC 5322.
// Remarks are limited to 50 characters in a character string.

// User は、ユーザー情報を表す構造体
type User struct {
	ID       string
	Password string
	EMail    string
	Name     string
	Age      uint8
	Note     string
	Auth     uint8
}

// NewUser は、新しくユーザー情報のインスタンスを作成する
func NewUser(email, passwd string) (*User, error) {
	if email == "" {
		return nil, errors.New("E-mail address is empty")
	}

	id := base64.StdEncoding.EncodeToString([]byte(email))
	return &User{ID: strings.TrimRight(id, "="), EMail: email, Password: passwd, Auth: 2}, nil
}
