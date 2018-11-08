package ctfsendai2018

// AuthUsecase is service manipulating authentication information
type AuthUsecase interface {
}

// NewAuthService is create new instance
func NewAuthService(rep AuthRepository) AuthUsecase {
	return &authService{rep}
}

type authService struct {
	rep AuthRepository
}
