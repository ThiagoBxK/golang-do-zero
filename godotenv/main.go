package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func getEnv(key string) string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	value := os.Getenv(key)
	return value
}

func main() {
	token := getEnv("ACCESS_TOKEN")

	fmt.Println(token)
}
