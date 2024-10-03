package main

import "fmt"

func counter() (func(int) int, func(int) int) {
	counter := 0

	// Retorna duas funções increment & decrement
	return func(increment int) int {
			counter += increment
			return counter
		}, func(decrement int) int {
			counter += decrement
			return counter
		}
}

func main() {
	increment, decrement := counter()

	fmt.Printf("Valor atual: %d\n", increment(10))
	fmt.Printf("Valor atual: %d\n", decrement(5))
}
