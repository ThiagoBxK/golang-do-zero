package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	// Escreve na tela um Hello World!
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	// Cria uma rota que executa a funcao handler
	http.HandleFunc("/", handleHome)

	address := "localhost:8080"
	fmt.Println("Server is running on: ", address)

	err := http.ListenAndServe(address, nil)

	if err != nil {
		log.Fatal("Error starting server")
	}
}
