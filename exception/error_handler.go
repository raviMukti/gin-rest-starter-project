package exception

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/raviMukti/gin-rest-starter-project/model/web"
)

func ErrorHandler(ctx *gin.Context, err interface{}) {

	if notFoundError(ctx, err) {
		return
	}

	if validationErrors(ctx, err) {
		return
	}

	internalServerError(ctx, err)
}

func notFoundError(ctx *gin.Context, err interface{}) bool {
	exception, ok := err.(NotFoundError)

	if ok {
		ctx.Writer.Header().Add("Content-Type", "application/json")
		ctx.Writer.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Error:       "API_ERROR",
			ErrorDetail: exception.Error,
			Code:        http.StatusNotFound,
			Message:     http.StatusText(http.StatusNotFound),
			Data:        nil,
		}

		ctx.JSON(http.StatusNotFound, webResponse)

		return true
	} else {
		return false
	}
}

func validationErrors(ctx *gin.Context, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)

	if ok {
		ctx.Writer.Header().Add("Content-Type", "application/json")
		ctx.Writer.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Error:       "API_ERROR",
			ErrorDetail: exception.Error(),
			Code:        http.StatusBadRequest,
			Message:     http.StatusText(http.StatusBadRequest),
			Data:        nil,
		}

		ctx.JSON(http.StatusBadRequest, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(ctx *gin.Context, err interface{}) {
	ctx.Writer.Header().Add("Content-Type", "application/json")
	ctx.Writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Error:       "API_ERROR",
		ErrorDetail: "SERVER_ERROR",
		Code:        http.StatusInternalServerError,
		Message:     http.StatusText(http.StatusInternalServerError),
		Data:        nil,
	}

	ctx.JSON(http.StatusInternalServerError, webResponse)
}
