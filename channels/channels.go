package main

import (
	"fmt"
	"sync"
	"time"
)

// Simula a leitura de um arquivo grande
func readFile(wg *sync.WaitGroup, result chan string) {
	defer wg.Done()

	time.Sleep(time.Millisecond * 500)
	result <- "Arquivo lido..."
}

// Simula a requisicao de uma api dos dados do usuario
func getUserData(wg *sync.WaitGroup, result chan string) {
	defer wg.Done()

	time.Sleep(time.Millisecond * 300)
	result <- "Dados do usúario..."
}

func main() {
	var startTime = time.Now()
	var wg sync.WaitGroup              // Cria um WaitGroup
	var results = make(chan string, 2) // Cria um canal de tamanho 2

	wg.Add(2) // Espera que duas goruntines sejam concluidas para continuar o programa

	go readFile(&wg, results)
	go getUserData(&wg, results)

	wg.Wait()      // Espera as goruntines terminar
	close(results) // Fecha o canal após o termino das goruntines

	// Lé os valores do canal
	for value := range results {
		fmt.Println(value)
	}

	// Tempo decorrido
	fmt.Print(time.Since(startTime))
}
