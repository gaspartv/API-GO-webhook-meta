package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gaspartv/API-GO-webhook-meta/configs"
	"github.com/gaspartv/API-GO-webhook-meta/rabbitmq"
	"github.com/gaspartv/API-GO-webhook-meta/routers"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

type Message struct {
	PhoneNumberId string          `json:"phoneNumberId"`
	AppToken      string          `json:"appToken"`
	Data          json.RawMessage `json:"data"`
}

var (
	logger configs.Logger
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Erro ao carregar o arquivo .env")
		os.Exit(1)
	}

	logger = *configs.GetLogger("main")

	url := os.Getenv("RABBIT_MQ_URL")
	rmq, err := rabbitmq.NewRabbitMQ(url)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer rmq.Close()

	go func() {
		routers.Initialize()
	}()

	queueNames := []string{os.Getenv("RABBIT_MQ_RECEIVED")}

	for _, queueName := range queueNames {
		err = rmq.DeclareQueue(queueName)
		if err != nil {
			log.Fatalf("Failed to declare queue %s: %v", queueName, err)
		}
		msgs, err := rmq.Consume(queueName)
		if err != nil {
			log.Fatalf("Failed to register a consumer for queue %s: %v", queueName, err)
		}

		go consumeMessages(queueName, msgs)
	}

	forever := make(chan bool)
	fmt.Println("[RabbitMQ-debug] Listening for messages...")
	<-forever
}

func consumeMessages(queueName string, msgs <-chan amqp.Delivery) {
	metaURL := os.Getenv("META_URL")
	for d := range msgs {
		var msg Message
		err := json.Unmarshal(d.Body, &msg)
		if err != nil {
			log.Printf("Failed to unmarshal message: %v", err)
			continue
		}

		apiUrl := fmt.Sprintf("%s/%s/messages?access_token=%s", metaURL, msg.PhoneNumberId, msg.AppToken)
		err = sendMessageToMetaAPI(msg.Data, apiUrl)
		if err != nil {
			log.Printf("Failed to send message to META API: %v", err)
			continue
		}
	}
}

func sendMessageToMetaAPI(body []byte, url string) error {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send message, status code: %d", resp.StatusCode)
	}

	fmt.Println("Message successfully sent to META API")
	return nil
}
