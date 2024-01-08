package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gaspartv/API-GO-webhook-meta/rabbitmq"
	"github.com/gin-gonic/gin"
)

func FacebookReceiveHandler(ctx *gin.Context) {
	rmq, err := rabbitmq.NewRabbitMQ(os.Getenv("RABBIT_MQ_URL"))
	if err != nil {
		logger.ErrorF("Connection failed: %v\n", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer rmq.Close()

	var request FacebookReceiveRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := request.Validate(); err != nil {
		logger.ErrorF("Validation error: %v\n", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	body, err := json.Marshal(&request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	err = rmq.Publish(os.Getenv("RABBIT_MQ_SEND"), body)
	if err != nil {
		logger.ErrorF("Failed to publish message to RabbitMQ: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	sendSuccess(ctx, "Successful")
}
