package handlers

import (
	"github.com/gaspartv/API-GO-webhook-meta/configs"
	"github.com/gaspartv/API-GO-webhook-meta/rabbitmq"
)

var (
	logger *configs.Logger
)

var rmq *rabbitmq.RabbitMQ

func InitializeHandler() {
	logger = configs.GetLogger("handler")
}
