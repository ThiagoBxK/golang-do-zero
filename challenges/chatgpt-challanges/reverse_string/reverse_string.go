package main

import "fmt"

func reverseString(input string) string {
	var result string

	// Neste loop, a string é concatenada repetidamente,
	// o que pode ser ineficiente, especialmente com strings grandes.
	// Uma alternativa seria usar um slice de runas para melhor performance.
	for index := range input {
		result += string(input[len(input)-index-1])
	}

	return result
}

func main() {
	fmt.Printf("Hello ao inverso %s", reverseString("Hello"))
}
