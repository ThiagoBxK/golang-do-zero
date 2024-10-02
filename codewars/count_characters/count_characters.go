package main

import "fmt"

// Desafio CodeWars : Count characters in your string

func count(str string) map[string]int {
	countMap := make(map[string]int)

	for _, char := range str {
		countMap[string(char)]++
	}

	return countMap
}

func main() {
	fmt.Println(count("abccc"))
}
