package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/135yshr/ctfsendai2018"
	"github.com/gin-gonic/gin"
)

const importFilePath = "./users.txt"

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
	if err := createUsers(); err != nil {
		fmt.Println(err)
		return
	}
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

func createUsers() error {
	path, err := filepath.Abs(importFilePath)
	if err != nil {
		return err
	}
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		if err := sc.Err(); err != nil {
			return err
		}
		text := sc.Text()
		cols := strings.Split(text, "\t")
		age, err := strconv.ParseUint(cols[4], 10, 8)
		if err != nil {
			fmt.Println(err)
		}
		auth, err := strconv.ParseUint(cols[6], 10, 8)
		if err != nil {
			fmt.Println(err)
			auth = 2
		}
		u := &ctfsendai2018.User{
			ID:       cols[0],
			EMail:    cols[1],
			Password: cols[2],
			Name:     cols[3],
			Age:      uint8(age),
			Note:     cols[5],
			Auth:     uint8(auth),
		}
		if err := reg.UserRepository.Add(u); err != nil {
			return err
		}
		fmt.Printf("Success %v", u)
	}
	return nil
}
