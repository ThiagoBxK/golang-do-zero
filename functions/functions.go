package main

import (
	"fmt"
	"math"
)

// Somar 2 números
func sum(a int, b int) int {
	return a + b
}

// Calcula a equação do segundo grau
func calcDelta(a, b, c float64) float64 {
	return math.Pow(b, 2) - 4*a*c
}

type QuadraticEquation struct {
	x1 float64
	x2 float64
}

func calcQuadracticEquation(a, b, c float64) QuadraticEquation {
	var delta = calcDelta(a, b, c)

	x1 := func() float64 {
		return (-b + math.Sqrt(delta)) / (2 * a)
	}
	x2 := func() float64 {
		return (-b - math.Sqrt(delta)) / (2 * a)
	}

	return QuadraticEquation{x1: x1(), x2: x2()}
}

func main() {
	fmt.Println(calcQuadracticEquation(2, -4, 1))
}
