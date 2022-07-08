package service

import (
	"os"

	"github.com/gin-gonic/gin"
)

type HelloService interface {
	Greeting(ctx *gin.Context) gin.H
}

type HelloServiceImpl struct {
}

func NewHelloServiceImpl() HelloService {
	return &HelloServiceImpl{}
}

// Greeting implements HelloService
func (service *HelloServiceImpl) Greeting(ctx *gin.Context) gin.H {

	return gin.H{
		"app_name": os.Getenv("APP_NAME"),
		"message":  "Hello World!",
	}
}
