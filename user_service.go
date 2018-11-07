package ctfsendai2018

import "errors"

// UserService is service manipulating user information
type UserService interface {
	Add(uint8, *User) error
	List(uint8) ([]*User, error)
}

// NewUserService is create new instance
func NewUserService(rep UserRepository) UserService {
	return &userService{rep}
}

type userService struct {
	rep UserRepository
}

func (s *userService) Add(auth uint8, u *User) error {
	if auth != 1 {
		return errors.New("Don't have permission")
	}
	return s.rep.Add(u)
}

func (s *userService) List(auth uint8) ([]*User, error) {
	return s.rep.List()
}
