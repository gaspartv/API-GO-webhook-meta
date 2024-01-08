package main

import (
	"os"

	"github.com/gaspartv/API-GO-webhook-meta/configs"
	"github.com/gaspartv/API-GO-webhook-meta/routers"
)

var (
	logger configs.Logger
)

func main() {
	// Environment variables
	envVariables := make(map[string]string)
	envVariables["FACEBOOK_VERIFY_TOKEN"] = "e87e4588-794a-4d5b-91c6-28699a275c7c"
	for key, value := range envVariables {
		os.Setenv(key, value)
	}

	logger = *configs.GetLogger("main")

	// Initialize Router
	routers.Initialize()
}
