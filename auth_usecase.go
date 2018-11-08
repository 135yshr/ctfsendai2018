package ctfsendai2018

// AuthUsecase is service manipulating authentication information
type AuthUsecase interface {
	Register(string, string) error
}

// NewAuthUsecase is create new instance
func NewAuthUsecase(a AuthRepository, u UserRepository) AuthUsecase {
	return &authUsecase{a, u}
}

type authUsecase struct {
	authRep AuthRepository
	userRep UserRepository
}

func (s *authUsecase) Register(email, passwd string) error {
	u, err := NewUser(email, passwd)
	if err != nil {
		return err
	}
	return s.userRep.Add(u)
}
