package web

import (
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Code int    `json:"code"`
	Status     string `json:"status"`
	Message    string `json:"message"`
}

type response struct {
	Data interface{} `json:"data"`
}

func Failure(ctx *gin.Context, code int, status, message string) {
	ctx.JSON(code, errorResponse{
		Code: code,
		Status:     status,
		Message:    message,
	})
}

func Success(ctx *gin.Context, code int, data interface{}) {
	ctx.JSON(code, response{data})
}