package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Opening",
	})
}
