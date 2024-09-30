package main

import "fmt"

func main() {
	var year int = 2024

	// Endereço de memoria da variavel year
	var pointer *int = &year

	fmt.Println(pointer, *pointer)

	// Desreferencia o ponteiro e atribui um novo valor
	*pointer = 2025
	fmt.Println(year)
}
