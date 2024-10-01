package main

import (
	"errors"
	"fmt"
	"log"
)

func sum(a, b int) (int, error) {
	result := a + b

	if result > 10 {
		return 0, errors.New("Ops! valor máximo excedido")
	}

	return result, nil
}

func main() {
	res, err := sum(5, 99)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(res)
}
