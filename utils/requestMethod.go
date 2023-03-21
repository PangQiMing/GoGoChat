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
