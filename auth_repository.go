package ctfsendai2018

// AuthRepository is repository manipulating authentication information
type AuthRepository interface {
}

// NewAuthRepository is create new instance
func NewAuthRepository() AuthRepository {
	return &authRepository{}
}

type authRepository struct{}
