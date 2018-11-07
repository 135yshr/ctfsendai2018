package ctfsendai2018

// Authentication is structure that represents authentication information
type Authentication struct {
	Token string
}

// AuthRepository is repository manipulating authentication information
type AuthRepository interface {
	Login(*User) (*Authentication, error)
}

// NewAuthRepository is create new instance
func NewAuthRepository() AuthRepository {
	return &authRepository{}
}

type authRepository struct{}

func (r *authRepository) Login(u *User) (*Authentication, error) {
	return &Authentication{Token: "eyJleHAiOjE0OTQ1MTQyNjksInVzZXIiOiLnrqHnkIbogIUifQo="}, nil
}
