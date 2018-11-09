package ctfsendai2018

import (
	"net/http"

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

func (h *authHandler) Register(c *gin.Context) {
	m := make(map[string]string)
	if err := c.BindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"error":   true,
		})
		return
	}
	if err := h.uc.Register(m["email"], m["password"]); err != nil {
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
	})
}

func (h *authHandler) Login(c *gin.Context) {
	m := make(map[string]string)
	if err := c.BindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"error":   true,
		})
		return
	}
	auth, err := h.uc.Login(m["email"], m["password"])
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
		"token":   auth.Token,
		"error":   false,
		"flag":    "TmV4dCBGTEFHOiBDb25ncmF0dWxhdGlvbnMhIVlvdV9oYXZlX2NsZWFyZWRfYWxsX3RoZV9wcm9ibGVtcw",
	})
}
