package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/Naw3/reinchat/internal/llm"
)

func main() {

	// Load env variables
	err := godotenv.Load(".env.dev")
	if err != nil {
		log.Fatalf("Error loading .env.dev file: %v", err)
	}

	url := os.Getenv("BASE_URL")
	model := os.Getenv("MODEL")
	apiKey := os.Getenv("API_KEY")

	client := llm.NewClient(url, apiKey, model)

	bodyResponse, err := client.Chat("Hi, how are you today?")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("API Response :")
	fmt.Println(string(bodyResponse))
}
