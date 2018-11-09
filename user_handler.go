package ctfsendai2018

import (
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
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
	token, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		b := jwt.UnsafeAllowNoneSignatureType
		return b, nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  http.StatusUnauthorized,
			"message": err.Error(),
			"error":   true,
		})
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	auth := claims["auth"].(float64)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	u := &User{}
	if err := c.BindJSON(u); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"error":   true,
		})
		return
	}
	if err := h.uc.Add(uint8(auth), u); err != nil {
		if strings.Index(err.Error(), "Don't have permission") != -1 {
			c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"status":  http.StatusUnauthorized,
				"message": err.Error(),
				"error":   true,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"error":   true,
		})
		return
	}
	c.JSON(http.StatusCreated, map[string]interface{}{
		"status":  http.StatusCreated,
		"message": "success",
		"error":   false,
		"user":    u,
		"flag":    "Congratulations!!You_have_cleared_all_the_problems",
	})
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
