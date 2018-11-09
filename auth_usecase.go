package ctfsendai2018

import (
	"errors"
	"strings"
)

// AuthUsecase is service manipulating authentication information
type AuthUsecase interface {
	Register(string, string) error
	Login(string, string) (*Authentication, error)
}

// NewAuthUsecase is create new instance
func NewAuthUsecase(a AuthRepository, u UserRepository) AuthUsecase {
	return &authUsecase{a, u}
}

type authUsecase struct {
	authRep AuthRepository
	userRep UserRepository
}

func (u *authUsecase) Register(email, passwd string) error {
	user, err := NewUser(email, passwd)
	if err != nil {
		return err
	}
	return u.userRep.Add(user)
}

func (u *authUsecase) Login(email, passwd string) (*Authentication, error) {
	user, err := u.userRep.FetchByEMail(email)
	// user, err := NewUser(email, passwd)
	if err != nil {
		if strings.Index(err.Error(), "Not found") != -1 {
			return nil, errors.New("Failed login")
		}
		return nil, err
	}
	return u.authRep.Login(user)
}
