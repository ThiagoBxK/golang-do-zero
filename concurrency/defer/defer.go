package main

import "fmt"

func finalyMessage() {
	fmt.Println("2 - A execução do programa foi concluída.")
}

func main() {
	// Função que será executada no final do programa
	defer finalyMessage()

	// Função anonima que também será executada ao final do programa
	defer func() {
		fmt.Println("1 - A execução do programa foi concluída.")
	}()

	fmt.Println("Rodando programa...")
}
