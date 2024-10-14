package main

import (
	"fmt"
	"regexp"
)

func countVowel(str string) int {
	regex := regexp.MustCompile(`[aeiouAEIOU]`)
	vowels := regex.FindAllString(str, -1)

	return len(vowels)
}

func main() {
	fmt.Printf("Hello World contém %v vogais", countVowel("Hello World!"))
}
