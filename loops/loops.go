package main

import (
	"fmt"
)

func main() {
	// Simulando download
	for i := 0; i < 100; i++ {
		fmt.Printf("Baixando... %d%%\n", i+1)
	}

	// While loop
	var sum int = 2

	for sum < 100 {
		fmt.Printf("Baixando... %d%%\n", sum+1)
		sum++
	}
}
