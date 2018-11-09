package ctfsendai2018

import (
	"net/http"

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

func (h *userHandler) Add(c *gin.Context) {

}

func (h *userHandler) List(c *gin.Context) {
	list, err := h.uc.List(1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"error":   true,
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "success",
		"error":   false,
		"users":   list,
	})
}
