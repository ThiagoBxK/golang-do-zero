package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Função que obtém uma variável de ambiente por chave
func getEnvironmentVar(key string) string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	value := os.Getenv(key)
	return value
}

func main() {
	fmt.Println(getEnvironmentVar("API_KEY"))
}
