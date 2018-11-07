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

type userRepository struct{}

func (r *userRepository) Add(u *User) error {
	return nil
}

func (r *userRepository) List() ([]*User, error) {
	return []*User{&User{}}, nil
}
