package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/raviMukti/gin-rest-starter-project/controller"
	"github.com/raviMukti/gin-rest-starter-project/exception"
	"github.com/raviMukti/gin-rest-starter-project/helper"
	"github.com/raviMukti/gin-rest-starter-project/model/web"
	"github.com/raviMukti/gin-rest-starter-project/service"
)

func Init() {
	err := godotenv.Load()
	helper.PanicIfError(err)
}

var (
	helloService    service.HelloService       = service.NewHelloServiceImpl()
	helloController controller.HelloController = controller.NewHelloController(helloService)
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.HandleMethodNotAllowed = true
	// Register Panic Handler
	router.Use(gin.CustomRecovery(exception.ErrorHandler))
	// Register No Route Handler
	router.NoRoute(noRoutetoHost)
	// Register Method Not Allowed Handler
	router.NoMethod(noMethodAllowed)

	hello := router.Group("/hello")
	{
		hello.GET("/", helloController.Greeting)
	}

	return router
}

func noRoutetoHost(ctx *gin.Context) {
	ctx.Writer.Header().Add("Content-Type", "application/json")
	ctx.Writer.WriteHeader(http.StatusNotFound)

	webResponse := web.WebResponse{
		Error:       "API_ERROR",
		ErrorDetail: ctx.Request.URL.Path + " NOT FOUND",
		Code:        http.StatusNotFound,
		Message:     http.StatusText(http.StatusNotFound),
		Data:        "",
	}

	ctx.JSON(http.StatusNotFound, webResponse)
}

func noMethodAllowed(ctx *gin.Context) {
	ctx.Writer.Header().Add("Content-Type", "application/json")
	ctx.Writer.WriteHeader(http.StatusMethodNotAllowed)

	webResponse := web.WebResponse{
		Error:       "API_ERROR",
		ErrorDetail: ctx.Request.Method + " NOT ALLOWED",
		Code:        http.StatusMethodNotAllowed,
		Message:     http.StatusText(http.StatusMethodNotAllowed),
		Data:        "",
	}

	ctx.JSON(http.StatusMethodNotAllowed, webResponse)
}
