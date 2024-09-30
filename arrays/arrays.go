package main

import "fmt"

func main() {
	// []int -> Não especifica o tamanho
	// [10]int -> Especifica o tamanho

	// Cria um array de 0 a 5 elementos
	var users = make([]int, 0, 5)
	fmt.Println(users[0:5])

	var numbers = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// Adicionar elementos, cria um novo pedaço
	numbers = append(numbers, 10)

	var total int = 0

	for index := 0; index < len(numbers); index++ {
		total += numbers[index]
		fmt.Println(numbers[index])
	}

	fmt.Println(numbers, total)

	// Indice 5 ate o 10
	fmt.Println(numbers[5:10])
	fmt.Println(numbers[5:])

	// Indice inicial ate o 5
	fmt.Println(numbers[:5])

	fmt.Println(cap(numbers), len(numbers))

	// Array de arrays
	var board = [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	fmt.Println(board)
}
