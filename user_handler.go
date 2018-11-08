package ctfsendai2018

import (
	"github.com/gin-gonic/gin"
)

// UserHandler is handler manipulating authentication information
type UserHandler interface {
	Add(*gin.Context)
	List(*gin.Context)
}

// NewUserHandler is create new instance
func NewUserHandler(uc UserUsecase) UserHandler {
	return &userHandler{uc}
}

type userHandler struct {
	uc UserUsecase
}

func (h *userHandler) Add(*gin.Context) {

}

func (h *userHandler) List(*gin.Context) {

}
