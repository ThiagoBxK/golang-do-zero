package main

import (
	"errors"
	"fmt"
	"log"
)

func sum(a, b int) (int, error) {
	result := a + b

	// Caso a soma total seja maior que 10 um erro será gerado
	if result > 10 {
		return 0, errors.New("Ops, valor máximo exedido!")
	}

	return result, nil
}

func main() {
	result, err := sum(5, 10)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(result)
}
