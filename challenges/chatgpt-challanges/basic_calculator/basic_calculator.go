package main

import "fmt"

func calc(a, b float64, operator string) float64 {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			fmt.Println("Erro: Divisão por zero")
			return 0
		}
		return a / b
	default:
		fmt.Println("Operador inválido")
		return 0
	}
}

func main() {
	fmt.Print(calc(5, 2, "+"))
}
