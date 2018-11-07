package ctfsendai2018

import (
	"errors"

	jwt "github.com/dgrijalva/jwt-go"
)

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
	if u == nil {
		return nil, errors.New("User information is empty")
	}
	token := jwt.New(jwt.GetSigningMethod("none"))
	token.Claims = jwt.MapClaims{
		"user": u.Name,
		"auth": u.Auth,
		"exp":  1924952399,
	}
	tokenString, err := token.SignedString(jwt.UnsafeAllowNoneSignatureType)
	if err != nil {
		return nil, err
	}
	return &Authentication{Token: tokenString}, nil
}
