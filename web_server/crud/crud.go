package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Estrutura da api
type Posts []struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// Função para ler arquivos
func readFile(filePath string) []byte {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Erro ao ler o arquivo: %v", err)
	}

	return data
}

// Obter posts
func getPosts() Posts {
	data := readFile("D:/Projetos/golang-do-zero/web_server/crud/posts.json")

	var posts Posts
	err := json.Unmarshal(data, &posts)

	if err != nil {
		log.Fatalf("Erro ao decodificar o arquivo: %v", err)
	}

	return posts
}

func handleUser(w http.ResponseWriter, r *http.Request) {
	posts := getPosts()

	// Retornar os posts como JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func main() {
	http.HandleFunc("/posts", handleUser)
	fmt.Println("Servidor rodando na porta: ", "8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error ao iniciar o servidor: %w", err)
	}
}
