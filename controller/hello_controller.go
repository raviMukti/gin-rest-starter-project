package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raviMukti/gin-rest-starter-project/model/web"
	"github.com/raviMukti/gin-rest-starter-project/service"
)

type HelloController interface {
	Greeting(ctx *gin.Context)
}

type HelloControllerImpl struct {
	HelloService service.HelloService
}

func NewHelloController(service service.HelloService) HelloController {
	return &HelloControllerImpl{
		HelloService: service,
	}
}

// Greeting implements HelloController
func (controller *HelloControllerImpl) Greeting(ctx *gin.Context) {
	response := controller.HelloService.Greeting(ctx)

	webResponse := web.WebResponse{
		Error:       "",
		ErrorDetail: "",
		Code:        http.StatusOK,
		Message:     http.StatusText(http.StatusOK),
		Data:        response,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
