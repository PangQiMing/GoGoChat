package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// RequestMethodPost 请求方法POST
func RequestMethodPost(ctx *gin.Context) {
	if ctx.Request.Method != "POST" {
		ctx.Writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

// RequestMethodGet 请求方法GET
func RequestMethodGet(ctx *gin.Context) {
	if ctx.Request.Method != "GET" {
		ctx.Writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

// RequestMethodPut 请求方法Put
func RequestMethodPut(ctx *gin.Context) {
	if ctx.Request.Method != "PUT" {
		ctx.Writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

// RequestMethodDelete 请求方法Delete
func RequestMethodDelete(ctx *gin.Context) {
	if ctx.Request.Method != "DELETE" {
		ctx.Writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
