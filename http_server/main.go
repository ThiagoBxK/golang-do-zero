package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func handlerHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Olá!")
}

func getEnv(key string) string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	value := os.Getenv(key)
	return value
}

func main() {
	http.HandleFunc("/", handlerHome)

	address := getEnv("HTTP_ADDRESS")
	fmt.Println("Server is running on: ", address)

	err := http.ListenAndServe(address, nil)

	if err != nil {
		log.Fatal("Error starting server")
	}

}
