package main

import (
	"fmt"
	"math"
)

func solveQuadraticEquation(a, b, c float64) []float64 {
	// Calculo de delta
	var delta = math.Pow(b, 2) - 4*a*c

	// Retorna um array com os valores de X1 & X2
	return []float64{
		(-b + math.Sqrt(delta)) / (2 * a),
		(-b - math.Sqrt(delta)) / (2 * a),
	}
}

func main() {
	fmt.Println(solveQuadraticEquation(2, -4, 1))
}
