package ctfsendai2018

import (
	"github.com/gin-gonic/gin"
)

// AuthHandler is handler manipulating authentication information
type AuthHandler interface {
	Register(*gin.Context)
	Login(*gin.Context)
}

// NewAuthHandler is create new instance
func NewAuthHandler(uc AuthUsecase) AuthHandler {
	return &authHandler{uc}
}

type authHandler struct {
	uc AuthUsecase
}

func (h *authHandler) Register(*gin.Context) {
}

func (h *authHandler) Login(*gin.Context) {

}
