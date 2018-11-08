package main

import (
	"github.com/135yshr/ctfsendai2018"
	"github.com/gin-gonic/gin"
)

var reg *registory

type registory struct {
	UserRepository ctfsendai2018.UserRepository
	UserUsecase    ctfsendai2018.UserUsecase
	UserHandler    ctfsendai2018.UserHandler

	AuthRepository ctfsendai2018.AuthRepository
	AuthUsecase    ctfsendai2018.AuthUsecase
	AutHandler     ctfsendai2018.AuthHandler
}

func main() {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.POST("/register", reg.AutHandler.Register)
		v1.POST("/login", reg.AutHandler.Login)

		v1.GET("/users", reg.UserHandler.List)
		v1.POST("/users", reg.UserHandler.Add)
	}
	r.Run()
}

func init() {
	reg = &registory{}
	reg.UserRepository = ctfsendai2018.NewUserRepository()
	reg.UserUsecase = ctfsendai2018.NewUserUsecase(reg.UserRepository)
	reg.UserHandler = ctfsendai2018.NewUserHandler(reg.UserUsecase)

	reg.AuthRepository = ctfsendai2018.NewAuthRepository()
	reg.AuthUsecase = ctfsendai2018.NewAuthUsecase(reg.AuthRepository, reg.UserRepository)
	reg.AutHandler = ctfsendai2018.NewAuthHandler(reg.AuthUsecase)
}
