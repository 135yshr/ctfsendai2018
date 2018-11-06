// Package ctfsendai2018 is a pseudo customer management system created for the task of CTF Sendai 2018.
// There are three vulnerabilities in this system, and we exploit the vulnerability from the outside to obtain a hidden flag.
package ctfsendai2018

import (
	"fmt"
)

// Specification of user information
// User information has information of name, age, e-mail address, remarks.
// The name must be between 1 and 20 characters in a string.
// The age shall be a numerical value of 0 or more.
// The mail address is a character string and conforms to the mailbox of RFC 5321 and the addr-spec of RFC 5322.
// Remarks are limited to 50 characters in a character string.

// User は、ユーザー情報を表す構造体
type User struct {
}

// NewUser は、新しくユーザー情報のインスタンスを作成する
func NewUser(name string) (*User, error) {
	if name == "" {
		return nil, fmt.Errorf("name can not be empty")
	}
	return &User{}, nil
}
