package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func WhatsappValidateHandler(ctx *gin.Context) {
	mode := ctx.Query("hub.mode")
	if mode != "subscribe" {
		logger.Error("Unsupported mode specified hub.mode")
		sendError(ctx, http.StatusBadRequest, "Unsupported mode specified hub.mode")
		return
	}

	verifyToken := ctx.Query("hub.verify_token")
	if verifyToken != os.Getenv("WHATSAPP_VERIFY_TOKEN") {
		logger.Error("Check if the token is valid")
		sendError(ctx, http.StatusBadRequest, "Check if the token is valid")
		return
	}

	challenge := ctx.Query("hub.challenge")
	if challenge == "" {
		logger.Error("Challenge is empty")
		sendError(ctx, http.StatusBadRequest, "Challenge is empty")
		return
	}
	sendSuccess(ctx, challenge)
}
