package handler

import (
	"github.com/Gabriel-Bitencort/gopportunities.git/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}
