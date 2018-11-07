package ctfsendai2018

// UserRepository is repository manipulating user information
type UserRepository interface {
	Add(*User) error
	List() ([]*User, error)
}

// NewUserRepository is Create new instance
func NewUserRepository() UserRepository {
	return &userRepository{}
}

type userRepository struct {
	list []*User
}

func (r *userRepository) Add(u *User) error {
	r.list = append(r.list, u)
	return nil
}

func (r *userRepository) List() ([]*User, error) {
	return r.list, nil
}
