package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Opening",
	})
}
