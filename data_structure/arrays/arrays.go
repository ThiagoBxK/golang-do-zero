package main

import "fmt"

func main() {
	// Array um tipo de dado com tamanho fixo
	// Tamanho do array no momento de build
	arrayA := [5]int{1, 2, 3, 4, 5}
	arrayB := [5]int{5, 4, 3, 2, 1}

	// Array de arrays
	matrix := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	// For usando range
	for _, value := range arrayA {
		fmt.Println(value)
	}

	// For em uma matriz
	for row, _ := range matrix {
		for column, _ := range matrix[row] {
			fmt.Printf("Linha: %d, Coluna: %d\n", row, column)
		}
	}

	// Arrays podem ser comparados
	fmt.Printf("Arrays iguais: %t\n", arrayA == arrayB)

	// Visualizar tamanho de um array
	fmt.Printf("Tamanho do array: %d\n", len(arrayA))

	// Fatiando um array
	fmt.Println(arrayA[:2]) // Pegando do índice 0 até o 2
	fmt.Println(arrayA[2:]) // Pegando do índice 2 até o final do array
}
