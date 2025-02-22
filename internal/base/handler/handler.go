package basehandler

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	baseerror "github.com/ladmakhi81/learning-management-system/internal/base/error"
)

type Handler func(handler *gin.Context) (*Response, error)

type Response struct {
	Data       any  `json:"data"`
	StatusCode uint `json:"statusCode"`
}

func NewResponse(data any, statusCode uint) *Response {
	return &Response{
		Data:       data,
		StatusCode: statusCode,
	}
}

func BaseHandler(handler Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res, err := handler(ctx)
		if err != nil {
			// client error
			clientErr, clientErrOk := err.(baseerror.ClientErr)
			if clientErrOk {
				ctx.JSON(
					int(clientErr.StatusCode),
					gin.H{"message": clientErr.Message},
				)

				return
			}

			// client validation error
			clientValidationErr, clientValidationErrOk := err.(baseerror.ClientValidationErr)
			if clientValidationErrOk {
				ctx.JSON(
					int(clientErr.StatusCode),
					gin.H{"message": clientValidationErr.Message, "errors": clientValidationErr.Errors},
				)

				return
			}

			// server error
			serverErr, serverErrOk := err.(baseerror.ServerErr)
			if serverErrOk {
				serverErrID := rand.Intn(1000000000)
				fmt.Println(serverErr)
				ctx.JSON(
					http.StatusInternalServerError,
					gin.H{
						"message": "internal server error",
						"trackID": serverErrID,
						"error":   serverErr.Message,
					},
				)
				return
			}
		}

		// success response
		ctx.JSON(int(res.StatusCode), res)
	}
}
