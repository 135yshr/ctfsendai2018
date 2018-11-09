package ctfsendai2018

import (
	"fmt"
	"sort"
)

// UserRepository is repository manipulating user information
type UserRepository interface {
	Add(*User) error
	List() ([]*User, error)
	FetchByEMail(string) (*User, error)
}

// NewUserRepository is Create new instance
func NewUserRepository() UserRepository {
	return &userRepository{m: make(map[string]*User)}
}

type userRepository struct {
	m map[string]*User
}

func (r *userRepository) Add(u *User) error {
	r.m[u.EMail] = u
	return nil
}

func (r *userRepository) List() ([]*User, error) {
	l := make([]string, len(r.m))
	x := 0
	for k := range r.m {
		l[x] = k
		x++
	}
	sort.Strings(l)
	list := make([]*User, len(l))
	for n, k := range l {
		list[n] = r.m[k]
	}
	return list, nil
}

func (r *userRepository) FetchByEMail(email string) (*User, error) {
	if v, ok := r.m[email]; ok {
		return v, nil
	}
	return nil, fmt.Errorf("Not found %s", email)
}
