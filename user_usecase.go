package ctfsendai2018

import "errors"

// UserUsecase is service manipulating user information
type UserUsecase interface {
	Add(uint8, *User) error
	List(uint8) ([]*User, error)
}

// NewUserUsecase is create new instance
func NewUserUsecase(rep UserRepository) UserUsecase {
	return &userUsecase{rep}
}

type userUsecase struct {
	rep UserRepository
}

func (s *userUsecase) Add(auth uint8, u *User) error {
	if auth != 1 {
		return errors.New("Don't have permission")
	}
	return s.rep.Add(u)
}

func (s *userUsecase) List(auth uint8) ([]*User, error) {
	return s.rep.List()
}
