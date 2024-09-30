package main

import "fmt"

func message() string {
	fmt.Println("Chegou aqui!")
	return "Olá"
}

func main() {
	defer fmt.Println("Carregou por ultimo!")
	defer fmt.Println(message()) // Vai rodar a função e adiar o retorno

	fmt.Println("Carregou primeiro!")
}
