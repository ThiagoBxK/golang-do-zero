package main

import "fmt"

func main() {
	// Simulando download
	for i := 1; i < 101; i++ {
		fmt.Printf("Baixando... %d%%\n", i)
	}

	// Equivalente ao while
	index := 0

	for index < 10 {
		fmt.Println(index)
		index++
	}

	// Usando range
	users := []string{"Ana", "Maria", "Pedro", "Juliana"}

	for index, name := range users {
		fmt.Println(index, name)
	}
}
