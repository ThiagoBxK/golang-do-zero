package main

import (
	"fmt"
	"strings"
)

func reverseString(input string) string {
	var result string
	for index := range input {
		result += string(input[len(input)-index-1])
	}

	return result
}

func isPalindrome(input string) bool {
	input = strings.ReplaceAll(strings.ToLower(input), " ", "")
	return input == reverseString(input)
}

func main() {
	fmt.Printf(`A string é um palíndromo "A man a plan a canal Panama" %v`, isPalindrome("A man a plan a canal Panama"))
}
