package ctfsendai2018

import (
	"errors"
	"fmt"
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
	if err != nil {
		return nil, errors.New("Failed login")
	}
	if user.Password != passwd {
		fmt.Println(user.Password, passwd)
		return nil, errors.New("Do not mutch user or password")
	}
	return u.authRep.Login(user)
}
