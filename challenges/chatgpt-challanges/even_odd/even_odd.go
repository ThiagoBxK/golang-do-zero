package main

import "fmt"

func checkEvenOdd(num int) string {
	switch num%2 == 0 {
	case true:
		return "Par"
	default:
		return "Impár"
	}
}

func main() {
	fmt.Printf("Número 4 é: %s", checkEvenOdd(4))
}
