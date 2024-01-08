package main

import (
	"fmt"
	"os"

	"github.com/gaspartv/API-GO-webhook-meta/configs"
	"github.com/gaspartv/API-GO-webhook-meta/routers"
	"github.com/joho/godotenv"
)

var (
	logger configs.Logger
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Erro ao carregar o arquivo .env")
		os.Exit(1)
	}

	logger = *configs.GetLogger("main")

	routers.Initialize()
}
